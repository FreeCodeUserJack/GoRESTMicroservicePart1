package repository

import (
	"net/http"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)


func TestUserRepositoryIntegration(t *testing.T) {
	t.Run("get user back from user repository", func(t *testing.T) {
		userRepository := UserRepositoryInMemoryImpl{}

		userId := uint64(1)
	
		want := &domain.User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"}
	
		got, err := userRepository.GetUserById(userId)
	
		utils.AssertNoApplicationError(t, err)
	
		utils.AssertUserFound(t, got, want)
	})

	t.Run("get err when user not in user repository", func(t *testing.T) {
		userRepository := UserRepositoryInMemoryImpl{}

		userId := uint64(4)

		_, err := userRepository.GetUserById(userId)

		utils.AssertApplicationError(t, err, http.StatusNotFound, httpNotFoundCode)
	})
}