package db

import "github.com/jinzhu/gorm"

// Delivery is a process of bringing fruits in a container
type Delivery struct {
	gorm.Model

	Customer  Customer
	Container Container
	Location  Location // we might not need this cause of Container
	Fruit     Fruit
	Status    int // dunno
}
