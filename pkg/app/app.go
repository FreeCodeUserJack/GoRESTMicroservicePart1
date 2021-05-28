package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/controllers"
)

func StartApp() {
	fmt.Println("starting app")
	// controllers.UserControllerHi()

	userController := controllers.UserControllerImpl{}

	http.HandleFunc("/users/", userController.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error listening on port 8080")
	}
}