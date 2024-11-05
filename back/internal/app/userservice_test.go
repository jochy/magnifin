package app

import (
	"context"
	"magnifin/internal/app/mocks"
	"magnifin/internal/app/model"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Login(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo, "test_sign_key")

	expectedUser := &model.User{ID: 1, Username: "testuser"}
	mockRepo.On("GetUserByUsernameAndPassword", mock.Anything, "testuser", "password").Return(expectedUser, nil)

	user, err := userService.Login(context.Background(), "testuser", "password")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_Create(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo, "test_sign_key")

	expectedUser := &model.User{ID: 1, Username: "newuser"}
	mockRepo.On("CreateUser", mock.Anything, "newuser", "password").Return(expectedUser, nil)

	user, err := userService.Create(context.Background(), "newuser", "password")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserService_GenerateJWT(t *testing.T) {
	userService := NewUserService(nil, "test_sign_key")

	user := &model.User{ID: 1, Username: "testuser"}
	token, err := userService.GenerateJWT(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("test_sign_key"), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
}

func TestUserService_FromJWT(t *testing.T) {
	mockRepo := mocks.NewUserRepository(t)
	userService := NewUserService(mockRepo, "test_sign_key")

	user := &model.User{ID: 1, Username: "testuser"}
	token, _ := userService.GenerateJWT(context.Background(), user)

	mockRepo.On("GetUserByID", mock.Anything, int32(1)).Return(user, nil)

	parsedUser, err := userService.FromJWT(context.Background(), token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}

func TestUserService_FromJWT_InvalidToken(t *testing.T) {
	userService := NewUserService(nil, "test_sign_key")

	_, err := userService.FromJWT(context.Background(), "invalid_token")
	assert.Error(t, err)
}
