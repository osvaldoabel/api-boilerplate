package repositories

import (
	"fmt"
	"osvaldoabel/users-api/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Find(id string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
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

/**
	Description: Updates an User Entity.

	Return: user, error
**/
func (repository *UserRepositoryDb) Update(user *domain.User) (*domain.User, error) {

	err := repository.Db.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepositoryDb) Find(id string) (*domain.User, error) {
	var user domain.User
	repository.Db.Preload("user").First(&user, "id = ?", id)

	if user.ID == "" {
		return nil, fmt.Errorf("This User doesn't exist.")
	}

	return &user, nil
}
