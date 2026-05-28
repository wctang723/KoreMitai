-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    create_at timestamp NOT NULL,
    update_at timestamp NOT NULL CHECK (age(update_at, create_at) >= '0 SECOND'::interval),
    user_id text UNIQUE NOT NULL,
    email text UNIQUE NOT NULL,
    hashed_password text NOT NULL
);

-- +goose Down
DROP TABLE users;
