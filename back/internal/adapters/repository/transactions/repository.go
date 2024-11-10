package transactions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"time"
)

type Repository struct {
	db database.Service
}

func NewRepository(db database.Service) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByAccountIDAndProviderTransactionID(ctx context.Context, accountID int32, providerTransactionID string) (*model.Transaction, error) {
	transaction, err := r.db.FindTransactionByAccountIDAndProviderTransactionID(ctx, database.FindTransactionByAccountIDAndProviderTransactionIDParams{
		AccountID:             accountID,
		ProviderTransactionID: providerTransactionID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(transaction, nil), nil
}

func (r *Repository) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	trs, err := r.db.CreateTransaction(ctx, database.CreateTransactionParams{
		AccountID:             transaction.AccountID,
		ProviderTransactionID: transaction.ProviderTransactionID,
		BankTransactionID:     repository.ToSqlNullString(transaction.BankTransactionID),
		Amount:                fmt.Sprintf("%f", transaction.Amount),
		Currency:              transaction.Currency,
		Direction:             string(transaction.Direction),
		Status:                string(transaction.Status),
		OperationAt:           transaction.OperationAt,
		CounterpartyName:      repository.ToSqlNullString(transaction.CounterpartyName),
		CounterpartyAccount:   repository.ToSqlNullString(transaction.CounterpartyAccount),
		Reference:             repository.ToSqlNullString(transaction.Reference),
	})
	if err != nil {
		return nil, err
	}

	return toDomain(trs, nil), nil
}

func (r *Repository) Update(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	trs, err := r.db.UpdateTransaction(ctx, database.UpdateTransactionParams{
		ID:                    transaction.ID,
		ProviderTransactionID: transaction.ProviderTransactionID,
		BankTransactionID:     repository.ToSqlNullString(transaction.BankTransactionID),
		Amount:                fmt.Sprintf("%f", transaction.Amount),
		Currency:              transaction.Currency,
		Direction:             string(transaction.Direction),
		Status:                string(transaction.Status),
		OperationAt:           transaction.OperationAt,
		CounterpartyName:      repository.ToSqlNullString(transaction.CounterpartyName),
		CounterpartyAccount:   repository.ToSqlNullString(transaction.CounterpartyAccount),
		Reference:             repository.ToSqlNullString(transaction.Reference),
	})
	if err != nil {
		return nil, err
	}

	return toDomain(trs, nil), nil
}

func (r *Repository) DeleteByConnectionID(ctx context.Context, connectionID int32) error {
	err := r.db.DeleteTransactionsEnrichmentsByConnectionID(ctx, connectionID)
	if err != nil {
		return err
	}

	err = r.db.DeleteTransactionsByConnectionID(ctx, connectionID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllByUserBetweenDates(ctx context.Context, user *model.User, from time.Time, to time.Time) ([]model.Transaction, error) {
	transactions, err := r.db.GetTransactionsByUserIDAndBetweenDates(ctx, database.GetTransactionsByUserIDAndBetweenDatesParams{
		UserID:        user.ID,
		OperationAt:   from,
		OperationAt_2: to,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting transactions: %w", err)
	}

	trs := make([]model.Transaction, len(transactions))
	for i, tx := range transactions {
		trs[i] = *toEnrichedDomain(&tx)
	}

	return trs, nil
}

func (r *Repository) GetTransactionMinMaxDateByUser(ctx context.Context, user *model.User) (*model.TransactionMinAndMax, error) {
	row, err := r.db.GetTransactionsMinAndMaxDateByUserID(ctx, user.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting transactions min and max date: %w", err)
	}

	// check if row.minDate is time.Time
	if _, ok := (row.MinDate).(time.Time); !ok {
		return nil, fmt.Errorf("error getting transactions min and max date: %w", err)
	} else if _, ok := (row.MaxDate).(time.Time); !ok {
		return nil, fmt.Errorf("error getting transactions min and max date: %w", err)
	}

	return &model.TransactionMinAndMax{
		Min: (row.MinDate).(time.Time), //nolint:forcetypeassert
		Max: (row.MaxDate).(time.Time), //nolint:forcetypeassert
	}, nil
}
