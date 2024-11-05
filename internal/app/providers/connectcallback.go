package providers

import (
	"context"
	"errors"
	"log/slog"
	"magnifin/internal/app/model"
)

func (s *ProviderService) ConnectCallback(
	ctx context.Context,
	user *model.User,
	connector *model.Connector,
	sid string,
	providerConnectionID *string,
) error {
	redirectSession, err := s.redirectSessionsRepository.GetRedirectSessionByID(ctx, sid)
	if err != nil {
		return err
	} else if redirectSession == nil {
		return errors.New("redirect session not found")
	}

	provider, err := s.providerRepository.GetByID(ctx, connector.ProviderID)
	if err != nil {
		return err
	} else if provider == nil {
		return errors.New("provider not found in db")
	}

	providerUser, err := s.providerUserRepository.GetByProviderIDAndUserID(ctx, provider.ID, user.ID)
	if err != nil {
		return err
	} else if providerUser == nil {
		return errors.New("provider user not found in db")
	}

	providerConnectionIDToUse := providerConnectionID
	if providerConnectionIDToUse == nil {
		providerConnectionIDToUse = redirectSession.ProviderConnectionID
	}

	if providerConnectionIDToUse == nil {
		return errors.New("provider connection ID not found")
	}

	providerConnection, err := s.ports[provider.Name].GetConnectionByID(ctx, provider, providerUser, connector, *providerConnectionIDToUse)
	if err != nil {
		return err
	} else if providerConnection == nil {
		return errors.New("provider connection not found")
	}

	providerConnection.Status = model.ConnectionStatusSyncInProgress

	var savedConnection *model.Connection
	if redirectSession.InternalConnectionID != nil {
		savedConnection, err = s.connectionRepository.GetByID(ctx, *redirectSession.InternalConnectionID)
	} else {
		savedConnection, err = s.connectionRepository.GetByProviderUserIDAndProviderConnectionID(ctx, providerUser.ID, redirectSession.ID)
	}

	if err != nil {
		return err
	}

	if savedConnection != nil {
		providerConnection.ID = savedConnection.ID
		savedConnection, err = s.connectionRepository.Update(ctx, providerConnection)
	} else {
		savedConnection, err = s.connectionRepository.Create(ctx, providerConnection)
	}

	if err != nil {
		return err
	}

	slog.Info("Connection upserted " + string(savedConnection.ID))

	return nil
}
