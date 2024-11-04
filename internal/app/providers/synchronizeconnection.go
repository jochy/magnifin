package providers

import (
	"context"
	"errors"
	"log/slog"
)

func (s *ProviderService) SynchronizeConnection(ctx context.Context, connectionID int32) error {
	connection, err := s.connectionRepository.GetByID(ctx, connectionID)
	if err != nil {
		return err
	} else if connection == nil {
		return errors.New("connection not found")
	}

	providerUserID, err := s.providerUserRepository.GetByID(ctx, connection.ProviderUserID)
	if err != nil {
		return err
	} else if providerUserID == nil {
		return errors.New("provider user not found")
	}

	provider, err := s.providerRepository.GetByID(ctx, providerUserID.ProviderID)
	if err != nil {
		return err
	} else if provider == nil {
		return errors.New("provider not found")
	}

	providerPort := s.ports[provider.Name]
	if providerPort == nil {
		return errors.New("provider port not found")
	}

	accounts, err := providerPort.GetAccounts(ctx, provider, providerUserID, connection)
	if err != nil {
		return err
	} else if accounts == nil {
		return errors.New("no accounts on the connection, weird")
	}

	for _, account := range accounts {
		savedAccount, err := s.accountsRepository.GetByConnectionIDAndProviderAccountID(ctx, connectionID, account.ProviderAccountID)
		if err != nil {
			return err
		}

		if savedAccount == nil {
			slog.Debug("Creating account")
			_, err = s.accountsRepository.Create(ctx, &account)
		} else {
			slog.Debug("Updating account")
			account.ID = savedAccount.ID
			_, err = s.accountsRepository.Update(ctx, &account)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
