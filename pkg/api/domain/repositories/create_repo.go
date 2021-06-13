package repositories

import (
	"strings"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)

// this is request sent by client of our microservice
type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (c *CreateRepoRequest) Validate() errors.ApiError {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return errors.NewBadRequestApiError("invalid repository name")
	}
	return nil
}

// this is what we will give back to client using our microservice
type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status"`
	Results []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error errors.ApiError `json:"error"`
}