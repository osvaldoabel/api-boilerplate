package services_test

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserServiceInsert(t *testing.T) {

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", "123456")
	require.Nil(t, err)

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()
	userService.User = user
	userService.UserRepository = repositories.NewUserRepository(db)

	err = userService.Insert()
	require.Nil(t, err)
}

func TestUserServiceUpdate(t *testing.T) {

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", "123456")
	require.Nil(t, err)

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()
	userService.User = user
	userService.UserRepository = repositories.NewUserRepository(db)

	err = userService.Insert()
	require.Nil(t, err)

	found, err := userService.UserRepository.Find(userService.User.ID)
	userService.User.Name = "UPDATED - Osvaldo Abel"
	userService.User.Address = "UPDATED - my new adderess"
	userService.User.Email = "new.email@example.com"
	found, err = userService.UserRepository.Update(userService.User)
	require.Nil(t, err)
	require.NotEqual(t, found.CreatedAt, found.UpdatedAt)
	require.Equal(t, found.ID, found.ID)
}
func TestUserServiceDelete(t *testing.T) {

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", "123456")
	require.Nil(t, err)

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()
	userService.User = user
	userService.UserRepository = repositories.NewUserRepository(db)

	err = userService.Insert()
	require.Nil(t, err)

	err = userService.Delete(userService.User.ID)

	require.Nil(t, err)
}
