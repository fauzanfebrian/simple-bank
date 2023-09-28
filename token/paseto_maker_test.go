package token

import (
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	require := require.New(t)

	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(err)
	require.NotEmpty(token)

	payload, err := maker.VerifyToken(token)
	require.NoError(err)
	require.NotEmpty(payload)

	require.NotZero(payload.ID)
	require.Equal(username, payload.Username)
	require.WithinDuration(issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(expiredAt, payload.ExpiredAt, time.Second)
}

func TestPasetoExpiredToken(t *testing.T) {
	require := require.New(t)

	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(err)

	username := util.RandomOwner()
	duration := time.Minute * -20

	expiredToken, err := maker.CreateToken(username, duration)
	require.NoError(err)
	require.NotEmpty(expiredToken)

	payload, err := maker.VerifyToken(expiredToken)
	require.ErrorIs(err, ErrTokenExpired)
	require.NotEmpty(payload)
}
