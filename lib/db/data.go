package db

import (
	// github.com/denisenkom/go-mssqldb
	// "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Data is a struct which can be used to establish
// a database connection.
// Host, user, pass and dbname is needed
// which can be passed to GetSession
// to get a cached db connection
type Data struct {
	Host     string
	Username string
	Password string
	Database string
}
