-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    id,
    token,
    created_at,
    updated_at,
    expires_at)
VALUES (
    $1,
    $2,
    NOW(),
    NOW(),
    NOW() + interval '60 days')
RETURNING
    *;
