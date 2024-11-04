package connectors

import (
	"context"
	"magnifin/internal/app/model"
)

type Repository interface {
	SearchByName(ctx context.Context, name string) ([]model.Connector, error)
	LikeSearchByName(ctx context.Context, name string) ([]model.Connector, error)
	GetByID(ctx context.Context, id int32) (*model.Connector, error)
}

type ProviderService interface {
	Connect(ctx context.Context, user *model.User, connector *model.Connector, params *model.ConnectParams) (*model.ConnectInstruction, error)
	ConnectCallback(ctx context.Context, user *model.User, connector *model.Connector, connectionID string, providerConnectionID *string) error
}

type Service struct {
	repository      Repository
	providerService ProviderService
}

func NewConnectorService(repository Repository, providerService ProviderService) *Service {
	return &Service{
		repository:      repository,
		providerService: providerService,
	}
}

func (s *Service) GetByID(ctx context.Context, id int32) (*model.Connector, error) {
	return s.repository.GetByID(ctx, id)
}
