package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func init() {
	// we check in this func if all tables are registered in the database

	db := gorm.DB{}
	gorm.Open("sqlite3", "")
	// TODO: open db
	check(db.HasTable(&Customer{}), "customer")

}

func check(valid bool, name string) {
	if !valid {
		panic(fmt.Errorf("table %s does not exist", name))
	}
}
