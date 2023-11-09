-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username,
    secret_code,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;