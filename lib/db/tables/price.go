package tables

var Price string = `
CREATE TABLE IF NOT EXISTS price (
	price_id SERIAL PRIMARY KEY,
	start DATE NOT NULL,
	until DATE,
	value INTEGER NOT NULL
);`

var PriceIndex string = `
CREATE INDEX IF NOT EXISTS
price_id_index on price(price_id);`
