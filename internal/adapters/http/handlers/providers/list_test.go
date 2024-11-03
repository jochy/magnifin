package providers

import (
	"errors"
	"magnifin/internal/adapters/http/handlers/providers/mocks"
	"magnifin/internal/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		mockListResponse   []model.Provider
		mockListError      error
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "success",
			mockListResponse: []model.Provider{
				{ID: 1, Name: "Provider1", Enabled: true},
				{ID: 2, Name: "Provider2", Enabled: false},
			},
			mockListError:      nil,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"providers":[{"id":1,"name":"Provider1","enabled":true},{"id":2,"name":"Provider2","enabled":false}]}`,
		},
		{
			name:               "error",
			mockListResponse:   nil,
			mockListError:      errors.New("some error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"some error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mocks.NewService(t)
			handler := NewHandler(mockService)

			mockService.On("ListProviders", mock.Anything).Return(tt.mockListResponse, tt.mockListError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			handler.List(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.JSONEq(t, tt.expectedResponse, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}
