package connectors

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	SearchByName(ctx context.Context, name string) ([]model.Connector, error)
	Connect(ctx context.Context, user *model.User, connectorID int32, params *model.ConnectParams) (*model.ConnectInstruction, error)
	ConnectCallback(ctx context.Context, connectorID int32, sid string, providerConnectionID *string) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
