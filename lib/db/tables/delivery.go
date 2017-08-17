package tables

var Delivery string = `
CREATE TABLE IF NOT EXISTS delivery (
	delivery_id SERIAL PRIMARY KEY,
	customer_id INTEGER REFERENCES customer,
	location_id INTEGER REFERENCES location,
	fruit_id INTEGER REFERENCES fruit,

	amount NUMERIC(6,2),
	typ INTEGER NOT NULL,
	status INTEGER NOT NULL,

	created_at TIMESTAMP,
	created_where INTEGER NOT NULL,
	created_by INTEGER NOT NULL,
	created_api_time TIMESTAMP
);`

var DeliveryIndex string = `
CREATE INDEX IF NOT EXISTS
delivery_id_index on delivery(delivery_id);`
