-- name: CreateUser :execresult
INSERT INTO users (hash_id, email, password) VALUES (sqlc.arg(hash_id), sqlc.arg(email), sqlc.arg(password));

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = sqlc.arg(email) LIMIT 1;

-- name: FindUserByHashId :one
SELECT * FROM users
WHERE hash_id = sqlc.arg(hash_id) LIMIT 1;

-- name: FindUser :one
SELECT * FROM users
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET
    email = COALESCE(sqlc.narg("email"), email),
    password = COALESCE(sqlc.narg("password"), password),
    fullname = COALESCE(sqlc.narg("fullname"), fullname)
WHERE id = sqlc.arg(id);

-- name: DeleteUser :exec
DELETE FROM users WHERE id = sqlc.arg(id);
