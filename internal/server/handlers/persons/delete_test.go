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
	router := gin.Default()

	mockService := &MockPersonService{}

	handler := &PersonHandler{
		service: &service.PersonService{},
	}

	router.DELETE("/person", handler.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/person?id=123", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

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
