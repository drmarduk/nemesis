package tables

// Price is a current price for each fruit
var Price string = `
CREATE TABLE IF NOT EXISTS price (
	price_id SERIAL PRIMARY KEY,
	fruit_id INTEGER REFERENCES fruit,
	start_date DATE NOT NULL CHECK(start_date < end_date),
	end_date DATE CHECK(end_date > start_date),
	amount INTEGER NOT NULL,
	is_banned BOOLEAN DEFAULT FALSE,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_where INTEGER,
	created_by INTEGER,
	created_api_time TIMESTAMP
);`

var PriceIndex string = `
CREATE INDEX IF NOT EXISTS
price_id_index on price(price_id);`
