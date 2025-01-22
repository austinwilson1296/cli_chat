-- name: GetUser :one
SELECT * FROM users
WHERE username = ?;

-- name: CreateUser :exec
INSERT INTO users(id,username,password)
VALUES (
    ?,
    ?,
    ?
);