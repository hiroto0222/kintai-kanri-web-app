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

-- name: UpdateClockIn :exec
UPDATE "ClockIns"
SET "clocked_out" = $1
WHERE "id" = $2;

-- name: DeleteClockIn :exec
DELETE FROM "ClockIns"
WHERE "id" = $1;

-- name: ListClockInsAndClockOuts :many
SELECT
  ci.id AS clock_in_id,
  ci.employee_id,
  ci.clock_in_time,
  co.id AS clock_out_id,
  co.clock_out_time
FROM "ClockIns" AS ci
  LEFT JOIN "ClockOuts" co
  ON ci.id = co.clock_in_id
WHERE ci.employee_id = $1
ORDER BY ci.clock_in_time DESC;
