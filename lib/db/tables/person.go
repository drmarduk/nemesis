package tables

var Person string = `
/* create table user */
create table if not exists person (
    id serial primary key,
    login_name text not null,
    password text not null,

    created_at timestamp,
    created_where int,
    created_by int,
    created_at_api_time timestamp
);
`

var PersonIndex string = `
    create index if not exists
    person_id_index on person;
`
