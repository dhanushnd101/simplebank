package api

import (
	"github.com/stretchr/testify/require"
	"time"
	"github.com/techschool/simplebank/util"
	"github.com/techschool/simplebank/db/sqlc"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server{
	config :=util.Config{
		TokenSymmetricKey: util.RandomeString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err:= NewServer(config, store)
	require.NoError(t,err)

	return server

}

func TestMain(m *testing.M){

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}