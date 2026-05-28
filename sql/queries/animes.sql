-- name: GetAnimes :one
SELECT
    *
FROM
    animes
WHERE
    animes_id = $1;
