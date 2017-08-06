package db

import "github.com/jinzhu/gorm"

// Location represents the House where a customre can bring is fruits
type Location struct {
	gorm.Model
	Name    string
	Surname string
}
