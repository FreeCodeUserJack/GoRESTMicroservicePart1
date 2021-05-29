package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/controllers"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/repository"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
)

func init() {
	fmt.Println("init func of app.go")
}

func StartApp() {
	fmt.Println("starting app")
	// controllers.UserControllerHi()

	userRepo := repository.UserRepositoryInMemoryImpl{
		DatabaseUrl: "url",
		Username: "user",
		Password: "pass",
	}

	userServiceImpl := services.NewUserServiceImpl(userRepo)
	
	userController := controllers.NewUserControllerImpl(userServiceImpl)

	http.HandleFunc("/users/", userController.GetUserById)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error listening on port 8080")
	}
}