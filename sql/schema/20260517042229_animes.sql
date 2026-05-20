-- +goose Up
CREATE TABLE animes (
    animes_id uuid,
    created_at timestamp,
    updated_at timestamp,
    title text NOT NULL,
    PRIMARY KEY (animes_id)
);

-- +goose Down
DROP TABLE animes;
