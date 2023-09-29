package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"time"

	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	"github.com/fauzanfebrian/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
		GinMode:             gin.TestMode,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

type resData[T any] struct {
	Data T `json:"data"`
}

func requireBodyMatch[T any](t *testing.T, body *bytes.Buffer, actualData T) {
	bodyData, err := io.ReadAll(body)
	require.NoError(t, err)

	var resData resData[T]
	err = json.Unmarshal(bodyData, &resData)
	require.NoError(t, err)
	require.Equal(t, actualData, resData.Data)
}
