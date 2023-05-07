package testutils

import (
	"testing"
	"time"

	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/server"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *server.Server {
	config := config.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := server.NewServer(config, store)
	require.NoError(t, err)

	return server
}
