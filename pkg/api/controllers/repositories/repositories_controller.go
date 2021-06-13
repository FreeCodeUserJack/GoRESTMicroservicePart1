package repositories

import (
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/services"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/gin-gonic/gin"
)


func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)

	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}