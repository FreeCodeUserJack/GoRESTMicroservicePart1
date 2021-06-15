package apiApp

import (
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/log"
	"github.com/gin-gonic/gin"
)


var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("app started", "step:1", "status:pending")

	mapUrls()

	log.Info("urls mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}