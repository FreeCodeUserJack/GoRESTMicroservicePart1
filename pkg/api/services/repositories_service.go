package services

import (
	"net/http"
	"sync"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/config"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/github"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/providers/github_provider"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

type repoService struct {
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
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

func (r *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	resultChan := make(chan repositories.CreateRepositoriesResult)
	outChan := make(chan repositories.CreateReposResponse)
	defer close(outChan)

	var wg sync.WaitGroup

	go r.handleRepoResults(&wg, resultChan, outChan)

	for _, req := range requests {
		wg.Add(1)
		go r.CreateRepoConcurrent(req, resultChan)
	}

	wg.Wait()
	close(resultChan)

	result := <- outChan

	successCreations := 0
	for _, res := range result.Results {
		if res.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.GetStatus()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (r *repoService) handleRepoResults(wg *sync.WaitGroup, resultChan chan repositories.CreateRepositoriesResult, outChan chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for result := range resultChan {
		results.Results = append(results.Results, result)
		wg.Done()
	}

	outChan <- results
}

func (r *repoService) CreateRepoConcurrent(request repositories.CreateRepoRequest, outChan chan repositories.CreateRepositoriesResult) {
	if err := request.Validate(); err != nil {
		outChan <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	result, err := r.CreateRepo(request)

	if err != nil {
		outChan <- repositories.CreateRepositoriesResult{
			Response: nil,
			Error: err,
		}
		return
	}

	outChan <- repositories.CreateRepositoriesResult{
		Response: result,
		Error: nil,
	}
}