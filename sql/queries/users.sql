-- name: CreateUser :one
INSERT INTO users (
    id,
    created_at,
    updated_at,
    user_id,
    email)
VALUES (
    gen_random_uuid (),
    NOW(),
    NOW(),
    $1,
    $2)
RETURNING
    *;

-- name: GetUserInfo :one
SELECT
    *
FROM
    users
WHERE
    user_id = $1;
