package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/hiroto0222/kintai-kanri-web-app/db/mock"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
	"github.com/stretchr/testify/require"
)

// Employee のカスタムマッチャーを作成
type eqCreateEmployeeParamsMatcher struct {
	arg      db.CreateEmployeeParams
	password string
}

func (e eqCreateEmployeeParamsMatcher) Matches(x interface{}) bool {
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

func (e eqCreateEmployeeParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateEmployeeParams(arg db.CreateEmployeeParams, password string) gomock.Matcher {
	return eqCreateEmployeeParamsMatcher{arg, password}
}

func TestSignUpEmployeeAPI(t *testing.T) {
	role := createTestRole()
	employee, password := createTestEmployee(t, role)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK (create non admin employee)",
			body: gin.H{
				"first_name": employee.FirstName,
				"last_name":  employee.LastName,
				"email":      employee.Email,
				"phone":      employee.Phone,
				"address":    employee.Address,
				"role_id":    employee.RoleID,
				"is_admin":   employee.IsAdmin,
				"password":   password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEmployeeParams{
					FirstName: employee.FirstName,
					LastName:  employee.LastName,
					Email:     employee.Email,
					Phone:     employee.Phone,
					Address:   employee.Address,
					RoleID:    employee.RoleID,
					IsAdmin:   employee.IsAdmin,
				}
				store.EXPECT().
					CreateEmployee(gomock.Any(), EqCreateEmployeeParams(arg, password)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchEmployee(t, recorder.Body, employee)
			},
		},
		{
			name: "OK (create admin employee)",
			body: gin.H{
				"first_name": employee.FirstName,
				"last_name":  employee.LastName,
				"email":      employee.Email,
				"phone":      employee.Phone,
				"address":    employee.Address,
				"role_id":    employee.RoleID,
				"is_admin":   true,
				"password":   password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEmployeeParams{
					FirstName: employee.FirstName,
					LastName:  employee.LastName,
					Email:     employee.Email,
					Phone:     employee.Phone,
					Address:   employee.Address,
					RoleID:    employee.RoleID,
					IsAdmin:   true,
				}
				store.EXPECT().
					CreateEmployee(gomock.Any(), EqCreateEmployeeParams(arg, password)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchEmployee(t, recorder.Body, employee)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// mockのコントローラを作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 実際DBに接続せず db.Store インターフェイスのmockを作成
			store := mockdb.NewMockStore(ctrl)

			// 作成したmockに対して期待する呼び出し関数とその引数、返り値を定義
			tc.buildStubs(store)

			// テストサーバを起動
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// テスト対象のURLとリクエストボディを定義
			url := "/api/auth/register"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// テストリクエストを作成
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func createTestEmployee(t *testing.T, role db.Role) (db.Employee, string) {
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
		RoleID:         role.ID,
	}, password
}

func createTestRole() db.Role {
	return db.Role{
		ID:   1,
		Name: "test role",
	}
}

func requireBodyMatchEmployee(t *testing.T, body *bytes.Buffer, employee db.Employee) {
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
