-- name: InsertTransactionItemAttribute :one
INSERT INTO transaction_item_attributes (
    id,
    transaction_id,
    attribute_id,
    value
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListTransactionItemAttributes :many
SELECT
    eia.id,
    eia.transaction_id,
    ca.name AS attribute_name,
    ca.data_type,
    eia.value
FROM transaction_item_attributes eia
JOIN category_attributes ca
  ON ca.id = eia.attribute_id
WHERE eia.transaction_id = $1;

-- name: GetEscrowItemAttribute :one
SELECT *
FROM transaction_item_attributes
WHERE transaction_id = $1
  AND attribute_id = $2;

-- name: DeleteTransactionItemAttributes :exec
DELETE FROM transaction_item_attributes
WHERE transaction_id = $1;
