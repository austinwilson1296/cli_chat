-- +goose Up
CREATE TABLE users(
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;