-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    image_url TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    deleted_at TEXT,
    CONSTRAINT unique_public_id
        UNIQUE (public_id),
    CONSTRAINT product_category_fk_1
        FOREIGN KEY (category_id) REFERENCES categories (id)
);


-- +goose Down
DROP TABLE IF EXISTS products;
