-- +goose Up
ALTER TABLE IF EXISTS users ADD COLUMN role_id INT REFERENCES roles(id) DEFAULT 1;

-- +goose Down
ALTER TABLE IF EXISTS users DROP COLUMN role_id;
