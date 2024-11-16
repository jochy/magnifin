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

type CounterpartyLogoEnrichmentResult struct {
	ID          *string
	Content     *string
	ContentType *string
}

type TransactionsRepository interface {
	GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error)
	GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error)
	GetByID(ctx context.Context, id int32) (*model.Transaction, error)
	StoreEnrichedData(ctx context.Context, data *model.TransactionEnrichment) (*model.TransactionEnrichment, error)
	ListAllUserCounterpartiesByTransID(ctx context.Context, transID int32) ([]string, error)
	Update(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	GetUserIDByTransactionID(ctx context.Context, id int32) (int32, error)
}

type CategoryRepository interface {
	GetAllRulesByUserFromTransID(ctx context.Context, transID int32) ([]model.CategoryRule, error)
	GetAllCategoriesByUserFromTransactionID(ctx context.Context, transID int32) ([]model.Category, error)
}

type ImageRepository interface {
	GetByID(ctx context.Context, id string) (*model.Image, error)
	Store(ctx context.Context, image *model.Image) (*model.Image, error)
}

type Enricher interface {
	CleanCounterpartyName(ctx context.Context, name *string, userCounterparties []string) (*CounterpartyEnrichmentResult, error)
	GetCounterpartyNameLogoURL(ctx context.Context, name *string) (*CounterpartyLogoEnrichmentResult, error)
	GuessCategory(ctx context.Context, keywords []string, categories []string) (*string, error)
}

type Notifier interface {
	Notify(userID int32, trs *model.Transaction)
}

type Service struct {
	TransactionsRepository TransactionsRepository
	CategoryRepository     CategoryRepository
	ImageRepository        ImageRepository
	Enricher               Enricher
	Notifier               Notifier
}

func NewTransactionsService(
	transactionsRepository TransactionsRepository,
	categoryRepository CategoryRepository,
	imageRepository ImageRepository,
	enricher Enricher,
	notifier Notifier,
) *Service {
	return &Service{
		TransactionsRepository: transactionsRepository,
		CategoryRepository:     categoryRepository,
		ImageRepository:        imageRepository,
		Enricher:               enricher,
		Notifier:               notifier,
	}
}
