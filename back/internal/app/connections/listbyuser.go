package connections

import (
	"context"
	"fmt"
	"magnifin/internal/app/model"
)

func (s *Service) ListConnections(ctx context.Context, user *model.User) ([]model.ConnectionWithAccounts, error) {
	// Could be handled better with one single query, which should improve performances.
	// But I want to go fast for now... so I'll fix it later.
	if user == nil {
		return nil, fmt.Errorf("user is required")
	}

	connections, err := s.ConnectionsRepository.ListActiveByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to list connections: %w", err)
	}

	connectionsWithAccounts := make([]model.ConnectionWithAccounts, len(connections))
	for i, connection := range connections {
		accounts, err := s.AccountsRepository.ListByConnection(ctx, connection.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to list accounts: %w", err)
		}

		connector, err := s.ConnectorsRepository.GetByID(ctx, connection.ConnectorID)
		if err != nil {
			return nil, fmt.Errorf("failed to get connector: %w", err)
		}

		connectionsWithAccounts[i] = model.ConnectionWithAccounts{
			Connection: &connection,
			Accounts:   accounts,
			Connector:  connector,
		}
	}

	return connectionsWithAccounts, nil
}
