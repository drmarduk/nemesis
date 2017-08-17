package db

import (
	"time"
)

// Price is a price for a fruit based on the year
type Price struct {
	Fruit  int
	Start  time.Time
	End    time.Time
	Amount float64

	CreatedAt      time.Time
	CreatedWhere   int
	CreatedBy      int
	CreatedAPITime time.Time
}
