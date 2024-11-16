package transactions

import (
	"context"
	"fmt"
	"magnifin/internal/adapters/jobs"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/scheduler"
)

func (s *Service) Update(ctx context.Context, id int32, category *int32, userCounterparty *string) (*model.Transaction, error) {
	tx, err := s.TransactionsRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var enrichment *model.TransactionEnrichment
	if tx.Enrichment != nil {
		enrichment = tx.Enrichment
	} else {
		enrichment = &model.TransactionEnrichment{}
	}

	enrichment.Category = category

	if userCounterparty != nil && ((enrichment.CounterpartyName != nil && hasChanged(enrichment.CounterpartyName, userCounterparty)) || (enrichment.CounterpartyName == nil && hasChanged(tx.CounterpartyName, userCounterparty))) {
		enrichment.UserCounterpartyName = userCounterparty
	}

	trs, err := s.TransactionsRepository.Update(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	// Re-enrich, in case the user has provided more data
	_ = scheduler.Scheduler.Trigger(context.Background(), jobs.TransactionEnrichInput{TransactionID: id})

	return trs, nil
}

func hasChanged(old *string, n *string) bool {
	if old == nil && n == nil {
		return false
	}

	if old == nil || n == nil {
		return true
	}

	return *old != *n
}
