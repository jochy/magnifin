package providers

import (
	"context"
	"errors"
	"magnifin/internal/app/model"
)

type ProviderRepository interface {
	List(ctx context.Context) ([]model.Provider, error)
	Update(ctx context.Context, provider *model.Provider) (*model.Provider, error)
	GetByName(ctx context.Context, name string) (*model.Provider, error)
}

type ConnectorRepository interface {
	Upsert(ctx context.Context, connectors *model.Connector) (*model.Connector, error)
}

type ProviderPort interface {
	Name() string
	ValidateConfiguration(provider *model.Provider) error
	ListConnectors(ctx context.Context, provider *model.Provider) ([]model.Connector, error)
}

type ProviderService struct {
	providerRepository  ProviderRepository
	connectorRepository ConnectorRepository
	ports               map[string]ProviderPort
}

func NewProviderService(
	repository ProviderRepository,
	connectorRepository ConnectorRepository,
	ports []ProviderPort,
) *ProviderService {
	p := make(map[string]ProviderPort)
	for _, port := range ports {
		p[port.Name()] = port
	}

	return &ProviderService{
		providerRepository:  repository,
		connectorRepository: connectorRepository,
		ports:               p,
	}
}

func (s *ProviderService) ListProviders(ctx context.Context) ([]model.Provider, error) {
	return s.providerRepository.List(ctx)
}

func (s *ProviderService) UpdateProvider(ctx context.Context, provider model.Provider) (*model.Provider, error) {
	port := s.ports[provider.Name]
	if port == nil {
		return nil, errors.New("provider not found")
	}

	err := port.ValidateConfiguration(&provider)
	if err != nil {
		return nil, err
	}

	return s.providerRepository.Update(ctx, &provider)
}
