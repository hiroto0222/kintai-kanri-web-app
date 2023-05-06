-- name: CreateEmployee :one
INSERT INTO "Employees" (
  first_name,
  last_name,
  email,
  phone,
  address,
  hashed_password,
  role_id,
  is_admin
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetEmployeeById :one
SELECT * FROM "Employees"
WHERE id = $1 LIMIT 1;

-- name: GetEmployeeByEmail :one
SELECT * FROM "Employees"
WHERE email = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM "Employees"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteEmployee :exec
DELETE FROM "Employees"
WHERE id = $1;
