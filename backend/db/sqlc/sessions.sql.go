// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: sessions.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
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
) RETURNING id, email, employee_id, refresh_token, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	EmployeeID   uuid.UUID `json:"employee_id"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.Email,
		arg.EmployeeID,
		arg.RefreshToken,
		arg.IsBlocked,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmployeeID,
		&i.RefreshToken,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, email, employee_id, refresh_token, is_blocked, expires_at, created_at FROM "Sessions"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmployeeID,
		&i.RefreshToken,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
