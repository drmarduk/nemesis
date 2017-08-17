package main

import (
	"database/sql"
	"fmt"
	"log"
)

// createDatabase gets a pure sql connection and
// operates on it, it exists the program on failure
func createDatabase(c *sql.DB, name, owner string) {
	q := fmt.Sprintf(`
		CREATE DATABASE %s
			WITH OWNER=%s
		;`,
		name,
		owner)

	r, err := c.Exec(q)
	if err != nil {
		log.Fatalf("Could not create database %s: %v\n", name, err)
	}

	log.Printf("Database %s created (%d rows affected)\n", name, getRowcount(r))
	return
}

func dropDatabase(c *sql.DB, name string) {
	q := fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, name)
	r, err := c.Exec(q)
	if err != nil {
		log.Fatalf("Could not delete database %s: %v\n", name, err)
	}

	log.Printf("Database %s dropped (%d rows affected)\n", name, getRowcount(r))
}
