package servermain

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"capnproto.org/go/capnp/v3"
	cpserver "capnproto.org/go/capnp/v3/server"
	"sandstorm.org/go/tempest/capnp/collection"
	"sandstorm.org/go/tempest/capnp/external"
	"sandstorm.org/go/tempest/capnp/grain"
	utilcp "sandstorm.org/go/tempest/capnp/util"
	grainagent "sandstorm.org/go/tempest/internal/capnp/grain-agent"
	"sandstorm.org/go/tempest/internal/capnp/system"
	"sandstorm.org/go/tempest/internal/common/types"
	"sandstorm.org/go/tempest/internal/config"
	"sandstorm.org/go/tempest/internal/server/container"
	"sandstorm.org/go/tempest/internal/server/database"
	"sandstorm.org/go/tempest/internal/server/session"
	"sandstorm.org/go/tempest/internal/server/tokenutil"
	"sandstorm.org/go/tempest/pkg/exp/util/assign"
	"zenhack.net/go/util"
	"zenhack.net/go/util/exn"
)

var ErrNotLoggedIn = errors.New("you are not logged in")

type externalApiImpl struct {
	server       *server
	userSession  session.UserSession
	sessionStore session.Store
}

func (api externalApiImpl) GetSessions(ctx context.Context, p external.ExternalApi_getSessions) error {
	if api.userSession.Credential.Type == "" {
		return ErrNotLoggedIn
	}
	return exn.Try0(func(throw exn.Thrower) {
		results, err := p.AllocResults()
		throw(err)
		p.Go()

		tx, err := api.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		role, err := tx.CredentialRole(api.userSession.Credential)
		throw(err)

		visitor := visitorSessionImpl{
			externalApiImpl: api,
		}
		results.SetVisitor(external.VisitorSession_ServerToClient(visitor))

		if role.Encompasses(types.RoleUser) {
			user := userSessionImpl{visitor: visitor}
			results.SetUser(external.UserSession_ServerToClient(user))
		}

		// TODO: admin
	})
	return nil
}

func (api externalApiImpl) Restore(ctx context.Context, p external.ExternalApi_restore) error {
	p.Go()
	return exn.Try0(func(throw exn.Thrower) {
		token, err := p.Args().SturdyRef()
		throw(err)
		results, err := p.AllocResults()
		throw(err)
		tx, err := api.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		v, err := tx.RestoreSturdyRef(database.SturdyRefKey{
			Token:     token,
			OwnerType: "external-api",
		})
		throw(err)
		if v.GrainID == "" {
			oid := system.SystemObjectId(v.ObjectID)
			switch oid.Which() {
			case system.SystemObjectId_Which_sharingToken:
				s := oid.SharingToken()
				id, err := s.GrainId()
				throw(err)
				_, seg := capnp.NewMultiSegmentMessage(nil)
				kv, err := utilcp.NewKeyValue(seg)
				throw(err)
				key, err := capnp.NewText(seg, id)
				throw(err)
				throw(kv.SetKey(key.ToPtr()))
				view, err := external.NewUiView(seg)
				throw(err)
				info, err := tx.GrainInfo(types.GrainID(id))
				throw(err)
				throw(view.SetTitle(info.Title))
				sessionToken, err := session.GrainSession{
					GrainID:   info.ID,
					SessionID: api.userSession.SessionID,
				}.Seal(api.sessionStore)
				throw(err)
				throw(view.SetSessionToken(sessionToken))
				throw(view.SetSubdomain(hex.EncodeToString(tokenutil.GenToken()[:16])))
				throw(view.SetController(external.UiView_Controller_ServerToClient(uiViewControllerImpl{
					GrainID: info.ID,
					Session: api.userSession,
					DB:      api.server.db,
				})))
				throw(kv.SetValue(view.ToPtr()))
				throw(results.SetCap(capnp.Client(assign.FixedGetter(kv.ToPtr()))))
			default:
				throw(fmt.Errorf("Restore not supported on system objects of type %v", oid.Which()))
			}
		} else {
			throw(errors.New("TODO: implement ExternalApi.restore() for non-system objects"))
		}
	})
}

func (api externalApiImpl) Authenticator(ctx context.Context, p external.ExternalApi_authenticator) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetAuthenticator(external.Authenticator_ServerToClient(authenticatorImpl{
		api: api,
	}))
	return nil
}

type authenticatorImpl struct {
	api externalApiImpl
}

