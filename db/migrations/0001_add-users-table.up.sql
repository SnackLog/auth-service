CREATE TABLE users (
    username VARCHAR(50) PRIMARY KEY,
    display_name VARCHAR(100) NOT NULL,
    salt VARCHAR(256) NOT NULL,
    password_hash VARCHAR(1024) NOT NULL
);


