-- name: CreateOneProductSpecificationValue :execresult
INSERT INTO product_specifications (
    product_id,
    specification_id,
    string_value,
    int_value,
    bool_value
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
);

-- name: GetAllProductSpecificationValuesByProductID :many
SELECT 
    ps.id,
    ps.product_id,
    ps.specification_id,
    ps.string_value,
    ps.int_value,
    ps.bool_value,
    s.type
FROM product_specifications ps
INNER JOIN specifications s ON s.id = ps.specification_id
WHERE 
    ps.product_id = ?
    AND s.deleted_at IS NULL;