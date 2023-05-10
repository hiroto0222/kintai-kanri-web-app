package db_test

import (
	"testing"

	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
)

func TestCreateClockIn(t *testing.T) {
	employee := CreateTestEmployee(t, db.CreateEmployeeParams{
		FirstName: "浩士",
		LastName:  "青山",
		Email:     "test@email.com",
		Phone:     "090-1234-5678",
		Address:   "東京都千代田区",
	})
	CreateTestClockIn(t, employee.ID)
}
