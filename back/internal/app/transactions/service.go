package transactions

import (
	"context"
	"magnifin/internal/app/model"
	"time"
)

type CounterpartyEnrichmentResult struct {
	CounterpartyName *string
	Method           *string
}

type TransactionsRepository interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
	GetByID(ctx context.Context, id int32) (*model.Transaction, error)
	StoreEnrichedData(ctx context.Context, data *model.TransactionEnrichment) error
	ListAllUserCounterpartiesByTransID(ctx context.Context, transID int32) ([]string, error)
}

type Enricher interface {
	CleanCounterpartyName(ctx context.Context, name *string, userCounterparties []string) (*CounterpartyEnrichmentResult, error)
	GetCounterpartyNameLogoURL(ctx context.Context, name *string) (*string, error)
}

type Service struct {
	TransactionsRepository TransactionsRepository
	Enricher               Enricher
}

func NewTransactionsService(
	transactionsRepository TransactionsRepository,
	enricher Enricher,
) *Service {
	return &Service{
		TransactionsRepository: transactionsRepository,
		Enricher:               enricher,
	}
}
