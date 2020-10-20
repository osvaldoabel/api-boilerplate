package services

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
	"osvaldoabel/users-api/utils/database"
)

type UserService struct {
	User           *domain.User
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{}
}

func (u *UserService) Insert(payload map[string]string) (*domain.User, error) {

	user, err := domain.NewUser(payload["Name"], payload["Email"], payload["Status"], payload["Address"], payload["Password"])
	if err != nil {
		return nil, err
	}

	db := database.NewDbTest()
	defer db.Close()

	u.User = user
	u.UserRepository = repositories.NewUserRepository(db)
	user, err = u.UserRepository.Insert(u.User)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Update(user *domain.User) error {
	_, err := u.UserRepository.Update(u.User)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Delete(id string) error {
	err := u.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
