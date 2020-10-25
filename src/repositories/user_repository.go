package repositories

import (
	"fmt"
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/utils/database"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Find(id string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	All(params map[string]string) []*domain.User
	Delete(id string) error
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

/**
	Description: "Instanceate" a new and clean UserRepositoryDb.
	It receives the Entity User (*domain.User) and returns

	Return: connection *gorm.DB
**/
func NewUserRepository() *UserRepositoryDb {
	db := database.NewDbConnection()

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
	repository.Db.Find(&user, "id=?", id)

	if user.ID == "" {
		return nil, fmt.Errorf("This User doesn't exist.")
	}

	return &user, nil
}

func (repository *UserRepositoryDb) All(params map[string]string) []*domain.User {
	var users []*domain.User

	repository.Db.Limit(params["Limit"]).Offset(params["Offset"]).Order("created_at DESC", true).Find(&users)
	return users
}

func (repository *UserRepositoryDb) Delete(id string) error {
	var user domain.User
	result := repository.Db.Delete(&user, "id = ?", id)

	if result.RowsAffected == 0 {
		return fmt.Errorf("This user doesn't exist")
	}

	return nil
}
