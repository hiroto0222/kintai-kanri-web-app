package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "Bearer"
	AuthorizationPayloadKey = "authorization_payload" // gin context に保存する際の key
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			// ctx.abort することで以降の処理を実行せずにレスポンスを返す
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		// authorization header のフォーマットを確認
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		// Bearer がついているか確認
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != strings.ToLower(AuthorizationTypeBearer) {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		// token を確認
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		// gin context に保存して後続の処理に移動
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
