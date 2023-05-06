package controllers_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/server"
)

func newTestServer(t *testing.T, store db.Store) *server.Server {
	config := config.Config{}
	server := server.NewServer(config, store)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
