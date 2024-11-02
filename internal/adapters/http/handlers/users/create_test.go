package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"magnifin/internal/adapters/http/handlers/users/mocks"
	"magnifin/internal/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    createRequest
		mockService    func(t *testing.T) *mocks.Service
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			name: "successful creation",
			requestBody: createRequest{
				Username: "testuser",
				Password: "password",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				mockService := mocks.NewService(t)
				mockService.EXPECT().Create(mock.Anything, "testuser", "password").Return(&model.User{Username: "testuser"}, nil).Once()
				mockService.EXPECT().GenerateJWT(mock.Anything, &model.User{Username: "testuser"}).Return("token", nil).Once()
				return mockService
			},
			expectedStatus: http.StatusOK,
			expectedBody:   gin.H{"token": "token"},
		},
		{
			name: "creation error",
			requestBody: createRequest{
				Username: "testuser",
				Password: "password",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				mockService := mocks.NewService(t)
				mockService.EXPECT().Create(mock.Anything, "testuser", "password").Return(nil, errors.New("creation error")).Once()
				return mockService
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"internal_error": "creation error"},
		},
		{
			name: "JWT generation error",
			requestBody: createRequest{
				Username: "testuser",
				Password: "password",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				mockService := mocks.NewService(t)
				mockService.EXPECT().Create(mock.Anything, "testuser", "password").Return(&model.User{Username: "testuser"}, nil).Once()
				mockService.EXPECT().GenerateJWT(mock.Anything, &model.User{Username: "testuser"}).Return("", errors.New("JWT error")).Once()
				return mockService
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"internal_error": "JWT error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := tt.mockService(t)
			handler := NewCreateHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, _ := json.Marshal(tt.requestBody)
			c.Request, _ = http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")

			handler.Handle(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			var responseBody gin.H
			_ = json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.Equal(t, tt.expectedBody, responseBody)

			mockService.AssertExpectations(t)
		})
	}
}
