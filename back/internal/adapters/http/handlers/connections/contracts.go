package connections

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	ListConnections(ctx context.Context, user *model.User) ([]model.ConnectionWithAccounts, error)
	DeleteConnection(ctx context.Context, user *model.User, id int32) error
}

type Handlers struct {
	Service Service
}

func NewHandlers(service Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}
