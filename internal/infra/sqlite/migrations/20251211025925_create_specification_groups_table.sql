-- +goose Up
CREATE TABLE IF NOT EXISTS specification_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    deleted_at TEXT,
    CONSTRAINT unique_public_id
        UNIQUE (public_id)
);

-- +goose Down
DROP TABLE IF EXISTS specification_groups;
