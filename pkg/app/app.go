package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	fmt.Println("init func of app.go")
	router = gin.Default()
}

func StartApp() {
	fmt.Println("starting app")
	// controllers.UserControllerHi()

	mapUrls()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("error listening on port 8080")
	}
}

