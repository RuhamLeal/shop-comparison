-- +goose Up
CREATE TABLE IF NOT EXISTS specifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id TEXT NOT NULL,
    specification_group_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    type TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    deleted_at TEXT,
    CONSTRAINT unique_public_id
        UNIQUE (public_id),
    CONSTRAINT specification_group_fk_1
        FOREIGN KEY (specification_group_id) REFERENCES specification_groups (id)
);

-- +goose Down
DROP TABLE IF EXISTS specifications;

