package token

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestPestoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomeString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username,duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t,payload)
	

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t,payload)
	require.WithinDuration(t, payload.IssuedAt,issuedAt,time.Second)
	require.WithinDuration(t, payload.ExpiredAt,expiredAt,time.Second)
	require.Equal(t, payload.Username,username)
	require.NotZero(t, payload.ID)
}

func TestExpiredPestoToken(t *testing.T){
	maker, err := NewPasetoMaker(util.RandomeString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomOwner(),-time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t,payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t,err)
	require.EqualError(t,err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
