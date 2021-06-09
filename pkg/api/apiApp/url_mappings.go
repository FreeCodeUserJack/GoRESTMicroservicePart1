package apiApp

import (
	healthcheck "github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/controllers/healthCheck"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/controllers/repositories"
)


func mapUrls() {
	router.GET("/health", healthcheck.HandleHealthCheck)
	router.POST("/repositories", repositories.CreateRepo)
}