package accounts

import (
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"strconv"
)

func toDomainAccount(account database.Account) *model.Account {
	balance, err := strconv.ParseFloat(account.Balance, 64)
	if err != nil {
		balance = 0
	}

	return &model.Account{
		ID:                account.ID,
		ConnectionID:      account.ConnectionID,
		ProviderAccountID: account.ProviderAccountID,
		Name:              repository.FromSqlNullString(account.Name),
		Type:              repository.FromSqlNullString(account.Type),
		Currency:          repository.FromSqlNullString(account.Currency),
		AccountNumber:     repository.FromSqlNullString(account.AccountNumber),
		Balance:           balance,
		BankAccountID:     repository.FromSqlNullString(account.BankAccountID),
	}
}
