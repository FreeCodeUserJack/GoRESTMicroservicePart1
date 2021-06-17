package oauth

import (
	"strings"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)


type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// pointer receiver type so we can actuall trimspace on the username/password
func (a *AccessTokenRequest) Validate() errors.ApiError {
	a.Username = strings.TrimSpace(a.Username)
	if a.Username == "" {
		return errors.NewBadRequestApiError("invalid username")
	}

	a.Password = strings.TrimSpace(a.Password)
	if a.Password == "" {
		return errors.NewBadRequestApiError("invalid password")
	}

	return nil
}