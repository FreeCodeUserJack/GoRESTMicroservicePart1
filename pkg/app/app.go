package app

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/controllers"
)

func StartApp() {
	fmt.Println("starting app")
	controllers.UserController()
}