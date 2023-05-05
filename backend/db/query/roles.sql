-- name: CreateRole :one
INSERT INTO "Roles" (name) VALUES ($1)
RETURNING *;

-- name: GetRoleByID :one
SELECT * FROM "Roles"
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM "Roles"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteRole :exec
DELETE FROM "Roles"
WHERE id = $1;
