package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetMockedContext(t *testing.T) {
	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:123/something", nil)

	request.Header = http.Header{"X-Mock": {"true"}}

	c := GetMockedContext(response, request)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("x-mock"))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}