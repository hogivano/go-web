package user

import (
	"go-web/helper"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	LoginUser(input LoginUserInput) (User, error)
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(password)

	newUser, err := s.repository.Login(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.SecretKey = helper.RandomString(20)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}
