-- +goose Up
CREATE TABLE users_review (
    review_id uuid,
    created_at timestamp,
    updated_at timestamp,
    star int NOT NULL CHECK (star >= 0 AND star <= 10),
    body text NOT NULL,
    user_id uuid REFERENCES users (user_id) ON DELETE CASCADE,
    animes_id uuid REFERENCES animes (animes_id) ON DELETE RESTRICT,
    PRIMARY KEY (review_id, user_id, animes_id)
);

-- +goose Down
DROP TABLE users_chirpy;
