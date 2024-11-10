package transactions

import (
	"context"
	"magnifin/internal/app/model"
	"time"
)

type Service interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
}

type Handlers struct {
	Service Service
}

func NewHandlers(service Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}
