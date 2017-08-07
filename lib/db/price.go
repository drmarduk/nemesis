package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Price is a price for a fruit based on the year
type Price struct {
	gorm.Model

	Fruit  Fruit
	Year   time.Time
	Amount float64
}
