package db

import "fmt"

// Fruit is a fruit type
type Fruit struct {
	ID    int
	Name  string
	IsBio bool
}

var (
	fruits []Fruit = []Fruit{
		Fruit{ID: 1, Name: "Äpfel", IsBio: false},
		Fruit{ID: 2, Name: "Bio-Äpfel", IsBio: true},
		Fruit{ID: 3, Name: "Birnen", IsBio: false},
		Fruit{ID: 4, Name: "Bio-Birnen", IsBio: false},
	}
)

// GetFruit returns a fruit instance based on the id
func GetFruit(id int) (Fruit, error) {
	for _, f := range fruits {
		if f.ID == id {
			return f, nil
		}
	}

	return Fruit{}, fmt.Errorf("fruit %d not found", id)
}

func checkFruit(fruit int) bool {
	for _, f := range fruits {
		if f.ID == fruit {
			return true
		}
	}
	return false
}
