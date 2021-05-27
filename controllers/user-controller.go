package controllers

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/services"
)

func UserController() {
	fmt.Println("user controller")
	services.PrintUser()
}