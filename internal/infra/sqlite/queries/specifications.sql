-- name: GetAllSpecifications :many
SELECT 
    s.id,
    s.public_id,
    s.title,
    s.type
FROM 
    specifications s
WHERE 
    s.specification_group_id = ?
    AND s.deleted_at IS NULL;

-- name: GetOneSpecificationByPublicID :one
SELECT 
    s.id,
    s.public_id,
    s.title,
    s.type,
    sg.id
FROM specifications s
INNER JOIN specification_groups sg ON s.specification_group_id = sg.id
WHERE 
    s.public_id = ?
    AND sg.deleted_at IS NULL
    AND s.deleted_at IS NULL
LIMIT 1;