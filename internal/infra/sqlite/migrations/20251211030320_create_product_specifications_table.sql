-- +goose Up
CREATE TABLE IF NOT EXISTS product_specifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER NOT NULL,
    specification_id INTEGER NOT NULL,
    string_value TEXT,
    int_value INTEGER,
    bool_value INTEGER,
    UNIQUE (product_id, specification_id),
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (specification_id) REFERENCES specifications (id)
);

-- +goose Down
DROP TABLE IF EXISTS product_specifications;
