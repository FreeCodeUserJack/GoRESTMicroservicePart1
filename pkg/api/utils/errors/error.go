package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)


type ApiError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
}

func (e *apiError) GetStatus() int {
	return e.Status
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) GetError() string {
	return e.Err
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Status: http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerApiError(message string) ApiError {
	return &apiError{
		Status: http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestApiError(message string) ApiError {
	return &apiError{
		Status: http.StatusBadRequest,
		Message: message,
	}
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		Status: statusCode,
		Message: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}