func (a authenticatorImpl) SendEmailAuthToken(ctx context.Context, p external.Authenticator_sendEmailAuthToken) error {
	return exn.Try0(func(throw exn.Thrower) {
		addr, err := p.Args().Address()
		throw(err)
		db := a.api.server.db
		tx, err := db.Begin()
		throw(err)
		defer tx.Rollback()

		// FIXME: sanitize addr?

		_, seg := capnp.NewSingleSegmentMessage(nil)
		oid, err := system.NewRootSystemObjectId(seg)
		throw(err)
		throw(oid.SetEmailLoginToken(addr))

		token := tokenutil.Gen128Base64()

		_, err = tx.SaveSturdyRef(
			database.SturdyRefKey{
				Token:     []byte(token),
				OwnerType: "external",
				Owner:     "",
			},
			database.SturdyRefValue{
				Expires:  time.Now().Add(10 * time.Minute),
				ObjectID: capnp.Struct(oid),
			},
		)
		throw(err)
		throw(tx.Commit())

		cfg := a.api.server.cfg
		throw(cfg.SMTP.SendMail(
			[]string{addr},
			[]byte(strings.Join([]string{
				"To: " + addr,
				"From: " + cfg.SMTP.Username,
				"Subject: Email Login Token",
				"",
				"Login in as " + addr + " by visiting:",
				"",
				cfg.HTTP.RootDomain + "/login/email/" + token,
				"",
				"Or entering " + token + " at the login prompt.",
			}, "\r\n")),
		))
	})
}

type visitorSessionImpl struct {
	externalApiImpl
}

func (s visitorSessionImpl) Views(ctx context.Context, p external.VisitorSession_views) error {
	res, err := p.AllocResults()
	if err != nil {
		return err
	}
	return res.SetViews(external.UiView_Keyring_ServerToClient(viewsPuller{
		externalApiImpl: s.externalApiImpl,
	}))
}

type viewsPuller struct {
	externalApiImpl
}

func (viewsPuller) Key(ctx context.Context, p collection.Puller_key) error {
	return capnp.Unimplemented("not implemented")
}

func (vp viewsPuller) Sync(ctx context.Context, p collection.Puller_sync) error {
	into := p.Args().Into()
	p.Go()
	return exn.Try0(func(throw exn.Thrower) {
		// TODO(cleanup): update our wrapper to support one-off queries without having to
		// create a whole transaction; this is too much boilerplate.
		tx, err := vp.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		accountID, err := tx.CredentialAccount(vp.userSession.Credential)
		throw(err)
		info, err := tx.AccountKeyring(accountID).AllUiViews()
		throw(err)
		throw(tx.Commit())

		throw(into.Clear(ctx, nil))
		for _, uiViewInfo := range info {
			throw(into.Upsert(ctx, func(p utilcp.KeyValue) error {
				key, err := capnp.NewText(p.Segment(), string(uiViewInfo.Grain.ID))
				throw(err)
				p.SetKey(key.ToPtr())
				g, err := external.NewUiView(p.Segment())
				throw(err)
				g.SetTitle(uiViewInfo.Grain.Title)
				sessionToken, err := session.GrainSession{
					GrainID:   uiViewInfo.Grain.ID,
					SessionID: vp.userSession.SessionID,
				}.Seal(vp.sessionStore)
				throw(err)
				g.SetSessionToken(sessionToken)
				g.SetSubdomain(hex.EncodeToString(tokenutil.GenToken()[:16]))
				g.SetController(external.UiView_Controller_ServerToClient(uiViewControllerImpl{
					GrainID: uiViewInfo.Grain.ID,
					Session: vp.userSession,
					DB:      vp.server.db,
				}))
				p.SetValue(g.ToPtr())
				return nil
			}))
		}
		fut, rel := into.Ready(ctx, nil)
		defer rel()
		throw(into.WaitStreaming())
		_, err = fut.Struct()
		throw(err)
	})
}

func (vp viewsPuller) Attach(ctx context.Context, p external.UiView_Keyring_attach) error {
	arg := p.Args().Controller()
	snapshot := capnp.Client(arg).Snapshot()
	defer snapshot.Release()
	brand := snapshot.Brand()
	srv, ok := cpserver.IsServer(brand)
	if !ok {
		return fmt.Errorf("not a server-side view controller (brand is type %T)", brand.Value)
	}
	ctrl, ok := srv.(uiViewControllerImpl)
	if !ok {
		return fmt.Errorf("not a view controller impl (type %T)", srv)
	}
	return exn.Try0(func(throw exn.Thrower) {
		tx, err := vp.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		accountID, err := tx.CredentialAccount(ctrl.Session.Credential)
		throw(err)
		keyring := tx.AccountKeyring(accountID)
		throw(keyring.AttachGrain(ctrl.GrainID, nil)) // FIXME: fill in permissions.
		throw(tx.Commit())
	})
}

