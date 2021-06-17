package oauth

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)

// mocked here - should be cassandra or DB connection save
var (
	tokens = make(map[string]*AccessToken)
)

func (a *AccessToken) Save() errors.ApiError {
	a.AccessToken = fmt.Sprintf("USER_%d", a.UserId)
	tokens[a.AccessToken] = a
	return nil
}

func GetAccessToken(accessToken string) (*AccessToken, errors.ApiError) {
	token, ok := tokens[accessToken]
	if !ok {
		return nil, errors.NewNotFoundApiError("token not valid")
	}
	return token, nil
}