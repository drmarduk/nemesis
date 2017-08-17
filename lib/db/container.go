package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Container is a main place to combine fruits
type Container struct {
	ID     int64
	Fruit  int
	Status int
	IsBio  bool

	CreatedAt      time.Time
	CreatedWhere   int
	CreatedBy      int
	CreatedAPITime time.Time

	ClosedAt      time.Time
	ClosedWhere   int
	ClosedBy      int
	ClosedAPITime time.Time
}

/*
	Status:
		- 1 open
		- 2 closed
		- 3 done
*/

// NewContainer creates a new db container
func NewContainer(c *sql.DB, location, fruit int, bio bool) (*Container, error) {
	if !checkLocation(location) {
		return nil, fmt.Errorf("location is not valid")
	}
	if !checkFruit(fruit) {
		return nil, fmt.Errorf("fruit is not valid")
	}
	query := `
		insert into container(id, created_at, fruit, status, location, bio)
		values(null, ?, ?, ?, ?, ?)`

	con := &Container{
		CreatedAt:    time.Now().UTC(),
		ClosedAt:     time.Time{}, // should be zero value
		Fruit:        fruit,
		Status:       1,
		CreatedWhere: location,
		IsBio:        bio,
	}

	result, err := c.Exec(query, con.CreatedAt, con.Fruit, con.Status, con.CreatedWhere, con.IsBio)
	if err != nil {
		log.Printf("error while inserting new container: %v\n", err)
		return nil, err
	}

	con.ID, err = result.LastInsertId()
	if err != nil {
		log.Printf("error while reading result ID: %v\n", err)
		return nil, err
	}

	return con, nil
}
