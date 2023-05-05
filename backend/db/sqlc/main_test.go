package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hiroto0222/kintai-kanri-web-app/config"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config, %q", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db, %q", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

// CreateTestEmployee creates a new employee for testing
func CreateTestEmployee(t *testing.T, arg CreateEmployeeParams) Employee {
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
func CreateTestRole(t *testing.T) Role {
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
