package database

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"sandstorm.org/go/tempest/internal/common/types"
)

// Test the CredentialAccount method
func TestCredentialAccount(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		addTestData(t, tx)

		id, err := tx.CredentialAccount(types.Credential{
			Type:     "dev",
			ScopedID: "Alice Dev Admin",
		})
		assert.NoError(t, err)
		assert.Equal(t, id, types.AccountID("id_alice"))

		id, err = tx.CredentialAccount(types.Credential{
			"dev",
			"Bob Dev User",
		})
		assert.NoError(t, err)
		assert.Equal(t, id, types.AccountID("id_bob"))
	})
}

func TestUiViews(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		addTestData(t, tx)

		views, err := tx.AccountKeyring("id_alice").AllUiViews()
		assert.NoError(t, err)

		assert.Equal(t, 1, len(views))
		assert.Equal(t, GrainInfo{
			ID:    "grain123",
			Title: "Example Grain",
			Owner: "id_alice",
		}, views[0].Grain)
		assert.Equal(t, 0, len(views[0].Permissions))
	})
}

/*
func TestAccountGrainPermissions(t *testing.T) {
	testWithTx(t, func(tx Tx) {
		addTestData(t, tx)

		expectedPerm := []bool{true, false, true}
		_, err := tx.NewSharingToken("id_bob", "grain123", expectedPerm, "Test share")
		require.NoError(t, err)
		actualPerm, err := tx.AccountGrainPermissions("id_bob", "grain123")
		require.NoError(t, err)
		require.Equal(t, expectedPerm, actualPerm)
	})
}
*/

// addTestData populates the database with some initial data.
func addTestData(t *testing.T, tx Tx) {
	accounts := []NewAccount{
		{
			ID:   "id_alice",
			Role: types.RoleAdmin,
			Profile: Profile{
				DisplayName:     "Alice",
				PreferredHandle: "alice",
			},
		},
		{
			ID:   "id_bob",
			Role: types.RoleUser,
			Profile: Profile{
				DisplayName:     "Bob",
				PreferredHandle: "bob",
			},
		},
	}

	credentials := []NewCredential{
		{
			AccountID: "id_alice",
			Login:     true,
			Credential: types.Credential{
				Type:     "dev",
				ScopedID: "Alice Dev Admin",
			},
		},
		{
			AccountID: "id_bob",
			Login:     true,
			Credential: types.Credential{
				Type:     "dev",
				ScopedID: "Bob Dev User",
			},
		},
	}

	packages := []Package{
		{
			ID: "abcdef",
			// FIXME: Manifest
		},
	}

	grains := []NewGrain{
		{
			GrainID: "grain123",
			PkgID:   "abcdef",
			OwnerID: "id_alice",
			Title:   "Example Grain",
		},
	}

	for _, a := range accounts {
		assert.NoError(t, tx.AddAccount(a), "Adding account: ", a)
	}
	for _, c := range credentials {
		assert.NoError(t, tx.AddCredential(c), "Adding credential: ", c)
	}
	for _, p := range packages {
		assert.NoError(t, tx.AddPackage(p), "Adding package: ", p)
		assert.NoError(t, tx.ReadyPackage(p.ID), "Readying package: ", p)
	}
	for _, g := range grains {
		assert.NoError(t, tx.AddGrain(g), "Adding grain: ", g)
	}
}

// testWithTx runs f in a fresh transaction, then commits. It fails
// the test if the commit fails.
func testWithTx(t *testing.T, f func(tx Tx)) {
	db := openTestDB(t)
	defer db.Close()
	tx, err := db.Begin()
	assert.NoError(t, err)
	defer mustCommit(t, tx)
	f(tx)
}

func mustCommit(t *testing.T, tx Tx) {
	assert.NoError(t, tx.Commit())
}

func openTestDB(t *testing.T) DB {
	// By default we just use an in-memory database, but it can be handy to
	// use an on-disk one so we can inspect it after a test failure to debug:
	dbPath := os.Getenv("TEST_DB_PATH")
	if dbPath == "" {
		dbPath = ":memory:"
	}

	sqlDB, err := sql.Open("sqlite3", dbPath)
	assert.NoError(t, err)
	db, err := InitDB(sqlDB)
	assert.NoError(t, err)
	return db
}
