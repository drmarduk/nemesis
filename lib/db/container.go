package db

import (
	"fmt"
	"time"
)

// Container is a main place to combine fruits
type Container struct {
	ID        int64
	CreatedAt time.Time
	ClosedAt  time.Time
	Fruit     int
	Status    int
	Location  int
	IsBio     bool
}

/*
	Status:
		- 1 open
		- 2 closed
		- 3 done
*/

// NewContainer creates a new db container
func NewContainer(location, fruit int, bio bool) (*Container, error) {
	if !checkLocation(location) {
		return nil, fmt.Errorf("location is not valid")
	}
	if !checkFruit(fruit) {
		return nil, fmt.Errorf("fruit is not valid")
	}
	query := `
		insert into container(id, created_at, fruit, status, location, bio)
		values(null, ?, ?, ?, ?, ?)`

	c, err := dbSession()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	con := &Container{
		CreatedAt: time.Now().UTC(),
		ClosedAt:  time.Time{}, // should be zero value
		Fruit:     fruit,
		Status:    1,
		Location:  location,
		IsBio:     bio,
	}

	result, err := c.Exec(query, con.CreatedAt, con.Fruit, con.Status, con.Location, con.IsBio)
	if err != nil {
		logger.Printf("error while inserting new container: %v\n", err)
		return nil, err
	}

	con.ID, err = result.LastInsertId()
	if err != nil {
		logger.Printf("error while reading result ID: %v\n", err)
		return nil, err
	}

	return con, nil
}
