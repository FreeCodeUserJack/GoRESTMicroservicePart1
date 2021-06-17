package oauth

import (
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/oauth-api/domain/oauth"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/oauth-api/services"
	"github.com/gin-gonic/gin"
)


func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(request); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	if private := c.GetHeader("X-Private"); private != "true" {
		c.String(http.StatusUnauthorized, "You do not have permission to call this endpoint")
		return
	}

	tokenId := c.Param("token_id")

	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusOK, token)
}