-- +goose Up

CREATE TABLE user_invitations (
    token TEXT PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expiry TIMESTAMP(0) WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_user_invitations_user_id
ON user_invitations(user_id);

-- +goose Down

DROP TABLE IF EXISTS user_invitations;