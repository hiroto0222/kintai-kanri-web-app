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
