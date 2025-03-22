-- name: CreateUser :execresult
INSERT INTO users (hash_id, email, password) VALUES (?, ?, ?);

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: FindUserByHashId :one
SELECT * FROM users
WHERE hash_id = ? LIMIT 1;

-- name: FindUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: UpdatePassword :exec
UPDATE users
SET password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
