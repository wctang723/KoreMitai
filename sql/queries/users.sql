-- name: CreateUser :one
INSERT INTO users (
    id,
    create_at,
    update_at,
    user_id,
    email,
    hashed_password)
VALUES (
    gen_random_uuid (),
    NOW(),
    NOW(),
    $1,
    $2,
    $3)
RETURNING
    *;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1;
