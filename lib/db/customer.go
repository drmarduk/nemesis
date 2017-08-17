package db

import (
	"errors"
	"time"
)

// Customer represents the database table of a customer
type Customer struct {
	ID        int
	PriapusID int

	Name    string
	Surname string
	Addon   string // Zusatz

	Street  string
	Zipcode string // we need the leading zero
	City    string

	Phone string
	Mobil string
	Mail  string

	Group      int
	IsBio      bool
	IsBanned   bool
	CanDeliver bool

	CreatedAt      time.Time
	CreatedWhere   int
	CreatedBy      int
	CreatedAPITime time.Time
}

// NewCustomer is called only once for each customer
func NewCustomer(name, surname, addon, street, zipcode, city, phone, mobil, mail string, location int) (*Customer, error) {
	if name == "" && surname == "" {
		return nil, errors.New("name and surname must be set")
	}
	if location == 0 {
		return nil, errors.New("location must not be zero")
	}

	c := &Customer{
		Name:    name,
		Surname: surname,

		Street:  street,
		Zipcode: zipcode,
		City:    city,

		Phone: phone,
		Mobil: mobil,
		Mail:  mail,

		// CreatedLocation: ,

		IsBio:      false,
		IsBanned:   false,
		CanDeliver: true,
	}
	// TODO: insert into database and set ID
	return c, nil
}
