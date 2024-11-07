package providers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"
	"time"
)

func (s *ProviderService) HandleSyncError(ctx context.Context, connectionID int32, syncErr error) error {
	slog.Warn(fmt.Sprintf("handling sync error for connection %d: %s", connectionID, syncErr))
	connection, err := s.connectionRepository.GetByID(ctx, connectionID)
	if err != nil {
		return fmt.Errorf("unable to get connection: %w", err)
	} else if connection == nil {
		return errors.New("connection not found")
	}

	errorMsg := syncErr.Error()
	if errors.Is(syncErr, model.ErrRateLimited) {
		connection.Status = model.ConnectionStatusRateLimited
		errorMsg = "Rate limited, the account will be synchronized later"
	} else {
		connection.Status = model.ConnectionStatusSuspended
	}

	connection.ErrorMessage = &errorMsg
	_, err = s.connectionRepository.Update(ctx, connection)
	if err != nil {
		return fmt.Errorf("unable to update connection: %w", err)
	}

	return nil
}

func (s *ProviderService) SynchronizeConnection(ctx context.Context, connectionID int32) error {
	connection, err := s.connectionRepository.GetByID(ctx, connectionID)
	if err != nil {
		return fmt.Errorf("unable to get connection: %w", err)
	} else if connection == nil {
		return errors.New("connection not found")
	}

	providerUserID, err := s.providerUserRepository.GetByID(ctx, connection.ProviderUserID)
	if err != nil {
		return fmt.Errorf("unable to get provider user: %w", err)
	} else if providerUserID == nil {
		return errors.New("provider user not found")
	}

	provider, err := s.providerRepository.GetByID(ctx, providerUserID.ProviderID)
	if err != nil {
		return fmt.Errorf("unable to get provider: %w", err)
	} else if provider == nil {
		return errors.New("provider not found")
	}

	providerPort := s.ports[provider.Name]
	if providerPort == nil {
		return errors.New("provider port not found")
	}

	slog.Info("Importing accounts")
	accounts, err := s.syncAccounts(ctx, connectionID, providerPort, provider, providerUserID, connection)
	if err != nil {
		return fmt.Errorf("unable to import accounts: %w", err)
	}

	slog.Info("Importing transactions")
	for _, account := range accounts {
		err = s.syncTransactions(ctx, providerPort, provider, providerUserID, connection, &account)
		if err != nil {
			return fmt.Errorf("unable to import transactions: %w", err)
		}
	}

	slog.Info("Connection synchronized")
	connection.Status = model.ConnectionStatusSynchronized
	now := time.Now()
	connection.LastSuccessfulSync = &now
	connection.ErrorMessage = nil
	_, err = s.connectionRepository.Update(ctx, connection)
	if err != nil {
		return fmt.Errorf("unable to update connection: %w", err)
	}

	return nil
}

func (s *ProviderService) syncAccounts(
	ctx context.Context,
	connectionID int32,
	providerPort ProviderPort,
	provider *model.Provider,
	providerUserID *model.ProviderUser,
	connection *model.Connection,
) ([]model.Account, error) {
	accounts, err := providerPort.GetAccounts(ctx, provider, providerUserID, connection)
	if err != nil {
		return nil, fmt.Errorf("unable to get accounts: %w", err)
	} else if accounts == nil {
		return nil, errors.New("no accounts on the connection, weird")
	}

	dbAccounts := make([]model.Account, len(accounts))
	for i, account := range accounts {
		savedAccount, err := s.accountsRepository.GetByConnectionIDAndProviderAccountID(ctx, connectionID, account.ProviderAccountID)
		if err != nil {
			return nil, fmt.Errorf("unable to get account: %w", err)
		}

		if savedAccount == nil {
			slog.Debug("Creating account")
			savedAccount, err = s.accountsRepository.Create(ctx, &account)
		} else {
			slog.Debug("Updating account")
			account.ID = savedAccount.ID
			savedAccount, err = s.accountsRepository.Update(ctx, &account)
		}

		if err != nil {
			return nil, fmt.Errorf("unable to save account: %w", err)
		}

		dbAccounts[i] = *savedAccount
	}
	return dbAccounts, nil
}

func (s *ProviderService) syncTransactions(
	ctx context.Context,
	providerPort ProviderPort,
	provider *model.Provider,
	providerUserID *model.ProviderUser,
	connection *model.Connection,
	account *model.Account,
) error {
	transactions, err := providerPort.GetTransactions(ctx, provider, providerUserID, connection, account)
	if err != nil {
		return fmt.Errorf("unable to get transactions: %w", err)
	}

	for _, transaction := range transactions {
		dbTransaction, err := s.transactionsRepository.GetByAccountIDAndProviderTransactionID(ctx, account.ID, transaction.ProviderTransactionID)
		if err != nil {
			return fmt.Errorf("unable to get transaction: %w", err)
		}

		if dbTransaction == nil {
			_, err = s.transactionsRepository.Create(ctx, &transaction)
		} else {
			transaction.ID = dbTransaction.ID
			_, err = s.transactionsRepository.Update(ctx, &transaction)
		}

		if err != nil {
			return fmt.Errorf("unable to save transaction: %w", err)
		}
	}

	return nil
}
