package services

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/repository"
)

type UserService interface {
	GetUserById(id uint64) (*domain.User, error)
}

type UserServiceImpl struct {}

func (u *UserServiceImpl) GetUserById(id uint64) (*domain.User, error) {
	fmt.Println("user service")

	user, err := repository.UserRepositoryImpl{
		DatabaseUrl: "url",
		Username: "user",
		Password: "pass",
	}.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserServiceHi() {
	fmt.Println("user service")
	repository.UserRepositoryHi()
}