package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// succeeds in creating a new role
func TestCreateRole(t *testing.T) {
	CreateTestRole(t)
}

// succeeds in getting a role by ID
func TestGetRoleByID(t *testing.T) {
	role1 := CreateTestRole(t)
	role2, err := testQueries.GetRoleByID(context.Background(), role1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, role2)
	require.Equal(t, role1.ID, role2.ID)
	require.Equal(t, role1.Name, role2.Name)
}

// succeeds in listing roles
func TestListRoles(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTestRole(t)
	}

	arg := ListRolesParams{
		Limit:  5,
		Offset: 5,
	}

	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roles, 5)

	for _, role := range roles {
		require.NotEmpty(t, role)
	}
}
