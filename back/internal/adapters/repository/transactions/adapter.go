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
		ID:                   enrichment.ID,
		TransactionID:        enrichment.TransactionID,
		CounterpartyLogo:     repository.FromSqlNullString(enrichment.CounterpartyLogo),
		Category:             repository.FromSqlNullInt32(enrichment.Category),
		CounterpartyName:     repository.FromSqlNullString(enrichment.CounterpartyName),
		Reference:            repository.FromSqlNullString(enrichment.Reference),
		Method:               repository.FromSqlNullString(enrichment.Method),
		UserCounterpartyName: repository.FromSqlNullString(enrichment.UserCounterpartyName),
	}
}

func toEnrichedDomain(trs *database.GetTransactionsByUserIDAndBetweenDatesRow) *model.Transaction {
	var enrichment *model.TransactionEnrichment
	if trs.ID_2.Valid {
		enrichment = &model.TransactionEnrichment{
			ID:                   trs.ID,
			TransactionID:        *repository.FromSqlNullInt32(trs.TransactionID),
			CounterpartyLogo:     repository.FromSqlNullString(trs.CounterpartyLogo),
			Category:             repository.FromSqlNullInt32(trs.Category),
			CounterpartyName:     repository.FromSqlNullString(trs.CounterpartyName_2),
			Reference:            repository.FromSqlNullString(trs.Reference_2),
			Method:               repository.FromSqlNullString(trs.Method),
			UserCounterpartyName: repository.FromSqlNullString(trs.UserCounterpartyName),
		}
	}

	amount, err := strconv.ParseFloat(trs.Amount, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to parse transaction amount: %s", err))
		amount = 0
	}

	return &model.Transaction{
		ID:                    trs.ID,
		AccountID:             trs.AccountID,
		ProviderTransactionID: trs.ProviderTransactionID,
		BankTransactionID:     repository.FromSqlNullString(trs.BankTransactionID),
		Amount:                amount,
		Currency:              trs.Currency,
		Direction:             model.TransactionDirection(trs.Direction),
		Status:                model.TransactionStatus(trs.Status),
		OperationAt:           trs.OperationAt,
		CounterpartyName:      repository.FromSqlNullString(trs.CounterpartyName),
		CounterpartyAccount:   repository.FromSqlNullString(trs.CounterpartyAccount),
		Reference:             repository.FromSqlNullString(trs.Reference),
		Enrichment:            enrichment,
	}
}

func toEnrichedDomainByID(trs database.GetTransactionByIDRow) *model.Transaction {
	var enrichment *model.TransactionEnrichment
	if trs.ID_2.Valid {
		enrichment = &model.TransactionEnrichment{
			ID:                   trs.ID,
			TransactionID:        *repository.FromSqlNullInt32(trs.TransactionID),
			CounterpartyLogo:     repository.FromSqlNullString(trs.CounterpartyLogo),
			Category:             repository.FromSqlNullInt32(trs.Category),
			CounterpartyName:     repository.FromSqlNullString(trs.CounterpartyName_2),
			Reference:            repository.FromSqlNullString(trs.Reference_2),
			Method:               repository.FromSqlNullString(trs.Method),
			UserCounterpartyName: repository.FromSqlNullString(trs.UserCounterpartyName),
		}
	}

	amount, err := strconv.ParseFloat(trs.Amount, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to parse transaction amount: %s", err))
		amount = 0
	}

	return &model.Transaction{
		ID:                    trs.ID,
		AccountID:             trs.AccountID,
		ProviderTransactionID: trs.ProviderTransactionID,
		BankTransactionID:     repository.FromSqlNullString(trs.BankTransactionID),
		Amount:                amount,
		Currency:              trs.Currency,
		Direction:             model.TransactionDirection(trs.Direction),
		Status:                model.TransactionStatus(trs.Status),
		OperationAt:           trs.OperationAt,
		CounterpartyName:      repository.FromSqlNullString(trs.CounterpartyName),
		CounterpartyAccount:   repository.FromSqlNullString(trs.CounterpartyAccount),
		Reference:             repository.FromSqlNullString(trs.Reference),
		Enrichment:            enrichment,
	}
}
