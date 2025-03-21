-- name: CreateUser :execresult
INSERT INTO users (username, password) VALUES (?, ?);

-- name: FindUserByUsername :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: FindUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET username = ?, password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
