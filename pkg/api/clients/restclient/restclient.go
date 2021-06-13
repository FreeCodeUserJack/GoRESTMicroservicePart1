package restclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var (
	enabledMocks = false
	mocks = make(map[string]*Mock)
)

type Mock struct {
	Url string
	HttpMethod string
	Response *http.Response
	BodyText string
	Error error
}

func (m Mock) GetMockId() string {
	return fmt.Sprintf("%s_%s", m.HttpMethod, m.Url)
}

func (m Mock) GetResponse() *http.Response {
	var res *http.Response

	if m.BodyText == `` {
		invalidCloser, _ := os.Open("something")
		res = &http.Response{
			StatusCode: m.Response.StatusCode,
			Body: invalidCloser,
		}
	} else {
		res = &http.Response{
			StatusCode: m.Response.StatusCode,
			Body: ioutil.NopCloser(strings.NewReader(m.BodyText)),
		}
	}
	return res
}

func StartMockups() {
	enabledMocks = true
}

func StopMockups() {
	enabledMocks = false
}

func AddMockup(mock Mock) {
	mocks[mock.GetMockId()] = &mock
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}


func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		// return local without calling any external resource
		mock := mocks[http.MethodPost + "_" + url]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Error
	}

	var bytes []byte
	var jsonErr error

	if reflect.ValueOf(body).Kind() != reflect.String {
		bytes, jsonErr = json.Marshal(body)

		if jsonErr != nil {
			fmt.Printf("error occurred marshalling! %v\n", body)
			return nil, jsonErr
		}
	}

	request, reqErr := http.NewRequest(http.MethodPost, url, strings.NewReader(string(bytes)))

	if reqErr != nil {
		fmt.Printf("error occurred creating new request: %v\n", reqErr)
		return nil, reqErr
	}

	request.Header = headers

	client := http.Client{}

	response, respErr := client.Do(request)

	if respErr != nil {
		fmt.Printf("error occurred for client.Do: %v", respErr)
		return nil, respErr
	}

	return response, nil
}