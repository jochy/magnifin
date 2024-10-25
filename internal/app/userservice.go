package app

import (
	"context"
	"magnifin/internal/adapters/repository/users"
	"magnifin/internal/app/model"
)

type UserService struct {
	userRepository users.Repository
}

func NewUserService(userRepository users.Repository) *UserService {
	return &UserService{
		userRepository: userRepository,
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
