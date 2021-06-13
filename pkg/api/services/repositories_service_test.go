package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/clients/restclient"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name: "",
		Description: "",
	}

	response, err := RepositoryService.CreateRepo(request)

	if err == nil {
		t.Fatal("expected err but didn't get one!")
	}

	if err.GetMessage() != "invalid repository name" {
		t.Errorf("got: %v, want: %v", err.GetMessage(), "invalid repository name")
	}

	if err.GetStatus() != http.StatusBadRequest {
		t.Errorf("got: %v, want: %v", err.GetStatus(), http.StatusBadRequest)
	}

	if response != nil {
		t.Errorf("expected nil response but got: %v", response)
	}
}

func TestCreateRepoErrorFromGithub(t *testing.T) {

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

	request := repositories.CreateRepoRequest{
		Name: "Sample Repo Name",
		Description: "Sample Repo Description",
	}

	response, err := RepositoryService.CreateRepo(request)

	if err == nil {
		t.Fatal("wanted error but got none")
	}

	if err.GetStatus() != http.StatusInternalServerError {
		t.Errorf("got: %v, want: %v", err.GetStatus(), http.StatusInternalServerError)
	}

	if err.GetMessage() != "invalid response body from github" {
		t.Errorf("got %v, want %v", err.GetMessage(), "invalid response body from github")
	}

	if response != nil {
		t.Errorf("response should be nil but is instead: %v", response)
	}
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"name": "RepoName"}`)),
		},
		BodyText: `{"id": 123, "name": "RepoName", "owner": {"login": "federico"}}`,
	})

	request := repositories.CreateRepoRequest{
		Name: "RepoName",
		Description: "RepoDescription",
	}

	response, err := RepositoryService.CreateRepo(request)

	if err != nil {
		t.Fatalf("didn't want error but got %v", err)
	}

	if response == nil {
		t.Fatal("wanted response but didn't get one")
	}

	if response.Name != request.Name {
		t.Errorf("got %v, want %v", response.Name, request.Name)
	}
}