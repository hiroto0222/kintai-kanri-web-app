package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	mockdb "github.com/hiroto0222/kintai-kanri-web-app/db/mock"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
	"github.com/hiroto0222/kintai-kanri-web-app/testutils"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/stretchr/testify/require"
)

func TestGetEmployee(t *testing.T) {
	role := testutils.CreateTestRole()
	employee, _ := testutils.CreateTestEmployee(t, role)
	adminEmployee, _ := testutils.CreateTestAdminEmployee(t, role)

	testCases := []struct {
		name          string
		employeeID    string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK (自分は自分を取得できる)",
			employeeID: employee.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeById(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				testutils.RequireBodyMatchEmployee(t, recorder.Body, employee)
			},
		},
		{
			name:       "OK (Adminは他人を取得できる)",
			employeeID: employee.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, adminEmployee.ID.String(), adminEmployee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeById(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				testutils.RequireBodyMatchEmployee(t, recorder.Body, employee)
			},
		},
		{
			name:       "UnAuthorized (他人は他人を取得できない)",
			employeeID: adminEmployee.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeById(gomock.Any(), gomock.Eq(adminEmployee.ID)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
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

			// テスト対象のURLを作成
			url := fmt.Sprintf("/api/employees/%s", tc.employeeID)

			// テストリクエストを作成
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListEmployees(t *testing.T) {
	role := testutils.CreateTestRole()
	employee, _ := testutils.CreateTestEmployee(t, role)
	adminEmployee, _ := testutils.CreateTestAdminEmployee(t, role)

	n := 5
	employees := make([]db.Employee, n)
	for i := 0; i < n; i++ {
		employees[i], _ = testutils.CreateTestEmployee(t, role)
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, adminEmployee.ID.String(), adminEmployee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListEmployeesParams{
					Limit:  int32(n),
					Offset: 0,
				}
				store.EXPECT().
					ListEmployees(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(employees, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEmployees(t, recorder.Body, employees)
			},
		},
		{
			name: "UnAuthorized",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
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

			// テスト対象のURLを作成
			url := "/api/employees"

			// テストリクエストを作成
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchEmployees(t *testing.T, body *bytes.Buffer, employees []db.Employee) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotEmployees []controllers.EmployeeResponse
	err = json.Unmarshal(data, &gotEmployees)

	var wantEmployees []controllers.EmployeeResponse
	for _, employee := range employees {
		wantEmployees = append(wantEmployees, controllers.NewUserResponse(employee))
	}

	require.NoError(t, err)
	require.Equal(t, wantEmployees, gotEmployees)
}
