-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    public_id CHAR(8) NOT NULL,
    category_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(2000) NULL,
    price INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    image_url TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    updated_at TIMESTAMP NOT NULL DEFAULT (datetime('now')),
    deleted_at TIMESTAMP,
    CONSTRAINT unique_public_id
        UNIQUE (public_id),
    CONSTRAINT product_category_fk_1
        FOREIGN KEY (category_id) REFERENCES categories (id)
);


-- +goose Down
DROP TABLE IF EXISTS products;
