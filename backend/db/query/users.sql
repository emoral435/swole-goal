-- name: CreateUser :one
INSERT INTO "users" (
    email, password, username
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUsers :one
SELECT * FROM "users"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "users"
ORDER BY id
LIMIT
    1
    OFFSET 2;

-- name: UpdateUsers :one
UPDATE "users"
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM "users"
WHERE id = $1;
