-- Test db 
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    password text NOT NULL,
    email text NOT NULL
);
