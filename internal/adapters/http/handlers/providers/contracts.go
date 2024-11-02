package providers

import (
	"context"
	"magnifin/internal/app/model"
)

type Service interface {
	List(ctx context.Context) ([]model.Provider, error)
	Update(ctx context.Context, provider model.Provider) (*model.Provider, error)
}

type providerResponse struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func toProviderResponse(provider model.Provider) providerResponse {
	return providerResponse{
		ID:      provider.ID,
		Name:    provider.Name,
		Enabled: provider.Enabled,
	}
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
