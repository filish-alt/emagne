-- name: CreateCategoryAttribute :one
INSERT INTO category_attributes (
    id,
    category_id,
    name,
    data_type,
    is_required
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: ListCategoryAttributes :many
SELECT *
FROM category_attributes
WHERE category_id = $1
ORDER BY name;

-- name: GetCategoryAttribute :one
SELECT *
FROM category_attributes
WHERE category_id = $1
  AND name = $2;

-- name: DeleteCategoryAttribute :exec
DELETE FROM category_attributes
WHERE id = $1;
