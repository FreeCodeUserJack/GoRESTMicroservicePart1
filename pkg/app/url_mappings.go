package app

import (
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/controllers"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/repository"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/services"
)


func mapUrls() {
	userRepo := repository.UserRepositoryInMemoryImpl{
		DatabaseUrl: "url",
		Username: "user",
		Password: "pass",
	}

	userServiceImpl := services.NewUserServiceImpl(userRepo)
	
	userController := controllers.NewUserControllerImpl(userServiceImpl)
	
	router.GET("/users/:userId", userController.GetUserById)
}