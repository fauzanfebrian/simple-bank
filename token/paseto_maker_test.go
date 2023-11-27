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
	role := util.DepositorRole
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payloadGenerated, err := maker.CreateToken(username, role, duration)
	require.NoError(err)
	require.NotEmpty(token)
	require.NotEmpty(payloadGenerated)

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
	role := util.DepositorRole
	duration := time.Minute * -20

	expiredToken, payloadGenerated, err := maker.CreateToken(username, role, duration)
	require.NoError(err)
	require.NotEmpty(expiredToken)
	require.NotEmpty(payloadGenerated)

	payload, err := maker.VerifyToken(expiredToken)
	require.ErrorIs(err, ErrTokenExpired)
	require.NotEmpty(payload)
}
