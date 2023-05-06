package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// succeeds in creating a new employee
func TestCreateEmployee(t *testing.T) {
	role := CreateTestRole(t)
	CreateTestEmployee(t, CreateEmployeeParams{
		FirstName: "浩士",
		LastName:  "青山",
		Email:     "test@email.com",
		Phone:     "090-1234-5678",
		Address:   "東京都千代田区",
		RoleID:    role.ID,
	})
}

// succeeds in getting an employee by ID
func TestGetEmployeeByID(t *testing.T) {
	role := CreateTestRole(t)
	employee1 := CreateTestEmployee(t, CreateEmployeeParams{
		FirstName: "浩士",
		LastName:  "青山",
		Email:     "test@email.com",
		Phone:     "090-1234-5678",
		Address:   "東京都千代田区",
		RoleID:    role.ID,
	})

	employee2, err := testQueries.GetEmployeeById(context.Background(), employee1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, employee2)
	requireEqualEmployee(t, employee1, employee2)
}

// succeeds in listing employees
func TestListEmployees(t *testing.T) {
	role := CreateTestRole(t)
	for i := 0; i < 10; i++ {
		CreateTestEmployee(t, CreateEmployeeParams{
			FirstName: "浩士",
			LastName:  "青山",
			Email:     fmt.Sprintf("test{%v}@email.com", i),
			Phone:     fmt.Sprintf("090-1234-567{%v}", i),
			Address:   "東京都千代田区",
			RoleID:    role.ID,
		})
	}

	arg := ListEmployeesParams{
		Limit:  5,
		Offset: 5,
	}

	employees, err := testQueries.ListEmployees(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, employees, 5)

	for _, employee := range employees {
		require.NotEmpty(t, employee)
	}
}

func requireEqualEmployee(t *testing.T, employee1, employee2 Employee) {
	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, employee1.FirstName, employee2.FirstName)
	require.Equal(t, employee1.LastName, employee2.LastName)
	require.Equal(t, employee1.Email, employee2.Email)
	require.Equal(t, employee1.Phone, employee2.Phone)
	require.Equal(t, employee1.Address, employee2.Address)
	require.Equal(t, employee1.RoleID, employee2.RoleID)
	require.Equal(t, employee1.CreatedAt, employee2.CreatedAt)
}
