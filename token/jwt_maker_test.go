package token

import (
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	require := require.New(t)

	maker, err := NewJWTMaker(util.RandomString(34))
	require.NoError(err)

	username := util.RandomOwner()
	duration := time.Minute

	// issuedAt := time.Now()
	// expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(err)
	require.NotEmpty(token)

	// payload, err := maker.VerifyToken(token)
	// fmt.Println(err)
	// require.NoError(err)
	// require.NotEmpty(payload)

	// require.NotZero(payload.ID)
	// require.Equal(username, payload.Username)
	// require.WithinDuration(issuedAt, payload.IssuedAt.Time, time.Second)
	// require.WithinDuration(expiredAt, payload.ExpiresAt.Time, time.Second)
}
