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
    deleted_at = ?
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