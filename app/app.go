package app

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/controllers"
)

func StartApp() {
	fmt.Println("started app")
	controllers.UserController()
	fmt.Println("finished app")
}