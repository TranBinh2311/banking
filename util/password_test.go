package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TesPassword(t *testing.T) {
	password := RandomString(6)

	hassedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hassedPassword)

	err = CheckPassword(password, hassedPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hassedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hassedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hassedPassword2)
	require.NotEqual(t, hassedPassword, hassedPassword2)
}
