package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Price is a price for a fruit based on the year
// Amount is in Cents!!!
type Price struct {
	ID       int
	Fruit    int
	Start    time.Time // start_date
	End      time.Time // end_date
	Amount   int
	IsBanned bool

	CreatedAt      time.Time
	CreatedWhere   int
	CreatedBy      int
	CreatedAPITime time.Time
}

// NewPrice inserts a new price for a fruit in the database.
// it must take a sql.connection as parameter
func NewPrice(c *sql.DB, fruitID, amount, where, by int, start, end time.Time) (Price, error) {
	var p Price
	q := fmt.Sprintf(`
		INSERT INTO price(
			fruit_id,
			start_date,
			end_date,
			amount,
			created_where,
			created_by
		) VALUES($1, $2, $3, $4, $5, $6);`)
	r, err := c.Exec(
		q,
		fruitID,
		start,
		end,
		amount,
		where,
		by,
	)
	if err != nil {
		return p, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return p, err
	}

	// GetPrice for this id
	p, err = GetPrice(c, int(id))
	return p, err
}

// GetPrice returns the price by ID
func GetPrice(c *sql.DB, priceID int) (Price, error) {
	q :=
		`SELECT 
			price_id,
			fruit_id, 
			start_date, 
			end_date,
			value,
			created_at,
			created_where,
			created_by,
			created_api_time
		FROM
			price
		where
			price_id = ?
		ORDER BY
			start_date DESC;
		`
	row := c.QueryRow(q, priceID)
	var r Price

	err := row.Scan(
		&r.ID,
		&r.Fruit,
		&r.Start,
		&r.End,
		&r.Amount,
		&r.CreatedAt,
		&r.CreatedWhere,
		&r.CreatedBy,
		&r.CreatedAPITime,
	)
	if err != nil {
		return r, err
	}
	return r, nil
}
