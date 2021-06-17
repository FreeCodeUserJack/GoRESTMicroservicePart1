package app

import (
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/controllers/healthCheck"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/oauth-api/controller/oauth"
)


func mapUrls() {
	router.GET("/health", healthCheck.HandleHealthCheck)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}