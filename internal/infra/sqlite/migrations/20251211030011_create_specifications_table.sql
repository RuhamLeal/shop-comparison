-- +goose Up
CREATE TABLE IF NOT EXISTS specifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id CHAR(8) NOT NULL,
    specification_group_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    deleted_at TIMESTAMP,
    CONSTRAINT unique_public_id
        UNIQUE (public_id),
    CONSTRAINT specification_group_fk_1
        FOREIGN KEY (specification_group_id) REFERENCES specification_groups (id)
);

-- +goose Down
DROP TABLE IF EXISTS specifications;

