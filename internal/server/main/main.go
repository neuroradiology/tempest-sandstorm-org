package servermain

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"zenhack.net/go/tempest/internal/server/database"
	"zenhack.net/go/tempest/internal/server/logging"
	"zenhack.net/go/tempest/internal/server/session"
	"zenhack.net/go/util"
)

func Main() {
	initStorage()
	lg := logging.NewLogger()
	cfg := ConfigFromSettings(lg)
	listenAddr := ":" + cfg.HTTP.Port
	db := util.Must(database.Open())
	sessionStore := session.NewStore(util.Must(session.GetKeys()))
	srv := newServer(cfg, lg, db, sessionStore)
	defer srv.Release()

	http.Handle("/", srv.Handler())
	lg.Info("Listening",
		"root-domain", cfg.HTTP.RootDomain,
		"listen-addr", listenAddr,
	)
	httpSrv := &http.Server{Addr: listenAddr}
	go monitorSignals(httpSrv)
	panic(httpSrv.ListenAndServe())
}

func monitorSignals(srv *http.Server) {
	defer srv.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		// Signals that would normally kill us. Instead, stop the server
		// and let main() do shutdown.
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	defer signal.Stop(sigs)
	<-sigs
}
