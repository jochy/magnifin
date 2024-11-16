package transactions

import (
	"context"
	"magnifin/internal/app/model"
	"time"
)

type Service interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
	Update(ctx context.Context, id int32, category *int32, userCounterparty *string) (*model.Transaction, error)
}

type Handlers struct {
	Service   Service
	PublicURL string
}

func NewHandlers(service Service, publicURL string) *Handlers {
	return &Handlers{
		Service:   service,
		PublicURL: publicURL,
	}
}
