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
	GetByID(ctx context.Context, id int32) (*model.Provider, error)
}

type ConnectorRepository interface {
	Upsert(ctx context.Context, connectors *model.Connector) (*model.Connector, error)
}

type ProviderUserRepository interface {
	GetByProviderIDAndUserID(ctx context.Context, providerID int32, userID int32) (*model.ProviderUser, error)
	Save(ctx context.Context, providerID int32, userID int32, providerUserID string) (*model.ProviderUser, error)
}

type ProviderPort interface {
	Name() string
	ValidateConfiguration(provider *model.Provider) error
	ListConnectors(ctx context.Context, provider *model.Provider) ([]model.Connector, error)
	CreateProviderUser(ctx context.Context, provider *model.Provider, user *model.User) (*model.ProviderUser, error)
	Connect(ctx context.Context, provider *model.Provider, providerUser *model.ProviderUser, connector *model.Connector, params *model.ConnectParams) (*model.ConnectInstruction, error)
}

type ProviderService struct {
	providerRepository     ProviderRepository
	connectorRepository    ConnectorRepository
	providerUserRepository ProviderUserRepository
	ports                  map[string]ProviderPort
}

func NewProviderService(
	repository ProviderRepository,
	connectorRepository ConnectorRepository,
	providerUserRepository ProviderUserRepository,
	ports []ProviderPort,
) *ProviderService {
	p := make(map[string]ProviderPort)
	for _, port := range ports {
		p[port.Name()] = port
	}

	return &ProviderService{
		providerRepository:     repository,
		connectorRepository:    connectorRepository,
		providerUserRepository: providerUserRepository,
		ports:                  p,
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

	// Update the bank list in background
	go func() {
		s.UpdateConnectorsList(context.Background())
	}()

	return s.providerRepository.Update(ctx, &provider)
}
