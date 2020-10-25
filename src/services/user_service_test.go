package services_test

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/services"
	"osvaldoabel/users-api/utils"
	"osvaldoabel/users-api/utils/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserServiceInsert(t *testing.T) {

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", 30, "123456")
	require.Nil(t, err)
	require.NotNil(t, user)

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()

	payload := &utils.UserPayload{
		Name:     "Osvaldo Abel",
		Email:    "teste@example.com",
		Status:   "active",
		Address:  "My  Street , 15-30",
		Age:      15,
		Password: "123456",
	}

	user, err = userService.Insert(payload)
	require.Nil(t, err)
}

func TestUserServiceUpdate(t *testing.T) {

	payload := &utils.UserPayload{
		Name:     "Osvaldo Abel",
		Email:    "teste@example.com",
		Status:   "active",
		Address:  "My  Street , 15-30",
		Age:      28,
		Password: "123456",
	}

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()
	user, err := userService.Insert(payload)
	require.Nil(t, err)
	require.NotNil(t, user)

	found, err := userService.UserRepository.Find(userService.User.ID)
	found.Name = "UPDATED - Osvaldo Abel"
	found.Address = "UPDATED - my new adderess"
	found.Email = "new.email@example.com"

	updated, err := userService.UserRepository.Update(found)
	require.Nil(t, err)
	require.NotEqual(t, updated.CreatedAt, updated.UpdatedAt)
	require.Equal(t, updated.ID, updated.ID)
}
func TestUserServiceDelete(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	payload := &utils.UserPayload{
		Name:     "Osvaldo Abel",
		Email:    "teste@example.com",
		Status:   "active",
		Address:  "My  Street , 15-30",
		Age:      30,
		Password: "123456",
	}

	userService := services.NewUserService()
	user, err := userService.Insert(payload)
	require.Nil(t, err)
	require.NotNil(t, user)

	err = userService.Delete(userService.User.ID)
	require.Nil(t, err)
}
