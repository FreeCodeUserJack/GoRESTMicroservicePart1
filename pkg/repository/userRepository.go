package repository

import (
	"fmt"
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
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

func (u UserRepositoryImpl) GetUserById(id uint64) (*domain.User, *utils.ApplicationError) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	// use our own error type instead of errors.New("...")
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user with id %d was not found", id),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}

func UserRepositoryHi() {
	fmt.Println("user repository")
	domain.UserDomainHi()
}