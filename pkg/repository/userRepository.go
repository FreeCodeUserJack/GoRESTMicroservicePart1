package repository

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
)


type UserRepository interface {
	GetUserById(id uint64) *domain.User
}

type UserRepositoryImpl struct {
	DatabaseUrl string
	Username string
	Password string
}

func (u UserRepositoryImpl) GetUserById(id uint64) (*domain.User, error) {
	return 	&domain.User{
		Id: id,
		FirstName: "John",
		LastName: "Doe",
		Email: "johndoe@example.com",
	}, nil
}

func UserRepositoryHi() {
	fmt.Println("user repository")
	domain.UserDomainHi()
}