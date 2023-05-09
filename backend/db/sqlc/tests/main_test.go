package db_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatalf("cannot load config, %q", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db, %q", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}

// CreateTestEmployee creates a new employee for testing
func CreateTestEmployee(t *testing.T, arg db.CreateEmployeeParams) db.Employee {
	employee, err := testQueries.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	t.Cleanup(func() {
		if err := testQueries.DeleteEmployee(context.Background(), employee.ID); err != nil {
			t.Errorf("failed to delete employee, %q", err)
		}
	})

	return employee
}

// CreateTestRole creates a new role for testing
func CreateTestRole(t *testing.T) db.Role {
	role, err := testQueries.CreateRole(context.Background(), "test role")
	require.NoError(t, err)
	require.NotEmpty(t, role)

	t.Cleanup(func() {
		if err := testQueries.DeleteRole(context.Background(), role.ID); err != nil {
			t.Errorf("failed to delete role, %q", err)
		}
	})

	return role
}

// CreateTestClockIn creates a new ClockIn for testing
func CreateTestClockIn(t *testing.T, employeeID uuid.UUID) db.ClockIn {
	clockIn, err := testQueries.CreateClockIn(context.Background(), employeeID)
	require.NoError(t, err)
	require.NotEmpty(t, clockIn)

	t.Cleanup(func() {
		if err := testQueries.DeleteClockIn(context.Background(), clockIn.ID); err != nil {
			t.Errorf("failed to delete clockin, %q", err)
		}
	})

	return clockIn
}
