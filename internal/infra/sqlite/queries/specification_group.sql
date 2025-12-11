-- name: GetAllSpecificationGroups :many
SELECT 
    sg.id,
    sg.public_id,
    sg.name,
    sg.description
FROM 
    specification_groups sg
WHERE 
    sg.deleted_at IS NULL;

-- name: GetOneSpecificationGroupByPublicID :one
SELECT 
    sg.id,
    sg.public_id,
    sg.name,
    sg.description
FROM specification_groups sg
WHERE 
    sg.public_id = ?
    AND sg.deleted_at IS NULL
LIMIT 1;