package providers

import (
	"context"
	"fmt"
	"magnifin/internal/app/model"
)

func (s *ProviderService) Delete(ctx context.Context, connection *model.Connection) error {
	providerUser, err := s.providerUserRepository.GetByID(ctx, connection.ProviderUserID)
	if err != nil {
		return fmt.Errorf("unable to get provider user: %w", err)
	}

	provider, err := s.providerRepository.GetByID(ctx, providerUser.ProviderID)
	if err != nil {
		return fmt.Errorf("unable to get provider: %w", err)
	}

	providerPort := s.ports[provider.Name]
	if providerPort == nil {
		return fmt.Errorf("provider port not found")
	}

	err = providerPort.DeleteConnection(ctx, provider, providerUser, connection)
	if err != nil {
		return fmt.Errorf("unable to delete connection on provider's side: %w", err)
	}

	return nil
}
