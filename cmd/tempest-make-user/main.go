package main

import (
	"flag"

	"sandstorm.org/go/tempest/internal/common/types"
	"sandstorm.org/go/tempest/internal/server/database"
	"sandstorm.org/go/tempest/internal/server/tokenutil"
	"zenhack.net/go/util"
)

var (
	typ      = flag.String("type", "", "credential type to use")
	scopedID = flag.String("id", "", "type-specific credential id")
	roleStr  = flag.String("role", string(types.RoleUser), "role the user should have")
)

func main() {
	accountID := types.AccountID(tokenutil.Gen128Base64())
	flag.Parse()
	role := types.Role(*roleStr)
	if !role.IsValid() {
		panic("Invalid role: " + role)
	}

	db, err := database.Open()
	util.Chkfatal(err)
	defer db.Close()
	tx, err := db.Begin()
	util.Chkfatal(err)
	defer tx.Rollback()
	util.Chkfatal(tx.AddAccount(database.NewAccount{
		ID:   accountID,
		Role: role,
	}))
	util.Chkfatal(tx.AddCredential(database.NewCredential{
		AccountID: accountID,
		Login:     true,
		Credential: types.Credential{
			Type:     types.CredentialType(*typ),
			ScopedID: *scopedID,
		},
	}))
	util.Chkfatal(tx.Commit())
}
