package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserControllerImpl(userService services.UserService) UserControllerImpl {
	return UserControllerImpl{
		UserService: userService,
	}
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	
	c.Writer.Header().Set("content-type", "application/json")

	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		convErr := &utils.ApplicationError{
			Message: fmt.Sprintf("failed to convert userid (%d) to an uint64", userId),
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}

		utils.RespondError(c, convErr)
		return
	}

	foundUser, userServiceErr := u.UserService.GetUserById(uint64(userId))

	if userServiceErr != nil {
		utils.RespondError(c, userServiceErr)
		return
	}

	utils.Respond(c, http.StatusOK, foundUser)
}

// func UserControllerHi() {
// 	fmt.Println("user controller")
// 	services.UserServiceHi()
// }