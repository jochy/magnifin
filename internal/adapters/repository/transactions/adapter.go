package transactions

import (
	"fmt"
	"log/slog"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"strconv"
)

func toDomain(transaction database.Transaction, enrichment *database.TransactionEnrichment) *model.Transaction {
	amount, err := strconv.ParseFloat(transaction.Amount, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to parse transaction amount: %s", err))
		amount = 0
	}

	return &model.Transaction{
		ID:                    transaction.ID,
		AccountID:             transaction.AccountID,
		ProviderTransactionID: transaction.ProviderTransactionID,
		BankTransactionID:     repository.FromSqlNullString(transaction.BankTransactionID),
		Amount:                amount,
		Currency:              transaction.Currency,
		Direction:             model.TransactionDirection(transaction.Direction),
		Status:                model.TransactionStatus(transaction.Status),
		OperationAt:           transaction.OperationAt,
		CounterpartyName:      repository.FromSqlNullString(transaction.CounterpartyName),
		CounterpartyAccount:   repository.FromSqlNullString(transaction.CounterpartyAccount),
		Reference:             repository.FromSqlNullString(transaction.Reference),
		Enrichment:            enrichmentToDomain(enrichment),
	}
}

func enrichmentToDomain(enrichment *database.TransactionEnrichment) *model.TransactionEnrichment {
	if enrichment == nil {
		return nil
	}

	return &model.TransactionEnrichment{
		ID:                  enrichment.ID,
		TransactionID:       enrichment.TransactionID,
		CounterpartyLogoURL: repository.FromSqlNullString(enrichment.CounterpartyLogoUrl),
		Category:            repository.FromSqlNullString(enrichment.Category),
		CounterpartyName:    repository.FromSqlNullString(enrichment.CounterpartyName),
		Reference:           repository.FromSqlNullString(enrichment.Reference),
	}
}
