-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
  set hashed_password = $2,
    is_email_verified = $3,
    password_changed_at = $4
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;