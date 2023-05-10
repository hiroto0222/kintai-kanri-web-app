package testutils

import (
	"testing"

	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
)

// CreateTestClockIn はテスト用の ClockIn を作成
func CreateTestClockIn(t *testing.T, employee db.Employee) db.ClockIn {
	return db.ClockIn{
		ID:         0,
		EmployeeID: employee.ID,
	}
}

// CreateTestClockOut はテスト用の ClockOut を作成
func CreateTestClockOut(t *testing.T, employee db.Employee, clockIn db.ClockIn) db.ClockOut {
	return db.ClockOut{
		ID:         0,
		EmployeeID: employee.ID,
		ClockInID:  clockIn.ID,
	}
}
