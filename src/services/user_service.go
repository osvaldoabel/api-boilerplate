package services

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
	"osvaldoabel/users-api/utils"
)

type UserService struct {
	User           *domain.User
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{}
}

func (u *UserService) Insert(payload *utils.UserPayload) (*domain.User, error) {

	user, err := domain.NewUser(payload.Name, payload.Email, payload.Status, payload.Address, payload.Age, payload.Password)
	if err != nil {
		return nil, err
	}

	u.User = user
	u.UserRepository = repositories.NewUserRepository()
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

func (u *UserService) All(params map[string]string) []*domain.User {
	return u.UserRepository.All(params)
}

func (u *UserService) Delete(id string) error {
	err := u.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
