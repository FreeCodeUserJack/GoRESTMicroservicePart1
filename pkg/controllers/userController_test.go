package controllers

import (
	"net/http"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)


func TestUserController(t *testing.T) {
	testUserServiceImpl := utils.TestUserServiceImpl{}

	t.Run("get existing user", func(t *testing.T) {
		userId := uint64(1)
		want := &domain.User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"}
	
		user, err := testUserServiceImpl.GetUserById(userId)
	
		utils.AssertNoApplicationError(t, err)
	
		utils.AssertEqualInstance(t, user, want)
	})

	t.Run("user not exists", func(t *testing.T) {
		userId := uint64(99)
		
		_, err := testUserServiceImpl.GetUserById(userId)

		utils.AssertApplicationError(t, err, http.StatusNotFound, "not found")
	})
}