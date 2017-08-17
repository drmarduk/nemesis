package main

import (
	"database/sql"
	"log"
)

// resetDB receives a database connection and deletes
// all user, rights and the database
func resetDB(c *sql.DB, db string, users []User) {
	t, err := c.Begin()
	if err != nil {
		log.Fatalf("Could not create transaction: %v\n", err)
	}

	for _, u := range users {
		dropRights(t, u.name, db)
		dropUser(t, u.name)
	}

	err = t.Commit()
	if err != nil {
		log.Fatalf("Could not commit transaction: %v\n", err)
	}
	dropDatabase(c, db)
	log.Printf("Reset done.")
}
