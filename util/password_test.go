package util

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPassword(t *testing.T){
	password := RandomeString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t,hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := RandomeString(6)

	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}