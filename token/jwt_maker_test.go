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
	require.WithinDuration(issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(expiredAt, payload.ExpiresAt.Time, time.Second)

	expiredToken, err := maker.CreateToken(username, duration*-20)
	require.NoError(err)

	payload, err = maker.VerifyToken(expiredToken)
	require.ErrorIs(err, ErrTokenExpired)
	require.Nil(payload)

	invalidToken := util.RandomString(32)
	payload, err = maker.VerifyToken(invalidToken)
	require.ErrorIs(err, ErrInvalidToken)
	require.Nil(payload)
}
