-- name: CreateClockOut :one
INSERT INTO "ClockOuts" (
  employee_id,
  clock_in_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetClockOut :one
SELECT * FROM "ClockOuts"
WHERE "id" = $1;

-- name: ListClockOuts :many
SELECT * FROM "ClockOuts"
WHERE "employee_id" = $1;
