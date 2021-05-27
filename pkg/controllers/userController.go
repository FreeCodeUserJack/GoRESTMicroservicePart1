package controllers

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
)


func UserController() {
	fmt.Println("user controller")
	services.UserService()
}