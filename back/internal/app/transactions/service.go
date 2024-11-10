package transactions

import (
	"context"
	"magnifin/internal/app/model"
	"time"
)

type TransactionsRepository interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
}

type Service struct {
	TransactionsRepository TransactionsRepository
}

func NewTransactionsService(transactionsRepository TransactionsRepository) *Service {
	return &Service{
		TransactionsRepository: transactionsRepository,
	}
}
