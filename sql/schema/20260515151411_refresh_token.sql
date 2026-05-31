-- +goose Up
CREATE TABLE refresh_tokens (
    id uuid REFERENCES users (id) ON DELETE CASCADE,
    token text UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL CHECK (age(updated_at, created_at) >= '0 SECOND'::interval),
    expires_at timestamp NOT NULL CHECK (age(expires_at, created_at) >= '0 SECOND'::interval),
    revoked_at timestamp DEFAULT NULL CHECK (age(revoked_at, created_at) >= '0 SECOND'::interval),
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE refresh_tokens;
