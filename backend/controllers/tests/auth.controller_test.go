package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/hiroto0222/kintai-kanri-web-app/db/mock"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/testutils"
	"github.com/stretchr/testify/require"
)

func TestRegisterEmployeeAPI(t *testing.T) {
	role := testutils.CreateTestRole()
	employee, password := testutils.CreateTestEmployee(t, role)

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
					CreateEmployee(gomock.Any(), testutils.EqCreateEmployeeParams(arg, password)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				testutils.RequireBodyMatchEmployee(t, recorder.Body, employee)
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
					CreateEmployee(gomock.Any(), testutils.EqCreateEmployeeParams(arg, password)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				testutils.RequireBodyMatchEmployee(t, recorder.Body, employee)
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
			server := testutils.NewTestServer(t, store)
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

func TestLogInEmployee(t *testing.T) {
	role := testutils.CreateTestRole()
	employee, password := testutils.CreateTestEmployee(t, role)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    employee.Email,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByEmail(gomock.Any(), gomock.Eq(employee.Email)).
					Times(1).
					Return(employee, nil)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()). // TODO:
					Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "EmployeeNotFound",
			body: gin.H{
				"email":    "notfound@email.com",
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Employee{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "WrongPassword",
			body: gin.H{
				"email":    employee.Email,
				"password": "wrongpassword",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByEmail(gomock.Any(), gomock.Eq(employee.Email)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
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
			server := testutils.NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			// テスト対象のURLとリクエストボディを定義
			url := "/api/auth/login"
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
