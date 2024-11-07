package connections

import (
	"context"
	"fmt"
	"magnifin/internal/app/model"
)

func (s *Service) DeleteConnection(ctx context.Context, user *model.User, id int32) error {
	// First: make sure the connection exists and belongs to the user
	cnx, err := s.ConnectionsRepository.GetByIDAndUser(ctx, id, user)
	if err != nil {
		return fmt.Errorf("error getting connection: %w", err)
	} else if cnx == nil {
		return fmt.Errorf("connection not found")
	}

	// Second: delete the connection on the provider side
	err = s.ProviderService.Delete(ctx, cnx)
	if err != nil {
		return fmt.Errorf("error deleting connection on provider's side: %w", err)
	}

	// Third: delete the connection on our side
	err = s.ConnectionsRepository.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting connection: %w", err)
	}
	err = s.AccountsRepository.DeleteByConnectionID(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting accounts: %w", err)
	}
	err = s.TransactionsRepository.DeleteByConnectionID(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting transactions: %w", err)
	}

	return nil
}
