package transactions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
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
