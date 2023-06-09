package controllers_test

import (
	"bytes"
	"encoding/json"
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

func TestCreateClockOut(t *testing.T) {
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
				store.EXPECT().GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).Times(1).Return(clockin, nil)

				arg := db.ClockOutTxParams{
					EmployeeID: employee.ID,
					ClockInID:  clockin.ID,
				}
				store.EXPECT().ClockOutTx(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "BadRequest (not yet clocked in)",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).Times(1).Return(db.ClockIn{}, nil)

				arg := db.ClockOutTxParams{
					EmployeeID: employee.ID,
					ClockInID:  clockin.ID,
				}
				store.EXPECT().ClockOutTx(gomock.Any(), gomock.Eq(arg)).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest (already clocked out)",
			body: gin.H{
				"employee_id": employee.ID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				testutils.AddAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, employee.ID.String(), employee.IsAdmin, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				clockin.ClockedOut = true // 既にClockOutしている状態にする
				store.EXPECT().GetMostRecentClockIn(gomock.Any(), gomock.Eq(employee.ID)).Times(1).Return(clockin, nil)

				arg := db.ClockOutTxParams{
					EmployeeID: employee.ID,
					ClockInID:  clockin.ID,
				}
				store.EXPECT().ClockOutTx(gomock.Any(), gomock.Eq(arg)).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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
			url := "/api/clockouts"
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
