package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/clients/restclient"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/test_utils"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()

	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))

	c := test_utils.GetMockedContext(response, request)

	CreateRepo(c)

	if http.StatusBadRequest != response.Code {
		t.Errorf("got %v, want %v", response.Code, http.StatusBadRequest)
	}

	fmt.Println(response.Body.String())

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	if err != nil {
		t.Fatalf("expected no err but got %v while unmarshalling", err)
	}

	if apiErr == nil {
		t.Fatal("expected apiErr obj but didn't get one")
	}

	if apiErr.GetStatus() != http.StatusBadRequest {
		t.Errorf("got %v, want %v", apiErr.GetStatus(), http.StatusBadRequest)
	}

	if apiErr.GetMessage() != "invalid json body" {
		t.Errorf("got %v, want %v", apiErr.GetMessage(), "invalid json body")
	}
}

func TestCreateRepoGithubError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser((strings.NewReader(`{"message": "Requires authentication"}`))),
		},
	})

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

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser((strings.NewReader(`{"id": 123}`))),
		},
	})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))

	c := test_utils.GetMockedContext(response, request)

	CreateRepo(c)

	if http.StatusCreated != response.Code {
		t.Errorf("got %v, want %v", response.Code, http.StatusCreated)
	}

	fmt.Println(response.Body.String())

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	if  err != nil {
		t.Fatalf("expected valid result body but error umarshalling: %v", err)
	}

	if result.Id != 123 {
		t.Errorf("got %v, want %v", result.Name, 123)
	}
}