package backend

import "fmt"

// Customer is a customer who is allowed to bring stuff
type Customer struct {
	ID      int
	Name    string
	Surname string

	Street  string
	Zipcode string
	City    string
}

// NewCustomer asdf
func NewCustomer(name, surname, street, zipcode, city string) (*Customer, error) {
	return nil, fmt.Errorf("not yet implemented")
}
