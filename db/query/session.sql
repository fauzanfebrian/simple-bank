-- name: CreateSession :one
INSERT INTO sessions (
    username,
    refresh_token,
    user_agent,
    client_ip,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: ListSessions :many
SELECT * FROM sessions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSession :exec
UPDATE sessions
  set refresh_token = $2,
    is_blocked = $3,
    expires_at = $4
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;