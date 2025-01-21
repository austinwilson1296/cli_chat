-- +goose Up
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    username TEXT,
    password TEXT
);

-- +goose Down
DROP TABLE users;