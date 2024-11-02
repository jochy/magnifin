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

func TestLoginHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		mockService    func(t *testing.T) *mocks.Service
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:        "Invalid JSON",
			requestBody: "invalid",
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				m := mocks.NewService(t)
				return m
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Login error",
			requestBody: loginRequest{
				Username: "user",
				Password: "pass",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				m := mocks.NewService(t)
				m.EXPECT().Login(mock.Anything, "user", "pass").Return(nil, errors.New("login error"))
				return m
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Invalid username or password",
			requestBody: loginRequest{
				Username: "user",
				Password: "pass",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				m := mocks.NewService(t)
				m.EXPECT().Login(mock.Anything, "user", "pass").Return(nil, nil)
				return m
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "JWT generation error",
			requestBody: loginRequest{
				Username: "user",
				Password: "pass",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				m := mocks.NewService(t)
				m.EXPECT().Login(mock.Anything, "user", "pass").Return(&model.User{ID: 1}, nil)
				m.EXPECT().GenerateJWT(mock.Anything, &model.User{ID: 1}).Return("", errors.New("jwt error"))
				return m
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Successful login",
			requestBody: loginRequest{
				Username: "user",
				Password: "pass",
			},
			mockService: func(t *testing.T) *mocks.Service {
				t.Helper()
				m := mocks.NewService(t)
				m.EXPECT().Login(mock.Anything, "user", "pass").Return(&model.User{ID: 1}, nil)
				m.EXPECT().GenerateJWT(mock.Anything, &model.User{ID: 1}).Return("token", nil)
				return m
			},
			expectedStatus: http.StatusOK,
			expectedBody:   loginResponse{Token: "token"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := tt.mockService(t)
			handler := NewHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			body, _ := json.Marshal(tt.requestBody)
			c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")

			handler.Login(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != nil {
				var responseBody loginResponse
				_ = json.Unmarshal(w.Body.Bytes(), &responseBody)
				assert.Equal(t, tt.expectedBody, responseBody)
			}
		})
	}
}
