-- +goose Up
CREATE TABLE users (
    id uuid,
    create_at timestamp NOT NULL,
    update_at timestamp NOT NULL CHECK (age(update_at, create_at) > 0),
    user_id text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    CONSTRAINT users_surrogate_key PRIMARY KEY (id)
);
