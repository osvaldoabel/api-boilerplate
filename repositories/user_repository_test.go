package repositories_test

import (
	"osvaldoabel/users-api/domain"
	database "osvaldoabel/users-api/framework/database"
	"osvaldoabel/users-api/repositories"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepositoryDbInsert(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := domain.NewUser("Osvaldo Abel", "teste@example.com", "active", "My  Street , 15-30", "123456")
	require.Nil(t, err)

	repo := repositories.NewUserRepository(db)
	user, err = repo.Insert(user)
	require.Nil(t, err)
}