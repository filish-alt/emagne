-- name: CreateTransaction :one
INSERT INTO transactions (
    title,
    role,
    currency,
    inspection_period,
    item_category_id,
    item_name,
    item_description,
    price,
    shipping_method,
    seller_email,
    seller_phone,
    buyer_email,
    buyer_phone,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: GetTransactionWithCategory :one
SELECT
    t.*,
    ic.name AS item_category_name,
    ic.description AS item_category_description
FROM transactions t
JOIN item_categories ic ON ic.id = t.item_category_id
WHERE t.id = $1
LIMIT 1;

-- name: ListTransactionsByCategory :many
SELECT
    t.*,
    ic.name AS item_category_name,
    ic.description AS item_category_description
FROM transactions t
JOIN item_categories ic ON ic.id = t.item_category_id
WHERE t.item_category_id = $1
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateTransactionStatus :one
UPDATE transactions
SET
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;
