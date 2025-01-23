-- name: CreateRoom :exec
INSERT INTO rooms(room_id,room_name,description,room_owner)
VALUES (
    ?,
    ?,
    ?,
    ?
);

-- name: ListRooms :many
SELECT * FROM rooms;

-- name: ListRoom :one
SELECT * FROM rooms
WHERE room_name = ?;