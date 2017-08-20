package main

import (
	"database/sql"
	"fmt"
	"log"
)

// grantRights receives a Transaction and tries to
// give the user the rights on a database
// crashes on failure
func grantRights(tx *sql.Tx, db, user string) {
	q := fmt.Sprintf(
		`GRANT ALL PRIVILEGES ON DATABASE %s TO %s;
		GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO %s;`,
		db,
		user,
		user,
	)
	r, err := tx.Exec(q)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Could not grant rights to %s on %s: %v\n", user, db, err)
	}
	log.Printf("Rights granted to %s on %s (%d rows affected)\n", user, db, getRowcount(r))
}

// dropRights removes the rights of a user on a database
// crahes on failure
func dropRights(tx *sql.DB, user, db string) {
	// q := fmt.Sprintf(`REASSIGN OWNED BY "%s" TO "postgres";`, user)
	q := fmt.Sprintf(
		`REVOKE ALL PRIVILEGES ON DATABASE %s FROM %s;`,
		db,
		user,
	)
	r, err := tx.Exec(q)
	if err != nil {
		log.Printf("Warning: Could not delete rights from %s on %s: %v\n", user, db, err)
		return
	}
	log.Printf("Rights removed from %s on %s (%d rows affected)\n", user, db, getRowcount(r))
}
