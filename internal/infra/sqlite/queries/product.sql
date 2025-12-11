-- name: GetAllProductsByCategoryId :many
SELECT
    p.id,
    p.public_id,
    p.name,
    p.description,
    p.price,
    p.rating,
    p.image_url,
    COUNT(p.id)       OVER () AS products_quantity 
FROM products p
WHERE 
    p.category_id = ?
    AND p.deleted_at IS NULL
LIMIT ? OFFSET ?;

-- name: GetAllProducts :many
SELECT
    p.id,
    p.public_id,
    p.name,
    p.description,
    p.price,
    p.rating,
    p.image_url,
    p.category_id,
    COUNT(p.id)       OVER () AS products_quantity 
FROM products p
WHERE 
    p.deleted_at IS NULL
LIMIT ? OFFSET ?;

-- name: GetOneProductByPublicId :one
SELECT
    p.id,
    p.public_id,
    p.name,
    p.category_id,
    p.description,
    p.price,
    p.rating,
    p.image_url
FROM products p
WHERE 
    p.public_id = ?
    AND p.deleted_at IS NULL
LIMIT 1;

-- name: CreateOneProduct :execresult
INSERT INTO products (
    public_id,
    name,
    description,
    price,
    rating,
    image_url,
    category_id
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

-- name: UpdateOneProduct :exec
UPDATE products
SET
    name = ?,
    description = ?,
    price = ?,
    rating = ?,
    image_url = ?
WHERE
    id = ?;

-- name: DeleteOneProduct :exec
UPDATE products
SET
    deleted_at = (datetime('now'))
WHERE
    id = ?;

-- name: CheckIfProductExists :one
SELECT
    p.id 
FROM products p
WHERE
    p.name = ?
    AND p.deleted_at IS NULL
    AND (
		sqlc.narg ('public_id') IS NULL
		OR p.public_id != sqlc.narg ('public_id')
	)
LIMIT
	1;

-- name: GetOneProductWithSpecificationsByPublicId :many
SELECT 
    p.id AS product_id,
    p.public_id AS product_public_id,
    p.name AS product_name,
    p.description AS product_description,
    p.price AS product_price,
    p.rating AS product_rating,
    p.image_url AS product_image_url,
    p.category_id AS product_category_id,
    sg.id AS specification_group_id,
    sg.public_id AS specification_group_public_id,
    sg.name AS specification_group_name,
    sg.description AS specification_group_description,
    s.id AS specification_id,
    s.public_id AS specification_public_id,
    s.title AS specification_title,
    s.type AS specification_type,
    ps.string_value AS specification_string_value,
    ps.int_value AS specification_int_value,
    ps.bool_value AS specification_bool_value
FROM product_specifications ps
INNER JOIN products p ON p.id = ps.product_id 
INNER JOIN specifications s ON ps.specification_id = s.id
INNER JOIN specification_groups sg ON s.specification_group_id = sg.id
WHERE 
	p.public_id = ?
	AND p.deleted_at IS NULL
	AND sg.deleted_at IS NULL
	AND s.deleted_at IS NULL;