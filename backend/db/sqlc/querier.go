// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateClockIn(ctx context.Context, employeeID uuid.UUID) (ClockIn, error)
	CreateClockOut(ctx context.Context, arg CreateClockOutParams) (ClockOut, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error)
	CreateRole(ctx context.Context, name string) (Role, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	DeleteClockIn(ctx context.Context, id int32) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
	DeleteRole(ctx context.Context, id int32) error
	GetClockIn(ctx context.Context, id int32) (ClockIn, error)
	GetClockOut(ctx context.Context, id int32) (ClockOut, error)
	GetEmployeeByEmail(ctx context.Context, email string) (Employee, error)
	GetEmployeeById(ctx context.Context, id uuid.UUID) (Employee, error)
	GetMostRecentClockIn(ctx context.Context, employeeID uuid.UUID) (ClockIn, error)
	GetRoleByID(ctx context.Context, id int32) (Role, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	ListClockIns(ctx context.Context, employeeID uuid.UUID) ([]ClockIn, error)
	ListClockInsAndClockOuts(ctx context.Context, employeeID uuid.UUID) ([]ListClockInsAndClockOutsRow, error)
	ListClockOuts(ctx context.Context, employeeID uuid.UUID) ([]ClockOut, error)
	ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employee, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	UpdateClockIn(ctx context.Context, arg UpdateClockInParams) error
}

var _ Querier = (*Queries)(nil)