func (s userSessionImpl) ListPackages(ctx context.Context, p external.UserSession_listPackages) error {
	// TODO: too much boilerplate in common with ListGrains; factor some of this out.
	p.Go()
	into := p.Args().Into()
	return exn.Try0(func(throw exn.Thrower) {
		tx, err := s.visitor.server.db.Begin()
		throw(err)
		defer tx.Rollback()
		dbPkgs, err := tx.CredentialPackages(s.visitor.userSession.Credential)
		throw(err)
		throw(tx.Commit())

		throw(into.Clear(ctx, nil))
		for _, dbPkg := range dbPkgs {
			throw(into.Upsert(ctx, func(p utilcp.KeyValue) error {
				key, err := capnp.NewText(p.Segment(), string(dbPkg.ID))
				throw(err)
				p.SetKey(key.ToPtr())
				pkg, err := external.NewPackage(p.Segment())
				throw(err)
				throw(pkg.SetManifest(dbPkg.Manifest))
				pkg.SetController(external.Package_Controller_ServerToClient(pkgController{
					visitorSessionImpl: s.visitor,
					pkg:                dbPkg,
				}))
				// TODO: controller
				p.SetValue(pkg.ToPtr())
				return nil
			}))
		}
		fut, rel := into.Ready(ctx, nil)
		defer rel()
		throw(into.WaitStreaming())
		_, err = fut.Struct()
		throw(err)
	})
}

type pkgController struct {
	visitorSessionImpl
	pkg database.Package
}

func newGrainID() types.GrainID {
	// Oversize buffer so we don't have to think too hard about decoded vs. encoded
	// length:
	var buf [64]byte
	_, err := rand.Read(buf[:])
	util.Chkfatal(err)
	return types.GrainID(base64.URLEncoding.EncodeToString(buf[:])[:22])
}

func (pc pkgController) Create(ctx context.Context, p external.Package_Controller_create) error {
	return exn.Try0(func(th exn.Thrower) {
		args := p.Args()
		actionIndex := args.ActionIndex()
		title, err := args.Title()
		exn.WrapThrow(th, "getting title", err)

		actions, err := pc.pkg.Manifest.Actions()
		exn.WrapThrow(th, "getting actions", err)
		if actionIndex >= uint32(actions.Len()) {
			th(errors.New("actionIndex out of bounds"))
		}
		grainID := newGrainID()

		tx, err := pc.server.db.Begin()
		exn.WrapThrow(th, "creating database transaction", err)

		defer tx.Rollback()
		accountID, err := tx.CredentialAccount(pc.userSession.Credential)
		exn.WrapThrow(th, "getting account id", err)

		err = os.MkdirAll(
			config.Localstatedir+"/sandstorm/grains/"+string(grainID)+"/sandbox",
			0770,
		)
		exn.WrapThrow(th, "creating grain sandbox directory", err)
		err = tx.AddGrain(database.NewGrain{
			GrainID: grainID,
			PkgID:   pc.pkg.ID,
			Title:   title,
			OwnerID: accountID,
		})
		exn.WrapThrow(th, "creating grain in database", err)

		startArg, err := exn.Try(func(throw exn.Thrower) string {
			_, seg := capnp.NewSingleSegmentMessage(nil)
			launchCmd, err := grainagent.NewRootLaunchCommand(seg)
			throw(err)
			launchCmd.SetInitGrain(actionIndex)
			return base64.StdEncoding.EncodeToString(seg.Data())
		})
		exn.WrapThrow(th, "encoding LaunchCommand", err)

		results, err := p.AllocResults()
		th(err)

		results.SetId(string(grainID))
		v, err := results.NewView()
		th(err)
		th(v.SetTitle(title))
		sessionToken, err := session.GrainSession{
			GrainID:   grainID,
			SessionID: pc.userSession.SessionID,
		}.Seal(pc.sessionStore)
		exn.WrapThrow(th, "creating grain session token", err)
		th(v.SetSessionToken(sessionToken))
		th(v.SetController(external.UiView_Controller_ServerToClient(uiViewControllerImpl{
			GrainID: grainID,
			Session: pc.userSession,
			DB:      pc.server.db,
		})))
		exn.WrapThrow(th, "commiting database transaction", tx.Commit())

		// TODO: maybe change container.Command so it can take tx instead of a DB?
		// But probably we shouldn't do the actual spawning in a tx anyway.
		c, err := container.Command{
			Log:     pc.server.log,
			DB:      pc.server.db,
			GrainID: grainID,
			Api:     grain.SandstormApi_ServerToClient(sandstormApiImpl{}),
			Args:    []string{startArg},
		}.Start(context.TODO())
		exn.WrapThrow(th, "starting container", err)
		pc.server.state.With(func(state *serverState) {
			state.containers.containersByGrainID[grainID] = c
		})
	})

}

type userSessionImpl struct {
	visitor visitorSessionImpl
}

func (s userSessionImpl) InstallPackage(ctx context.Context, p external.UserSession_installPackage) error {
	results, err := p.AllocResults()
	if err != nil {
		return err
	}
	results.SetStream(external.Package_InstallStream_ServerToClient(newInstallStream(s)))
	return nil
}
