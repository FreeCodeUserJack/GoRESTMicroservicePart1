package oauth

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)


const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?"
)

// temporary storage, will be replaced by DB access later
var (
	users = map[string]*User{
		"John": {Id: 123, Username: "John"},
		"Jane": {Id: 234, Username: "Jane"},
		"Jack": {Id: 345, Username: "Jack"},
	}
)

func GetUserByUsernameAndPassword(username, password string) (*User, errors.ApiError) {
	fmt.Println(queryGetUserByUsernameAndPassword)

	user, ok := users[username]
	if ok {
		return user, nil
	} else {
		return nil, errors.NewNotFoundApiError(fmt.Sprintf("user with username %s not found", username))
	}
}