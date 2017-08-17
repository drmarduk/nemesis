package db

import (
	"database/sql"
	"fmt"
)

var (
	session *sql.DB
)

// GetSession creates a database instance with the
// given ctx Information. It tries to establish a database
// connection, caches it and returns it to the user
// PostgreSQL is hardcoded!!!! keep in mind
// with ssl enabled
func GetSession(ctx Data) (*sql.DB, error) {
	if session != nil {
		return session, nil
	}
	c, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=require",
			ctx.Username,
			ctx.Password,
			ctx.Host,
			ctx.Database,
		),
	)
	if err != nil {
		return nil, err
	}
	// cache it
	session = c
	return session, nil
}

// ResetSession closes the connection, so that GetSession creates a new session
func ResetSession() error {
	err := session.Close()
	if err != nil {
		return err
	}
	session = nil
	return nil
}
