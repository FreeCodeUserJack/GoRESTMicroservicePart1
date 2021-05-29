package utils

import (
	"fmt"
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
)

var (
	users = []*domain.User{
		{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"},
		{Id: 2, FirstName: "Jane", LastName: "Summers", Email: "janesummers@example.com"},
		{Id: 3, FirstName: "Oberon", LastName: "Falls", Email: "oberonfalls@example.com"},
	}
)

type TestUserInMemoryRepository struct {
}

func (u TestUserInMemoryRepository) GetUserById(id uint64) (*domain.User, *ApplicationError) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	// use our own error type instead of errors.New("...")
	return nil, &ApplicationError{
		Message: fmt.Sprintf("user with id %d was not found", id),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}

type TestUserServiceImpl struct {
	numberCalls []uint64
}

func (t TestUserServiceImpl) GetUserById(id uint64) (*domain.User, *ApplicationError) {
	t.numberCalls = append(t.numberCalls, id)
	return TestUserInMemoryRepository{}.GetUserById(id)
}