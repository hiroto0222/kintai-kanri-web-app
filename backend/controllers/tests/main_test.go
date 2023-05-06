package controllers_test

import (
	"testing"

	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/server"
)

func newTestServer(t *testing.T, store db.Store) *server.Server {
	config := config.Config{}
	server := server.NewServer(config, store)
	return server
}
