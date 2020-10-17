package domain

import (
	"fmt"
	"osvaldoabel/users-api/utils"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `valid:"uuid" gorm:"type:uuid;primary_key"`
	Name      string    `valid:"notnull"`
	Email     string    `valid:"notnull"`
	Status    string    `valid:"notnull"`
	UpdatedAt time.Time `valid:"-"`
	CreatedAt time.Time `valid:"-"`
}

func (user *User) prepare() {
	user.ID = uuid.NewV4().String()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

/**
	Description: Validates user fields
	Return: nil|error
**/
func (user *User) Validate() error {

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	isValid := utils.IsValidEmail(user.Email)

	if !isValid {
		return fmt.Errorf(" %q  is not a Valid Email", user.Email)
	}

	return nil
}

/**
	Description: Creates new user
	Return: User|error
**/
func NewUser(name string, email string, status string) (*User, error) {
	user := &User{}
	user.prepare()

	user.Name = name
	user.Email = email
	user.Status = status

	err := user.Validate()

	if err != nil {
		return &User{}, err
	}

	return user, nil
}
