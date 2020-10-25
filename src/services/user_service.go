package services

import (
	"fmt"
	"osvaldoabel/users-api/src/domain"
	"osvaldoabel/users-api/src/repositories"
	"osvaldoabel/users-api/utils"
)

type UserService struct {
	User           *domain.User
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{UserRepository: repositories.NewUserRepository()}
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

func (u *UserService) Update(id string, payload *utils.UserPayload) (*domain.User, error) {
	user, err := u.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}
	if payload.Email != "" {
		user.Email = payload.Email
	}
	if payload.Status != "" {
		user.Status = payload.Status
	}
	if payload.Address != "" {
		user.Address = payload.Address
	}
	if payload.Age >= 0 {
		user.Age = payload.Age
	}
	if payload.Password != "" {
		user.Password = payload.Password
	}

	user, err = u.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) All(params map[string]string) []*domain.User {
	return u.UserRepository.All(params)
}

func (u *UserService) Find(id string) (*domain.User, error) {
	user, err := u.UserRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Delete(id string) error {

	found, err := u.UserRepository.Find(id)

	if err != nil {
		return err
	}

	if found.ID == "" {
		return fmt.Errorf("This User doesn't exist.")
	}

	return u.UserRepository.Delete(id)
}
