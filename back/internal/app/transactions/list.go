package transactions

import (
	"context"
	"fmt"
	"magnifin/internal/app/model"
	"time"
)

func (s *Service) GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error) {
	transactions, err := s.TransactionsRepository.GetAllByUserBetweenDates(ctx, user, from, to)
	if err != nil {
		return nil, fmt.Errorf("error getting transactions: %w", err)
	}

	return transactions, nil
}
