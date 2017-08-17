package main

import (
	"database/sql"
	"fmt"
	"log"
)

// createUser gets a Transaction and tries to add
// a user, crashed on failure
func createUser(tx *sql.Tx, user, pw string, admin bool) {
	_admin := ""
	if admin {
		_admin = "createdb"
	}

	q := fmt.Sprintf(
		`CREATE USER "%s" WITH LOGIN %s PASSWORD '%s'`,
		user,
		_admin,
		pw,
	)
	r, err := tx.Exec(q)
	if err != nil {
		tx.Rollback() // TODO: check return error
		log.Fatalf("Could not create user %s: %v\n", user, err)
	}
	log.Printf("User %s created (%d rows affected)\n", user, getRowcount(r))
}

// dropUser deletes a user
// crashed on failure
func dropUser(tx *sql.Tx, user string) {
	q := fmt.Sprintf(`DROP USER IF EXISTS %s;`, user)
	r, err := tx.Exec(q)
	if err != nil {
		log.Fatalf("Could not delete user %s: %v\n", user, err)
	}

	log.Printf("User %s deleted (%d rows affected)\n", user, getRowcount(r))
}
