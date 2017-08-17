package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/drmarduk/nemesis/lib/db"
	"github.com/drmarduk/nemesis/lib/db/tables"
)

/*
	At least a user with 'create database' priviliges must me on the server
*/

// User holds the user info
type User struct {
	name  string
	pw    string
	admin bool
}

// Table holds the table infos
type Table struct {
	name     string
	create   string
	index    string
	constant string
}

func main() {
	var stage, host, user, password, dbname string
	var reset bool
	flag.StringVar(&stage, "stage", "test", "the stage to deploy: test, prod. Default: test")
	flag.StringVar(&host, "host", "", "our postgres server to install our tables")
	flag.StringVar(&user, "user", "", "a username with rights to install a database")
	flag.StringVar(&password, "password", "", "the password for the given user")
	flag.BoolVar(&reset, "reset", false, "do we want to clean and delete the database and users, default: false")
	flag.Parse()

	if host == "" {
		log.Fatal("no host given")
	}
	if user == "" {
		log.Fatal("no user given")
	}
	dbname = fmt.Sprintf("%snemesis", strings.ToLower(stage))

	log.Printf("Stage:\t %s\n", stage)
	log.Printf("Host:\t %s\n", host)
	log.Printf("Database:\t %s\n", dbname)
	log.Printf("User:\t %s\n", user)
	log.Printf("Password:\t %s\n", password)

	ctxInit := db.Data{
		Host:     host,
		Database: "", // db does not exist yet
		Username: user,
		Password: password,
	}

	ctx := db.Data{
		Host:     host,
		Database: dbname,
		Username: user,
		Password: password,
	}

	users := []User{
		{"nemesis_admin", "nemesis_admin", true},
		{"nemesis_user", "nemesis_user", false},
	}

	tables := []Table{
		{"Customer", tables.Customer, tables.CustomerIndex, ""},
		{"Person", tables.Person, tables.PersonIndex, ""},
		{"Fruit", tables.Fruit, tables.FruitIndex, tables.FruitConstant},
	}

	c, err := db.GetSession(ctxInit)
	if err != nil {
		log.Fatalf("could not connect to host %s: %v\n", ctx.Host, err)
	}

	if reset {
		resetDB(c, dbname, users)
	} else {
		createDatabase(c, dbname, user)
		// database Created, we need a new connection with correct dsn with database
		err = db.ResetSession()
		if err != nil {
			log.Fatalf("error while resetting session: %v\n", err)
		}
		log.Println("Database connection closed, try to reconnect.")

		c, err = db.GetSession(ctx)
		if err != nil {
			log.Fatalf("could not connect to host %s: %v\n", ctx.Host, err)
		}
		err = c.Ping()
		if err != nil {
			log.Fatalf("could not ping to database: %v\n", err)
		}
		log.Println("Database connection reestablished.")
		setupDB(c, dbname, users, tables)
	}
}

func insertConstant(tx *sql.Tx, name, query string) {
	if query == "" {
		return
	}
	r, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Could not insert values in %s: %v\n", name, err)
	}
	log.Printf("Constants inserted in %s (%d rows affected)\n", name, getRowcount(r))
}
