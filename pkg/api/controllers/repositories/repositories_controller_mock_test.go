package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/services"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/test_utils"
)

var (
	funcCreateRepo func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type RepositoriesServiceMock struct {
}

func (r RepositoriesServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(request)
}

func (r RepositoriesServiceMock) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(requests)
}

func TestCreateRepoGithubErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &RepositoriesServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewApiError(http.StatusUnauthorized, "Requires authentication")
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "NewRepo"}`))

	c := test_utils.GetMockedContext(response, request)

	CreateRepo(c)

	if http.StatusUnauthorized != response.Code {
		t.Errorf("got %v, want %v", response.Code, http.StatusUnauthorized)
	}

	fmt.Println(response.Body.String())

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	if err != nil {
		t.Fatalf("expected no err but got %v while unmarshalling", err)
	}

	if apiErr == nil {
		t.Fatal("expected apiErr obj but didn't get one")
	}

	if apiErr.GetStatus() != http.StatusUnauthorized {
		t.Errorf("got %v, want %v", apiErr.GetStatus(), http.StatusUnauthorized)
	}

	if apiErr.GetMessage() != "Requires authentication" {
		t.Errorf("got %v, want %v", apiErr.GetMessage(), "Requires authentication")
	}
}

func TestCreateRepoNoErrorMockingEntireService(t *testing.T) {
	services.RepositoryService = &RepositoriesServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id: 123,
			Name: "MockName",	
			Owner: "MockOwner",
		}, nil
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "MockName"}`))

	c := test_utils.GetMockedContext(response, request)

	CreateRepo(c)

	if http.StatusCreated != response.Code {
		t.Errorf("got %v, want %v", response.Code, http.StatusCreated)
	}

	fmt.Println(response.Body.String())

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	if err != nil {
		t.Fatalf("expected valid result body but error umarshalling: %v", err)
	}

	if result.Id != 123 {
		t.Errorf("got %v, want %v", result.Name, 123)
	}
}