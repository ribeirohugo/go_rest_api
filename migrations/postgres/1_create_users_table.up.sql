CREATE TABLE users (
    id SERIAL,
    username VARCHAR(250),
    email VARCHAR(250) UNIQUE,
    password VARCHAR(250),
    phone VARCHAR(250),
    created TIMESTAMPTZ,
    updated TIMESTAMPTZ
);
