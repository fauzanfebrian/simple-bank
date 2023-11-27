-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email, role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password),hashed_password),
  full_name = COALESCE(sqlc.narg(full_name),full_name),
  email = COALESCE(sqlc.narg(email),email),
  password_changed_at = CASE WHEN sqlc.narg(hashed_password) IS NOT NULL THEN NOW() ELSE password_changed_at END,
  is_email_verified = COALESCE(sqlc.narg(is_email_verified),is_email_verified)
WHERE
  username = @username
RETURNING *;