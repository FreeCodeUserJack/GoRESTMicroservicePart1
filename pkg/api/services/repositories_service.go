package services

import (
	"strings"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/config"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/github"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/providers/github_provider"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

type repoService struct {
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestApiError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name: input.Name,
		Private: false,
		Description: input.Description,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
		Id: response.Id,
		Owner: response.Owner.Login,
		Name: response.Name,
	}, nil
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}