package connectors

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	SearchByName(ctx context.Context, name string) ([]model.Connector, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
