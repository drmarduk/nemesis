package tables

// Container ss
var Container string = `
CREATE TABLE IF NOT EXISTS container (
	container_id SERIAL PRIMARY KEY,
	fruit_id INTEGER NOT NULL,
	hash INTEGER NOT NULL,
	status INTEGER DEFAULT 1,
	is_bio boolean NOT NULL,

	created_at TIMESTAMP,
	created_where INTEGER,
	created_by INTEGER,
	created_api_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	closed_at TIMESTAMP,
	closed_where INTEGER,
	closed_by INTEGER,
	closed_api_time TIMESTAMP
);`

var ContainerIndex string = `
CREATE INDEX IF NOT EXISTS 
container_id_index on container(id);
`
