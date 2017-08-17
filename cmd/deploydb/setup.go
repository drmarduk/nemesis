package main

import (
	"database/sql"
	"log"
)

func setupDB(c *sql.DB, dbname string, users []User, tables []Table) {
	// start transaction
	t, err := c.Begin()
	if err != nil {
		c.Close()
		log.Fatalf("error while creating transaction: %v\n", err)
	}

	// Create User and grant rights
	for _, u := range users {
		createUser(t, u.name, u.pw, u.admin)
		grantRights(t, dbname, u.name)
	}

	// create tables with indexes
	for _, tbl := range tables {
		createTable(t, tbl.name, tbl.create)
		insertConstant(t, tbl.name, tbl.constant)
	}

	err = t.Commit()
	if err != nil {
		log.Fatalf("error while commiting transaction: %v\n", err)
	}
	log.Print("Setup done.")
}
