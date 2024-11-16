package transactions

import (
	"context"
	"errors"
	"fmt"
	"magnifin/internal/app/model"
	"time"
)

func (s *Service) ApplyCategoryRule(ctx context.Context, ruleID int32, userID int32) error {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user by id: %w", err)
	} else if user == nil {
		return errors.New("user not found")
	}

	rule, err := s.CategoryRepository.GetCategoryRuleByID(ctx, ruleID)
	if err != nil {
		return fmt.Errorf("error getting category rule by id: %w", err)
	} else if rule == nil {
		return errors.New("category rule not found")
	}

	minMax, err := s.TransactionsRepository.GetTransactionMinMaxDateByUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error getting min and max transaction dates: %w", err)
	} else if minMax == nil {
		return errors.New("no transactions found")
	}

	minDate := time.Date(minMax.Min.Year(), minMax.Min.Month(), 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(minMax.Max.Year(), minMax.Max.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)

	current := minDate
	for current.Before(maxDate) {
		transactions, err := s.TransactionsRepository.GetAllByUserBetweenDates(ctx, user, current, current.AddDate(0, 1, 0))
		if err != nil {
			return fmt.Errorf("error getting transactions by user between dates: %w", err)
		}

		for _, trs := range transactions {
			if trs.Enrichment == nil {
				trs.Enrichment = &model.TransactionEnrichment{
					TransactionID: trs.ID,
				}
			}

			keywords := transactionKeywords(&trs, trs.Enrichment)
			if evaluateRule(*rule, keywords) {
				trs.Enrichment.Category = &rule.CategoryID
				t, err := s.TransactionsRepository.Update(ctx, &trs)
				if err != nil {
					return fmt.Errorf("error updating transaction: %w", err)
				}

				s.Notifier.Notify(userID, t)
			}
		}

		current = current.AddDate(0, 1, 0)
	}

	return nil
}
