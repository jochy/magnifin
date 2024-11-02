package app

import (
	"context"
	"fmt"
	"magnifin/internal/adapters/repository/users"
	"magnifin/internal/app/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	userRepository users.Repository
	jwtSignKey     string
}

func NewUserService(userRepository users.Repository, jwtSignKey string) *UserService {
	return &UserService{
		userRepository: userRepository,
		jwtSignKey:     jwtSignKey,
	}
}

func (s *UserService) Login(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := s.userRepository.GetUserByUsernameAndPassword(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := s.userRepository.CreateUser(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GenerateJWT(_ context.Context, user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": strconv.Itoa(int(user.ID)),
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(s.jwtSignKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) FromJWT(ctx context.Context, token string) (*model.User, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSignKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userIDStr, err := t.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.GetUserByID(ctx, int32(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
