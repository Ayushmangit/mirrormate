-- +goose Up

ALTER TABLE users
ADD COLUMN updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW();

-- +goose Down

ALTER TABLE users
DROP COLUMN updated_at;
