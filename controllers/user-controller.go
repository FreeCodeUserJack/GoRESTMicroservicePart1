package controllers

import (
	"fmt"
	"log"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/services"
)

func UserController() {
	log.Println("hello")
	fmt.Println("user controller")
	services.PrintUser()
}