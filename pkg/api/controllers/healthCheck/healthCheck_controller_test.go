package healthCheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "good", health)
}

func TestHandleHealthCheck(t *testing.T) {
	response := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/health", nil)

	c := test_utils.GetMockedContext(response, request)

	HandleHealthCheck(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "good", response.Body.String())
}