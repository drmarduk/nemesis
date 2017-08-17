package tables

var Location string = `
CREATE TABLE IF NOT EXISTS location (
	location_id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	surname TEXT,
	addon TEXT,

	street TEXT,
	zipcode TEXT,
	city TEXT,

	phone TEXT,
	mobile TEXT,
	email TEXT
);`

var LocationIndex string = `
CREATE INDEX IF NOT EXISTS
location_id_index on location(location_id);
`
