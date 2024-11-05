package middlewares

import (
	"errors"
	"magnifin/internal/adapters/http/middlewares/mocks"
	"magnifin/internal/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthMiddleware_Authenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		userService    func(t *testing.T) *mocks.UserService
		authHeader     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "missing Authorization header",
			userService: func(t *testing.T) *mocks.UserService {
				t.Helper()
				m := mocks.NewUserService(t)
				return m
			},
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"missing Authorization header"}`,
		},
		{
			name: "invalid token",
			userService: func(t *testing.T) *mocks.UserService {
				t.Helper()
				m := mocks.NewUserService(t)
				m.EXPECT().FromJWT(mock.Anything, "invalid-token").Return(nil, nil)
				return m
			},
			authHeader:     "invalid-token",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"invalid token"}`,
		},
		{
			name: "technical error",
			userService: func(t *testing.T) *mocks.UserService {
				t.Helper()
				m := mocks.NewUserService(t)
				m.EXPECT().FromJWT(mock.Anything, "invalid-token").Return(nil, errors.New("technical error"))
				return m
			},
			authHeader:     "invalid-token",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"technical error"}`,
		},
		{
			name: "valid token",
			userService: func(t *testing.T) *mocks.UserService {
				t.Helper()
				m := mocks.NewUserService(t)
				m.EXPECT().FromJWT(mock.Anything, "valid-token").Return(&model.User{}, nil)
				return m
			},
			authHeader:     "valid-token",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := NewAuthMiddleware(tt.userService(t))
			router := gin.New()
			router.Use(middleware.Authenticate)
			router.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
