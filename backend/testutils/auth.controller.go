package testutils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
	"github.com/stretchr/testify/require"
)

// CreateTestEmployee はテスト用の Employee を作成
func CreateTestEmployee(t *testing.T, role db.Role) (db.Employee, string) {
	password := utils.RandomString(10)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	return db.Employee{
		FirstName:      "Hiroto",
		LastName:       "Aoyama",
		Email:          "test@email.com",
		Phone:          "090-1234-5678",
		Address:        "Tokyo",
		HashedPassword: hashedPassword,
		RoleID: sql.NullInt32{
			Int32: role.ID,
			Valid: true,
		},
	}, password
}

// CreateTestRole はテスト用の Role を作成
func CreateTestRole() db.Role {
	return db.Role{
		ID:   1,
		Name: "test role",
	}
}

// RequireBodyMatchRole は Employee 作成時の response の body を検証
func RequireBodyMatchEmployee(t *testing.T, body *bytes.Buffer, employee db.Employee) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var got db.Employee
	err = json.Unmarshal(data, &got)
	fmt.Println(got)
	require.NoError(t, err)

	require.Equal(t, employee.FirstName, got.FirstName)
	require.Equal(t, employee.LastName, got.LastName)
	require.Equal(t, employee.Email, got.Email)
	require.Equal(t, employee.Phone, got.Phone)
	require.Equal(t, employee.Address, got.Address)
	require.Equal(t, employee.RoleID, got.RoleID)
	require.Equal(t, employee.IsAdmin, got.IsAdmin)
	require.Empty(t, got.HashedPassword)
}

// Employee のカスタムマッチャーを作成
type EqCreateEmployeeParamsMatcher struct {
	arg      db.CreateEmployeeParams
	password string
}

func (e EqCreateEmployeeParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateEmployeeParams)
	if !ok {
		return false
	}

	err := utils.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e EqCreateEmployeeParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateEmployeeParams(arg db.CreateEmployeeParams, password string) gomock.Matcher {
	return EqCreateEmployeeParamsMatcher{arg, password}
}
