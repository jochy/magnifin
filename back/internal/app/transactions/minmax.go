package transactions

import (
	"context"
	"fmt"
	"magnifin/internal/app/model"
)

func (s *Service) GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error) {
	minmax, err := s.TransactionsRepository.GetTransactionMinMaxDateByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("transactions.Service.GetTransactionMinMaxDateByUser: %w", err)
	}

	return minmax, nil
}
