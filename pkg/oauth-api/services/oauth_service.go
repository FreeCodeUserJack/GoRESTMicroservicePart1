package services

import (
	"time"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/oauth-api/domain/oauth"
)

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

type oauthService struct {
}

func (o *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// validate against database
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token := &oauth.AccessToken{
		UserId: user.Id,
		Expires: time.Now().UTC().Add(time.Hour * 24).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	return token, nil
}

func (o *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	if token.IsNotExpired() {
		// refreshses token expiration
		token.Expires = time.Now().UTC().Add(time.Hour * 24).Unix()
		token.Save()
		return token, nil
	} else {
		return nil, errors.NewNotFoundApiError("token expired")
	}
}