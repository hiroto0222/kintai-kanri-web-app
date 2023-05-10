// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: clockins.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createClockIn = `-- name: CreateClockIn :one
INSERT INTO "ClockIns" (
  employee_id
) VALUES (
  $1
)
RETURNING id, employee_id, clocked_out, clock_in_time
`

func (q *Queries) CreateClockIn(ctx context.Context, employeeID uuid.UUID) (ClockIn, error) {
	row := q.db.QueryRowContext(ctx, createClockIn, employeeID)
	var i ClockIn
	err := row.Scan(
		&i.ID,
		&i.EmployeeID,
		&i.ClockedOut,
		&i.ClockInTime,
	)
	return i, err
}

const deleteClockIn = `-- name: DeleteClockIn :exec
DELETE FROM "ClockIns"
WHERE "id" = $1
`

func (q *Queries) DeleteClockIn(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteClockIn, id)
	return err
}

const getClockIn = `-- name: GetClockIn :one
SELECT id, employee_id, clocked_out, clock_in_time FROM "ClockIns"
WHERE "id" = $1
`

func (q *Queries) GetClockIn(ctx context.Context, id int32) (ClockIn, error) {
	row := q.db.QueryRowContext(ctx, getClockIn, id)
	var i ClockIn
	err := row.Scan(
		&i.ID,
		&i.EmployeeID,
		&i.ClockedOut,
		&i.ClockInTime,
	)
	return i, err
}

const getMostRecentClockIn = `-- name: GetMostRecentClockIn :one
SELECT id, employee_id, clocked_out, clock_in_time FROM "ClockIns"
WHERE "employee_id" = $1
ORDER BY "clock_in_time" DESC
LIMIT 1
`

func (q *Queries) GetMostRecentClockIn(ctx context.Context, employeeID uuid.UUID) (ClockIn, error) {
	row := q.db.QueryRowContext(ctx, getMostRecentClockIn, employeeID)
	var i ClockIn
	err := row.Scan(
		&i.ID,
		&i.EmployeeID,
		&i.ClockedOut,
		&i.ClockInTime,
	)
	return i, err
}

const listClockIns = `-- name: ListClockIns :many
SELECT id, employee_id, clocked_out, clock_in_time FROM "ClockIns"
WHERE "employee_id" = $1
`

func (q *Queries) ListClockIns(ctx context.Context, employeeID uuid.UUID) ([]ClockIn, error) {
	rows, err := q.db.QueryContext(ctx, listClockIns, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClockIn
	for rows.Next() {
		var i ClockIn
		if err := rows.Scan(
			&i.ID,
			&i.EmployeeID,
			&i.ClockedOut,
			&i.ClockInTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listClockInsAndClockOuts = `-- name: ListClockInsAndClockOuts :many
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
ORDER BY ci.clock_in_time DESC
`

type ListClockInsAndClockOutsRow struct {
	ClockInID    int32         `json:"clock_in_id"`
	EmployeeID   uuid.UUID     `json:"employee_id"`
	ClockInTime  time.Time     `json:"clock_in_time"`
	ClockOutID   sql.NullInt32 `json:"clock_out_id"`
	ClockOutTime sql.NullTime  `json:"clock_out_time"`
}

func (q *Queries) ListClockInsAndClockOuts(ctx context.Context, employeeID uuid.UUID) ([]ListClockInsAndClockOutsRow, error) {
	rows, err := q.db.QueryContext(ctx, listClockInsAndClockOuts, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListClockInsAndClockOutsRow
	for rows.Next() {
		var i ListClockInsAndClockOutsRow
		if err := rows.Scan(
			&i.ClockInID,
			&i.EmployeeID,
			&i.ClockInTime,
			&i.ClockOutID,
			&i.ClockOutTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClockIn = `-- name: UpdateClockIn :exec
UPDATE "ClockIns"
SET "clocked_out" = $1
WHERE "id" = $2
`

type UpdateClockInParams struct {
	ClockedOut bool  `json:"clocked_out"`
	ID         int32 `json:"id"`
}

func (q *Queries) UpdateClockIn(ctx context.Context, arg UpdateClockInParams) error {
	_, err := q.db.ExecContext(ctx, updateClockIn, arg.ClockedOut, arg.ID)
	return err
}
