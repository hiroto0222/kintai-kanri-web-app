package token_test

import (
	"testing"
	"time"

	"github.com/hiroto0222/kintai-kanri-web-app/token"
	"github.com/hiroto0222/kintai-kanri-web-app/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	email := utils.RandomEmail()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	jwtToken, err := maker.CreateToken(email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	payload, err := maker.VerifyToken(jwtToken)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := token.NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	jwtToken, err := maker.CreateToken(utils.RandomEmail(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	payload, err := maker.VerifyToken(jwtToken)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
