package transactions

import (
	"context"
	"errors"
	"fmt"
	"magnifin/internal/app/model"
)

func (s *Service) EnrichTransaction(ctx context.Context, transactionID int32) error {
	trs, err := s.TransactionsRepository.GetByID(ctx, transactionID)
	if err != nil {
		return fmt.Errorf("enrichTransaction: %w", err)
	} else if trs == nil {
		return errors.New("enrichTransaction: transaction not found")
	}

	userCounterparties, err := s.TransactionsRepository.ListAllUserCounterpartiesByTransID(ctx, trs.ID)
	if err != nil {
		return fmt.Errorf("enrichTransaction failed to get user counterparties: %w", err)
	}

	enrichedData := &model.TransactionEnrichment{TransactionID: trs.ID}
	if trs.Enrichment != nil {
		enrichedData = trs.Enrichment
	}

	hasEnriched := false

	if trs.CounterpartyName != nil && enrichedData.CounterpartyName == nil {
		cleanName, err := s.Enricher.CleanCounterpartyName(ctx, trs.CounterpartyName, userCounterparties)
		if err != nil {
			return fmt.Errorf("enrichTransaction failed to clean the name: %w", err)
		}

		if cleanName != nil {
			hasEnriched = true
			enrichedData.CounterpartyName = cleanName.CounterpartyName
			enrichedData.Method = cleanName.Method
		}
	}

	var counterpartyName *string
	counterpartyName = enrichedData.UserCounterpartyName
	if counterpartyName == nil {
		counterpartyName = enrichedData.CounterpartyName
	}

	if counterpartyName != nil {
		logoURL, err := s.Enricher.GetCounterpartyNameLogoURL(ctx, counterpartyName)
		if err != nil {
			return fmt.Errorf("enrichTransaction failed to get the logo: %w", err)
		}

		if logoURL != nil {
			hasEnriched = true
			enrichedData.CounterpartyLogo = logoURL.ID

			if _, err := s.ImageRepository.Store(ctx, &model.Image{
				ID:          *logoURL.ID,
				Content:     *logoURL.Content,
				ContentType: *logoURL.ContentType,
			}); err != nil {
				return fmt.Errorf("enrichTransaction failed to store the logo: %w", err)
			}
		}
	}

	if enrichedData.Category == nil {
		catID, err := s.ComputeCategory(ctx, trs, enrichedData)
		if err != nil {
			return fmt.Errorf("enrichTransaction failed to compute the category: %w", err)
		}

		if catID != nil {
			hasEnriched = true
			enrichedData.Category = catID
		}
	}

	if hasEnriched {
		if _, err := s.TransactionsRepository.StoreEnrichedData(ctx, enrichedData); err != nil {
			return fmt.Errorf("enrichTransaction failed to save the enriched data: %w", err)
		}
	}

	return nil
}
