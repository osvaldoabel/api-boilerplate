package repositories

import (
	"fmt"
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/utils"
	"osvaldoabel/users-api/utils/database"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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

func init() {
	godotenv.Load("./../src/.env")
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
	message := ""
	err := repository.Db.Create(user).Error

	if err != nil {
		message = fmt.Sprintf("Coulden't Insert User. [Name: %s].", user.Name)
		return nil, err
	}

	message = fmt.Sprintf("User Inserted. [ID: %s].", user.ID)
	utils.App_log(message)
	return user, nil
}

/**
	Description: Updates an User Entity.

	Return: user, error
**/
func (repository *UserRepositoryDb) Update(user *domain.User) (*domain.User, error) {

	err := repository.Db.Save(&user).Error
	message := ""

	if err != nil {
		message = fmt.Sprintf("Couldn't Update User [ID: %s].", user.ID)
		utils.App_log(message)
		return nil, err
	}

	message = fmt.Sprintf("User Updated. [ID: %s].", user.ID)
	utils.App_log(message)
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
	message := ""
	var user domain.User

	result := repository.Db.Delete(&user, "id = ?", id)

	if result.RowsAffected == 0 {
		message = fmt.Sprintf("This user doesn't exist. [ID: %s]", id)
		utils.App_log(message)
		return fmt.Errorf(message)
	}

	message = fmt.Sprintf("User [ID: %s] Deleted.", id)
	utils.App_log(message)
	return nil
}
