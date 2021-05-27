package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
)

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {}

func (u *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	w.Header().Set("content-type", "application/json")

	val, err := strconv.Atoi(userId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}

	userService := services.UserServiceImpl{}

	foundUser, err := userService.GetUserById(uint64(val))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(foundUser)

	if err != nil {
		log.Fatalf("error encoding %v", foundUser)
	}
	
	w.WriteHeader(http.StatusOK)
}

func UserControllerHi() {
	fmt.Println("user controller")
	services.UserServiceHi()
}