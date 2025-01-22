// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users(id,username,password)
VALUES (
    ?,
    ?,
    ?
)
`

type CreateUserParams struct {
	ID       interface{}
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.ID, arg.Username, arg.Password)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password FROM users
WHERE username = ?
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}
