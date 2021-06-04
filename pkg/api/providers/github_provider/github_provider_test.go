package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/clients/restclient"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/github"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	if headerAuthorization != "Authorization" {
		t.Errorf("got %v, want %v", headerAuthorization, "Authorization")
	}

	if headerAuthorizationFormat != "token %s" {
		t.Errorf("got %v, want %v", headerAuthorizationFormat, "token %s")
	}

	if urlCreateRepo != "https://api.github.com/user/repos" {
		t.Errorf("got %v, want %v", urlCreateRepo, "https://api.github.com/user/repos")
	}
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")

	if header != "token " + "abc123" {
		t.Errorf("got %v, want %v", header, "token " + "abc123")
	}
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	t.Run("invalid request should return err", func(t *testing.T) {
		restclient.FlushMocks()

		restclient.AddMockup(restclient.Mock{
			Url: urlCreateRepo,
			HttpMethod: http.MethodPost,
			Response: nil,
			Error: errors.New("invalid restclient response"),
		})

		response, err := CreateRepo("", github.CreateRepoRequest{})

		if err == nil {
			t.Fatal("expected error but didn't get one")
		}

		if err != nil && err.Message != "invalid restclient response" {
			t.Errorf("got: %v, want: %s", err.Message, "invalid restclient response")
		}

		if response != nil {
			t.Errorf("response should be nil but got: %v", response)
		}
	})
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMocks()

	invalidCloser, _ := os.Open("-blah")
	defer invalidCloser.Close()

	restclient.AddMockup(restclient.Mock{
		Url: urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	if err == nil {
		t.Fatal("expected error but didn't get one")
	}

	if err != nil && err.Message != "invalid response body from github" {
		t.Errorf("got: %v, want: %s", err.Message, "invalid response body from github")
	}

	if http.StatusInternalServerError != err.StatusCode {
		t.Errorf("got %v, want %v", err.StatusCode, http.StatusInternalServerError)
	}

	if response != nil {
		t.Errorf("response should be nil but got: %v", response)
	}
}

// when you call github api without auth then you get response containing "message" and "documentation_url" keys
	// when you unmarshal this msg and it throws an err (we mimic this by making message be an int)
		// {"message": 1}
func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMocks()

	restclient.AddMockup(restclient.Mock{
		Url: urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	if err == nil {
		t.Fatal("expected error but didn't get one")
	}

	if err != nil && err.Message != "could not unmarshal the body bytes" {
		t.Errorf("got: %v, want: %s", err.Message, "could not unmarshal the body bytes")
	}

	if http.StatusInternalServerError != err.StatusCode {
		t.Errorf("got %v, want %v", err.StatusCode, http.StatusInternalServerError)
	}

	if response != nil {
		t.Errorf("response should be nil but got: %v", response)
	}
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMocks()

	restclient.AddMockup(restclient.Mock{
		Url: urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	if err == nil {
		t.Fatal("expected error but didn't get one")
	}

	if err != nil && err.Message != "Requires authentication" {
		t.Errorf("got: %v, want: %s", err.Message, "Requires authentication")
	}

	if http.StatusUnauthorized != err.StatusCode {
		t.Errorf("got %v, want %v", err.StatusCode, http.StatusUnauthorized)
	}

	if response != nil {
		t.Errorf("response should be nil but got: %v", response)
	}
}

// when the 201 created response returned from github isn't properly mapped to our CreateRepoResponse struct in domain dir
func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMocks()

	restclient.AddMockup(restclient.Mock{
		Url: urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "Hello-World"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	if err == nil {
		t.Fatal("expected error but didn't get one")
	}

	if err != nil && err.Message != "Error unmarshalling body" {
		t.Errorf("got: %v, want: %s", err.Message, "Error unmarshalling body")
	}

	if http.StatusInternalServerError != err.StatusCode {
		t.Errorf("got %v, want %v", err.StatusCode, http.StatusInternalServerError)
	}

	if response != nil {
		t.Errorf("response should be nil but got: %v", response)
	}
}

func TestCreateRepoSuccess(t *testing.T) {
	restclient.FlushMocks()

	restclient.AddMockup(restclient.Mock{
		Url: urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"name": "Hello-World"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	if err != nil {
		t.Fatalf("didn't expect error but got one: %v", err)
	}

	if response == nil {
		t.Fatal("response should not be nil but is nil")
	}
}

// func TestDefer(t *testing.T) {
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// }