package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/clients/restclient"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	outChan := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, outChan)

	result := <- outChan

	assert.NotNil(t, result)

	if result.Response != nil {
		t.Errorf("response should be nil but is: %v", result.Response)
	}

	if result.Error == nil {
		t.Fatal("expected error but didn't get one")
	}

	assert.EqualValues(t, result.Error.GetStatus(), http.StatusBadRequest)
	assert.EqualValues(t, result.Error.GetMessage(), "invalid repository name")
}

func TestCreateRepoConcurrentGithubError(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name: "blah",
	}

	invalidBody, _ := os.Open("something")

	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidBody,
		},
	})

	outChan := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, outChan)

	result := <- outChan

	assert.NotNil(t, result)

	if result.Response != nil {
		t.Errorf("response should be nil but is: %v", result.Response)
	}

	if result.Error == nil {
		t.Fatal("expected error but didn't get one")
	}

	assert.EqualValues(t, result.Error.GetStatus(), http.StatusInternalServerError)
	assert.EqualValues(t, result.Error.GetMessage(), "invalid response body from github")
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name: "RepoName",
	}

	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "RepoName"}`)),
		},
		BodyText: `{"id": 123, "name": "RepoName", "owner": {"login": "federico"}}`,
	})

	outChan := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.CreateRepoConcurrent(request, outChan)

	result := <- outChan

	assert.NotNil(t, result)

	if result.Response == nil {
		t.Fatal("expected response but didn't get one")
	}

	if result.Error != nil {
		t.Errorf("expected no error but got one: %v", result.Error)
	}

	assert.EqualValues(t, result.Response.Name, "RepoName")
	assert.EqualValues(t, result.Response.Id, 123)
}

func TestHandleRepoResults(t *testing.T) {
	var wg sync.WaitGroup
	resultChan := make(chan repositories.CreateRepositoriesResult)
	outChan := make(chan repositories.CreateReposResponse)

	go func() {
		for i := 0; i < 5; i++ {
			resultChan <- repositories.CreateRepositoriesResult{
				Response: nil,
				Error: errors.NewApiError(http.StatusInternalServerError, "invalid json body"),
			}
		}
	}()

	wg.Add(5)

	service := repoService{}
	go service.handleRepoResults(&wg, resultChan, outChan)

	wg.Wait()
	close(resultChan)

	result := <-outChan

	assert.NotNil(t, result)
	assert.EqualValues(t, len(result.Results), 5)
	assert.EqualValues(t, http.StatusInternalServerError, result.Results[0].Error.GetStatus())
}

func TestCreateReposInvalidRequests(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "    "},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.GetStatus())
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.GetStatus())
	assert.Nil(t, result.Results[0].Response)
	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.GetMessage())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.GetMessage())
}

func TestCreateReposPartialRequests(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "Test Repo"},
	}

	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "RepoName"}`)),
		},
		BodyText: `{"id": 123, "name": "RepoName", "owner": {"login": "federico"}}`,
	})

	result, err := RepositoryService.CreateRepos(requests)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, res := range result.Results {
		if res.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, res.Error.GetStatus())
			assert.Nil(t, res.Response)
			assert.EqualValues(t, "invalid repository name", res.Error.GetMessage())
		} else {
			assert.Nil(t, res.Error)
			assert.EqualValues(t, 123, res.Response.Id)
			assert.EqualValues(t, "RepoName", res.Response.Name)
		}
	}
}

func TestCreateReposValidRequests(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "RepoName", "owner": {"login": "federico"}}`)),
		},
		BodyText: `{"id": 123, "name": "RepoName", "owner": {"login": "federico"}}`,
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "RepoName", result.Results[0].Response.Name)

	assert.EqualValues(t, 123, result.Results[1].Response.Id)
	assert.EqualValues(t, "RepoName", result.Results[1].Response.Name)
}