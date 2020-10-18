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

func TestIfNotValidEmail(t *testing.T) {
	_, err := NewUser("Osvaldo Abel", "wrong.com", "active")
	require.Error(t, err)
}
func TestNewUser(t *testing.T) {
	_, err := NewUser("Osvaldo Abel", "teste@example.com", "active")
	require.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	user, err := NewUser("Osvaldo Abel", "teste@example.com", "active")
	require.Nil(t, err)

	newName := "Osvaldo Abel updated"
	newEmail := "updated@example.com"
	newStatus := "inactive"
	updated, err := user.Update(newName, newEmail, newStatus)
	require.Nil(t, err)

	require.Equal(t, updated.ID, user.ID)
	require.Equal(t, updated.CreatedAt, user.CreatedAt)
	require.Equal(t, newName, user.Name)
	require.Equal(t, newEmail, user.Email)
	require.Equal(t, newStatus, user.Status)
	require.NotEqual(t, updated.CreatedAt, updated.UpdatedAt)
}
