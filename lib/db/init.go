package db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	conn    *sql.DB
	context Data
	logger  *log.Logger
)

// Data holds the datbase metadata
type Data struct {
	Driver   string
	Host     string
	Username string
	Password string
	Database string
}

func dbSession() (*sql.DB, error) {
	// TODO: change to either mssql, mysql or psql
	var err error
	conn, err = sql.Open(
		context.Driver,
		fmt.Sprintf("file:data.db"),
	)

	return conn, err
}

// InitializeDatabase creates all necessary connections
func InitializeDatabase(ctx Data) error {
	context = ctx
	_, err := dbSession()
	return err
}
