package main

import (
	"database/sql"
	"fmt"
	"log"
)

// createTable receives a Transaction and tries to
// create a new table
func createTable(tx *sql.Tx, name, query string) {
	r, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Could not create table %s: %v\n", name, err)
	}
	log.Printf("Table %s created (%d rows affected)\n", name, getRowcount(r))
}

func dropTable(c *sql.DB, name string) {
	q := fmt.Sprintf("DROP TABLE IF EXISTS %s;", name)
	r, err := c.Exec(q)
	if err != nil {
		log.Fatalf("Could not delete table %s: %v\n", name, err)
	}
	log.Printf("Table %s deleted (%d rows affected)\n", name, getRowcount(r))
}
