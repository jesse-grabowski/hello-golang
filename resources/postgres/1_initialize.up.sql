CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    username varchar(256) UNIQUE NOT NULL,
    password bytea NOT NULL
);