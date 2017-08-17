package db

// Location represents the House where a customre can bring is fruits
type Location struct {
	ID      int
	Name    string
	Surname string
	Addon   string // Zusatz

	Street  string
	Zipcode string // we need the leading zero
	City    string

	Phone string
	Mobil string
	Email string
}

func checkLocation(loc int) bool {
	return true
}
