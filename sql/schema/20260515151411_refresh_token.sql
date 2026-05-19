-- +goose Up
CREATE TABLE refresh_tokens (
    token text PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp CHECK (age(update_at, create_at) >= '0 SECOND'::interval),
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    expires_at timestamp CHECK (age(expires_at, create_at) >= '0 SECOND'::interval),
    revoked_at timestamp CHECK (age(revoked_at, create_at) >= '0 SECOND'::interval)
);

-- +goose Down
DROP TABLE refresh_tokens;
