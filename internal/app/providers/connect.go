package providers

import (
	"context"
	"errors"
	"magnifin/internal/app/model"
)

func (s *ProviderService) Connect(ctx context.Context, user *model.User, connector *model.Connector, params *model.ConnectParams) (*model.ConnectInstruction, error) {
	provider, err := s.providerRepository.GetByID(ctx, connector.ProviderID)
	if err != nil {
		return nil, err
	} else if provider == nil {
		return nil, errors.New("provider not found in db")
	}

	providerPort := s.ports[provider.Name]
	if providerPort == nil {
		return nil, errors.New("provider port not found")
	}

	if !provider.Enabled {
		return nil, errors.New("provider is not enabled")
	}

	// First: create provider user if needed
	providerUser, err := s.providerUserRepository.GetByProviderIDAndUserID(ctx, provider.ID, user.ID)
	if err != nil {
		return nil, err
	}

	if providerUser == nil {
		providerUser, err = providerPort.CreateProviderUser(ctx, provider, user)
		if err != nil {
			return nil, err
		}

		providerUser, err = s.providerUserRepository.Save(ctx, provider.ID, user.ID, providerUser.ProviderUserID)
		if err != nil {
			return nil, err
		}
	}

	// Second: create connection
	connectInstruction, err := providerPort.Connect(ctx, provider, providerUser, connector, params)
	if err != nil {
		return nil, err
	}

	return connectInstruction, nil
}
