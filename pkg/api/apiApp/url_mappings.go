package apiApp

import (
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/controllers/healthCheck"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/controllers/repositories"
)


func mapUrls() {
	router.GET("/health", healthCheck.HandleHealthCheck)
	router.POST("/repositories", repositories.CreateRepo)
}