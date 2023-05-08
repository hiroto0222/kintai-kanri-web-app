-- name: CreateSession :one
INSERT INTO "Sessions" (
  id,
  email,
  employee_id,
  refresh_token,
  is_blocked,
  expires_at,
  created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM "Sessions"
WHERE id = $1 LIMIT 1;
