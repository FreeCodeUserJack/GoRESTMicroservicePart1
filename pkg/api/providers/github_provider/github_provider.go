package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/clients/restclient"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/github"
)


const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	
	response, err := restclient.Post(urlCreateRepo, request, headers)

	// this error if internet is gone, i.e. github api is not called at all
	if err != nil {
		log.Printf("error when trying to create new repo in github: %v", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid response body from github",
		}
	}

	defer response.Body.Close()

	// error 300+ means error from github
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err = json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "could not unmarshal the body bytes",
			}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse

	// valid request
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Printf("error when trying to unmarshal valid github response body: %v", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "Error unmarshalling body",
		}
	}

	return &result, nil
}