package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateIfUserIsEmpty(t *testing.T) {
	user := &User{}
	err := user.Validate()
	require.Error(t, err)
}

func TestIfIdIsNotValidUUID(t *testing.T) {

	user := &User{}
	user.ID = "123"
	user.Name = "osvaldo abel"
	user.Email = "test@example.com"
	user.Status = "active"

	err := user.Validate()
	require.Error(t, err)
}

func TestNewUser(t *testing.T) {
	user, err := NewUser("Osvaldo Abel", "teste@example.com", "active")
	require.Nil(t, err)

	err = user.Validate()
	require.Nil(t, err)
}

func TestIfNotValidEmail(t *testing.T) {
	_, err := NewUser("Osvaldo Abel", "wrong.com", "active")
	require.Error(t, err)
}
