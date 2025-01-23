-- +goose Up
CREATE TABLE rooms(
    room_id TEXT PRIMARY KEY,
    room_name TEXT NOT NULL,
    description TEXT,
    room_owner TEXT NOT NULL,
    FOREIGN KEY(room_owner) REFERENCES users(id)

);

-- +goose Down
DROP TABLE rooms;