-- +goose Up
CREATE TABLE animes (
    animes_id uuid,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    title text NOT NULL,
    PRIMARY KEY (animes_id)
);

-- +goose Down
DROP TABLE animes;
