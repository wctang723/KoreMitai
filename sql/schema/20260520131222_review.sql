-- +goose Up
CREATE TABLE reviews (
    review_id uuid,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    star int NOT NULL CHECK (star >= 0 AND star <= 10),
    body text,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    animes_id uuid REFERENCES animes (animes_id) ON DELETE RESTRICT,
    PRIMARY KEY (review_id, user_id, animes_id)
);

-- +goose Down
DROP TABLE reviews;
