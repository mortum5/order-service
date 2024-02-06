-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY id DESC LIMIT $1;

-- name: CreateOrder :one
INSERT INTO orders (
    id,
    data
) VALUES (
    $1,
    $2
) RETURNING *;
