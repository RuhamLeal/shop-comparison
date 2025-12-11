-- name: GetOneCategoryByPublicID :one
SELECT 
    c.id,
    c.public_id,
    c.name,
    c.description
FROM categories c
WHERE 
    c.public_id = ?
    AND c.deleted_at IS NULL
LIMIT 1;

-- name: GetAllCategories :many
SELECT 
    c.id,
    c.public_id,
    c.name,
    c.description
FROM categories c
WHERE 
    c.deleted_at IS NULL;   