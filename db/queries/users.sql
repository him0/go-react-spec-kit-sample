-- name: GetUserByID :one
SELECT id, name, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, email, created_at, updated_at
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT id, name, email, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CreateUser :exec
INSERT INTO users (id, name, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateUser :exec
UPDATE users
SET name = $1, email = $2, updated_at = $3
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
