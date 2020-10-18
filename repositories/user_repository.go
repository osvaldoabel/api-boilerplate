package repositories

import (
	"osvaldoabel/users-api/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

/**
	Description: "Instanceate" a new and clean UserRepositoryDb.
	It receives the Entity User (*domain.User) and returns

	Return: connection *gorm.DB
**/
func NewUserRepository(db *gorm.DB) *UserRepositoryDb {
	return &UserRepositoryDb{Db: db}
}

/**
	Description: Create a new User registry.
	It receives the Entity User (*domain.User) and returns

	Return: user, error
**/
func (repository *UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := repository.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
