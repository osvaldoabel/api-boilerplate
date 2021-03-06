package repositories_test

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
	database "osvaldoabel/users-api/utils/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepositoryDbInsert(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", 26, "123456")
	require.Nil(t, err)

	repo := repositories.NewUserRepository()
	user, err = repo.Insert(user)

	require.Nil(t, err)
}

func TestUserRepositoryDbUpdate(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", 26, "123456")
	require.Nil(t, err)

	repo := repositories.NewUserRepository()
	user, err = repo.Insert(user)
	require.Nil(t, err)

	found, err := repo.Find(user.ID)
	require.Nil(t, err)

	found.Name = "new name"
	found.Email = "newemail@example.com"
	found.Address = "new address"

	err = found.Validate()
	require.Nil(t, err)

	updated, err := repo.Update(found)
	require.Nil(t, err)
	require.NotEqual(t, updated.Name, user.Name)
	require.NotEqual(t, updated.Email, user.Email)
	require.NotEqual(t, updated.Address, user.Address)
}

func TestUserRepositoryDbDelete(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", 26, "123456")
	require.Nil(t, err)

	repo := repositories.NewUserRepository()
	user, err = repo.Insert(user)
	require.Nil(t, err)

	err = repo.Delete(user.ID)
	require.Nil(t, err)
}
