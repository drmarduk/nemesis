package tables

// Fruit is the fruit table
var Fruit string = `
CREATE TABLE IF NOT EXISTS fruit (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	is_bio BOOLEAN
);
`

// FruitIndex is a index
var FruitIndex string = `
create index if not exists
fruit_id_index on "Fruit"(id);
`

// FruitConstant is a collection of constant entries
var FruitConstant string = `
insert into fruit(id, name, is_bio) values
	(1, 'Äpfel', false),
	(2, 'Bio-Äpfel', true),
	(3, 'Birnen', false),
	(4, 'Bio-Birnen', true);
`
