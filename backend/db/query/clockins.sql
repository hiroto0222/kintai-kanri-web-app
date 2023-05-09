-- name: CreateClockIn :one
INSERT INTO "ClockIns" (
  employee_id
) VALUES (
  $1
)
RETURNING *;

-- name: GetClockIn :one
SELECT * FROM "ClockIns"
WHERE "id" = $1;

-- name: ListClockIns :many
SELECT * FROM "ClockIns"
WHERE "employee_id" = $1;

-- name: GetMostRecentClockIn :one
SELECT * FROM "ClockIns"
WHERE "employee_id" = $1
ORDER BY "clock_in_time" DESC
LIMIT 1;

-- name: DeleteClockIn :exec
DELETE FROM "ClockIns"
WHERE "id" = $1;
