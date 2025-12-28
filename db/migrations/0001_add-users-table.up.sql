CREATE TABLE users (
    username VARCHAR(50) PRIMARY KEY,
    display_name VARCHAR(100) NOT NULL,
    salt VARCHAR(32) NOT NULL,
    password_hash VARCHAR(64) NOT NULL
);


