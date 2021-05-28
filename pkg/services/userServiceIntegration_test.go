package services

import (
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/repository"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)


func TestUserServiceIntegration(t *testing.T) {
	t.Run("get existing user", func(t *testing.T) {
		userServiceImpl := NewUserServiceImpl(repository.UserRepositoryInMemoryImpl{})

		userId := uint64(1)
	
		want := &domain.User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"}
	
		user, err := userServiceImpl.GetUserById(userId)
	
		utils.AssertNoApplicationError(t, err)
	
		utils.AssertUserFound(t, user, want)
	})

	t.Run("user not found return err", func(t *testing.T) {
		userServiceImpl := NewUserServiceImpl(repository.UserRepositoryInMemoryImpl{})

		userId := uint64(4)

		_, err := userServiceImpl.GetUserById(userId)

		utils.AssertError(t, err)
	})
}