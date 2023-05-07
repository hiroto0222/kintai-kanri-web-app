package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
	"github.com/hiroto0222/kintai-kanri-web-app/testutils"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name          string
	setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []testCase{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, middlewares.AuthorizationTypeBearer, "test@email.com", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テストサーバを起動
			server := testutils.NewTestServer(t, nil)
			recorder := httptest.NewRecorder()

			// 認証が必要な api パスを設定
			authRequiredPath := "/api/auth"
			server.Router.GET(
				authRequiredPath,
				middlewares.AuthMiddleware(server.TokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			// テストリクエストを作成
			request, err := http.NewRequest(http.MethodGet, authRequiredPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

// addAuthorization はテストリクエストに authorization header を追加する
func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	email string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(email, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(middlewares.AuthorizationHeaderKey, authorizationHeader)
}
