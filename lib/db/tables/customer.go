package tables

// Customer is the customer table
var Customer string = `
/* create table customer */
create table if not exists customer(
    id serial primary key,
    priapus_id int not null unique,
    
    name text not null,
    surname text not null,
    addon text,

    street text,
    zipcode text, /* cause leading zero */
    city text,

    phone text,
    mobil text,
    mail text,

    customer_group int,
    is_bio boolean,
    is_banned boolean,
    candeliver boolean,

    created_at timestamp,
    created_where int,
    created_by int, /* later references to user table */
    created_at_api_time timestamp
);
`

var CustomerIndex string = `
    create unique index if not exists
    customer_id_index on 'Customer'(id);

    create unique index if not exists
    customer_priapusid_index on 'Customer'(priapus_id);
`
