-- name: CreateItemCategory :one
INSERT INTO item_category (
    name,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetItemCategoryByID :one
SELECT *
FROM item_category
WHERE id = $1
LIMIT 1;

-- name: ListItemCategories :many
SELECT *
FROM item_category
ORDER BY name ASC;

-- name: UpdateItemCategory :one
UPDATE item_category
SET
    name = COALESCE($2, name),
    description = COALESCE($3, description)
WHERE id = $1
RETURNING *;

-- name: DeleteItemCategory :exec
DELETE FROM item_category
WHERE id = $1;

