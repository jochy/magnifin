package transactions

import (
	"context"
	"magnifin/internal/adapters"
	"magnifin/internal/app/model"
	"time"
)

type Service interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
	Update(ctx context.Context, id int32, category *int32, userCounterparty *string) (*model.Transaction, error)
}

type MapperInterface interface {
	ToPublicFormat(trs *model.Transaction) *adapters.EnrichedTransaction
}

type Handlers struct {
	Service Service
	Mapper  MapperInterface
}

func NewHandlers(service Service, mapper MapperInterface) *Handlers {
	return &Handlers{
		Service: service,
		Mapper:  mapper,
	}
}
