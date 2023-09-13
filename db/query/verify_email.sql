-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE id = $1 LIMIT 1;

-- name: ListVerifyEmails :many
SELECT * FROM verify_emails
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateVerifyEmail :exec
UPDATE verify_emails
  set secret_code = $2,
    is_used = $3,
    expired_at = $4
WHERE id = $1;

-- name: DeleteVerifyEmail :exec
DELETE FROM verify_emails
WHERE id = $1;