package test_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetMockedContext(response http.ResponseWriter, request *http.Request) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request

	return c
}