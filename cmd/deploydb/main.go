package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	var stage, host, user, password, dbname string
	flag.StringVar(&stage, "stage", "test", "the stage to deploy: test, prod. Default: test")
	flag.StringVar(&host, "host", "", "our postgres server to install our tables")
	flag.StringVar(&user, "user", "", "a username with rights to install a database")
	flag.StringVar(&password, "password", "", "the password for the given user")
	flag.Parse()

	if host == "" {
		log.Fatal("no host given")
	} else if user == "" {
		log.Fatal("no user given")
	}
	dbname = fmt.Sprintf("%sNemesis", strings.ToUpper(stage))

	log.Printf("Deploy stage: %s\n", stage)
	log.Printf("Host: %s\n", host)
	log.Printf("Database: %s\n", dbname)
	log.Printf("User: %s\n", user)
	log.Printf("Password: %s\n", password)
}
