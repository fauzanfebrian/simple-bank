package api

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/fauzanfebrian/simplebank/config"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	config.LoadConfig()
	os.Exit(m.Run())
}

func requireBodyMatch[T any](t *testing.T, body *bytes.Buffer, actualData T) {
	type getResData struct {
		Data T `json:"data"`
	}

	bodyData, err := io.ReadAll(body)
	require.NoError(t, err)

	var resData getResData
	err = json.Unmarshal(bodyData, &resData)
	require.NoError(t, err)
	require.Equal(t, actualData, resData.Data)
}
