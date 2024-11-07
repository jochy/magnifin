package accounts

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

func (r *Repository) GetByConnectionIDAndProviderAccountID(ctx context.Context, connectionID int32, providerAccountID string) (*model.Account, error) {
	account, err := r.db.GetAccountByConnectionIDAndProviderAccountID(ctx, database.GetAccountByConnectionIDAndProviderAccountIDParams{
		ConnectionID:      connectionID,
		ProviderAccountID: providerAccountID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomainAccount(account), err
}

func (r *Repository) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	accountEntity, err := r.db.CreateAccount(ctx, database.CreateAccountParams{
		ConnectionID:      account.ConnectionID,
		ProviderAccountID: account.ProviderAccountID,
		Name:              repository.ToSqlNullString(account.Name),
		Type:              repository.ToSqlNullString(account.Type),
		Currency:          repository.ToSqlNullString(account.Currency),
		AccountNumber:     repository.ToSqlNullString(account.AccountNumber),
		Balance:           fmt.Sprintf("%f", account.Balance),
		BankAccountID:     repository.ToSqlNullString(account.BankAccountID),
	})
	if err != nil {
		return nil, err
	}

	return toDomainAccount(accountEntity), nil
}

func (r *Repository) Update(ctx context.Context, account *model.Account) (*model.Account, error) {
	accountEntity, err := r.db.UpdateAccount(ctx, database.UpdateAccountParams{
		ID:                account.ID,
		ProviderAccountID: account.ProviderAccountID,
		Name:              repository.ToSqlNullString(account.Name),
		Type:              repository.ToSqlNullString(account.Type),
		Currency:          repository.ToSqlNullString(account.Currency),
		AccountNumber:     repository.ToSqlNullString(account.AccountNumber),
		Balance:           fmt.Sprintf("%f", account.Balance),
		BankAccountID:     repository.ToSqlNullString(account.BankAccountID),
	})
	if err != nil {
		return nil, err
	}

	return toDomainAccount(accountEntity), nil
}

func (r *Repository) ListByConnection(ctx context.Context, connectionID int32) ([]model.Account, error) {
	accounts, err := r.db.ListAccountsByConnectionID(ctx, connectionID)
	if err != nil {
		return nil, err
	}

	domainAccounts := make([]model.Account, len(accounts))
	for i, a := range accounts {
		domainAccounts[i] = *toDomainAccount(a)
	}

	return domainAccounts, nil
}

func (r *Repository) DeleteByConnectionID(ctx context.Context, connectionID int32) error {
	err := r.db.DeleteAccountByConnectionID(ctx, connectionID)
	if err != nil {
		return err
	}

	return nil
}
