-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id CHAR(8) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(2000) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    deleted_at TIMESTAMP,
    CONSTRAINT unique_public_id
        UNIQUE (public_id)
);

-- +goose Down
DROP TABLE IF EXISTS categories;

