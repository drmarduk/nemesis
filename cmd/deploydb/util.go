package main

import (
	"database/sql"
	"log"
)

func getRowcount(r sql.Result) int {
	c, err := r.RowsAffected()
	if err != nil {
		log.Printf("Could not get affected rowcount: %v\n", err)
		return -1
	}
	return int(c)
}
