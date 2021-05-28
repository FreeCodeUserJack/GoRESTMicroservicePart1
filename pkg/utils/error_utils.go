package utils

import (
	"encoding/json"
)


type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

func (a ApplicationError) String() string {
	jsonValue, _ := json.Marshal(a)
	return string(jsonValue)
}

func (a ApplicationError) Error() string {
	// return fmt.Sprintf("Application errors:\nMessage: %s\nStatusCode: %d\nCode: %s", a.Message, a.StatusCode, a.Code)
	jsonValue, _ := json.Marshal(a)
	return string(jsonValue)
}