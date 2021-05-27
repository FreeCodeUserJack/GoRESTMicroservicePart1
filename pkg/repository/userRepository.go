package repository

import (
	"errors"
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
)

// in-memory storage for now
var (
	users = []*domain.User{
		{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"},
		{Id: 2, FirstName: "Jane", LastName: "Summers", Email: "janesummers@example.com"},
		{Id: 3, FirstName: "Oberon", LastName: "Falls", Email: "oberonfalls@example.com"},
	}
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
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil,  errors.New("user with id: " + fmt.Sprintf("%v", id) + " not found")
}

func UserRepositoryHi() {
	fmt.Println("user repository")
	domain.UserDomainHi()
}