-- name: GetReviews :one
SELECT
    *
FROM
    reviews
WHERE
    review_id = $1;
