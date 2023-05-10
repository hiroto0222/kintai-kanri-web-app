package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/hiroto0222/kintai-kanri-web-app/db/mock"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
	"github.com/hiroto0222/kintai-kanri-web-app/testutils"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/stretchr/testify/require"
)

func TestCreateClockIn(t *testing.T) {
	role := testutils.CreateTestRole()
	employee, _ := testutils.CreateTestEmployee(t, role)
	clockin := testutils.CreateTestClockIn(t, employee)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(clockin, nil)
				store.EXPECT().
					GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(db.ClockIn{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchClockIn(t, recorder.Body, clockin)
			},
		},
		{
			name: "UnAuthorized",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, "UnAuthorized", employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "BadRequest, 退出打刻せずにまた出勤打刻した場合",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(0)
				store.EXPECT().
					GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					// 退出打刻していない状態
					Return(db.ClockIn{
						ID:          1,
						EmployeeID:  employee.ID,
						ClockInTime: time.Now(),
						ClockedOut:  false,
					}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "OK, 退出打刻していて出勤打刻した場合",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(clockin, nil)
				store.EXPECT().
					GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(db.ClockIn{
						ID:          1,
						EmployeeID:  employee.ID,
						ClockInTime: time.Now(),
						ClockedOut:  true,
					}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchClockIn(t, recorder.Body, clockin)
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
			url := "/api/clockins/"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// テストリクエストを作成
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchClockIn(t *testing.T, body *bytes.Buffer, clockIn db.ClockIn) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var got db.ClockIn
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, clockIn.ID, got.ID)
	require.Equal(t, clockIn.EmployeeID, got.EmployeeID)
	require.WithinDuration(t, clockIn.ClockInTime, got.ClockInTime, time.Second*2)
}
