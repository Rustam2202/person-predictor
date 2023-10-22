package persons

import (
	"context"
	"net/http"
	"net/http/httptest"
	"person-predicator/internal/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Create a mock service
	mockService := &MockPersonService{}

	// Create a test handler instance
	handler := &PersonHandler{
		service: &service.PersonService{},
	}

	// Define your API endpoint and attach the handler
	router.DELETE("/users", handler.Delete)

	// Create a test request to the API endpoint
	req, _ := http.NewRequest(http.MethodDelete, "/users?id=123", nil)

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(recorder, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the service method calls
	assert.True(t, mockService.DeleteCalled)
	assert.Equal(t, int64(123), mockService.DeleteID)
}

type MockPersonService struct {
	DeleteCalled bool
	DeleteID     int64
}

func (m *MockPersonService) Delete(ctx context.Context, id int64) error {
	m.DeleteCalled = true
	m.DeleteID = id
	return nil // or return an error if needed
}
