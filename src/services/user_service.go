package services

import (
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
)

type UserService struct {
	User           *domain.User
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{}
}

func (u *UserService) Insert() error {
	_, err := u.UserRepository.Insert(u.User)

	if err != nil {
		return err
	}

	return nil
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
