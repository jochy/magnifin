package users

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	Create(ctx context.Context, username string, password string) (*model.User, error)
	Login(ctx context.Context, username string, password string) (*model.User, error)
	GenerateJWT(ctx context.Context, user *model.User) (string, error)
}

type loginResponse struct {
	Token string `json:"token"`
}
