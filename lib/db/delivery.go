package db

import (
	"time"
)

// Delivery is a process of bringing fruits in a container
type Delivery struct {
	ID         int
	CustomerID int
	LocationID int
	FruitID    int
	Amount     float64

	Typ    int // normal, Löschen, Korrektur
	Status int

	CreatedAt      time.Time
	CreatedWhere   int
	CreatedBy      int
	CreatedAPITime time.Time
}
