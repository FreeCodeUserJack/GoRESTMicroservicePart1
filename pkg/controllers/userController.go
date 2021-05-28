package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
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

func (u UserControllerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	w.Header().Set("content-type", "application/json")

	val, err := strconv.Atoi(userId)

	if err != nil {
		convErr := &utils.ApplicationError{
			Message: fmt.Sprintf("failed to convert userid (%s) to an uint64", userId),
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(convErr.String())
		return
	}

	foundUser, userServiceErr := u.UserService.GetUserById(uint64(val))

	if userServiceErr != nil {
		w.WriteHeader(userServiceErr.StatusCode)
		
		w.Write([]byte(userServiceErr.String()))
		return
	}

	err = encoder.Encode(foundUser)

	if err != nil {
		log.Fatalf("error encoding %v", foundUser)
	}
	
	w.WriteHeader(http.StatusOK)
}

// func UserControllerHi() {
// 	fmt.Println("user controller")
// 	services.UserServiceHi()
// }