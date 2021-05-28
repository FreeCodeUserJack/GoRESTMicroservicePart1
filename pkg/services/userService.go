package services

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/repository"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)

type UserService interface {
	GetUserById(id uint64) (*domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserServiceImpl {
	return UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (u UserServiceImpl) GetUserById(id uint64) (*domain.User, *utils.ApplicationError) {
	fmt.Println("user service")

	user, err := u.UserRepository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// func UserServiceHi() {
// 	fmt.Println("user service")
// 	repository.UserRepositoryHi()
// }