package app

import (
	"context"
	"magnifin/internal/app/model"
)

type ProviderRepository interface {
	List(ctx context.Context) ([]model.Provider, error)
	Update(ctx context.Context, provider *model.Provider) (*model.Provider, error)
}

type ProviderService struct {
	repository ProviderRepository
}

func NewProviderService(repository ProviderRepository) *ProviderService {
	return &ProviderService{repository}
}

func (s *ProviderService) List(ctx context.Context) ([]model.Provider, error) {
	return s.repository.List(ctx)
}

func (s *ProviderService) Update(ctx context.Context, provider model.Provider) (*model.Provider, error) {
	return s.repository.Update(ctx, &provider)
}
