package transactions

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"magnifin/internal/app/model"
	"slices"
	"strconv"
	"strings"
)

func (s *Service) ComputeCategory(ctx context.Context, transaction *model.Transaction, enrichment *model.TransactionEnrichment) (*int32, error) {
	if transaction == nil || enrichment == nil {
		return nil, errors.New("computeCategory: transaction or enrichment is nil")
	}

	rules, err := s.CategoryRepository.GetAllRulesByUserFromTransID(ctx, transaction.ID)
	if err != nil {
		return nil, err
	}

	transactionKeywords := transactionKeywords(transaction, enrichment)

	for _, rule := range rules {
		if evaluateRule(rule, transactionKeywords) {
			return &rule.CategoryID, nil
		}
	}

	slog.Debug(fmt.Sprintf("No category found for transaction %d, let's use AI to guess one", transaction.ID))
	userCategories, err := s.CategoryRepository.GetAllCategoriesByUserFromTransactionID(ctx, transaction.ID)
	if err != nil {
		return nil, fmt.Errorf("computeCategory: failed to get user categories: %w", err)
	}

	aiCategory, err := s.Enricher.GuessCategory(ctx, transactionKeywords, categoryNames(userCategories))
	if err != nil {
		return nil, fmt.Errorf("computeCategory: failed to guess the category: %w", err)
	}

	if aiCategory == nil {
		// No category guessed, return nil
		return nil, nil
	}

	category := getCategoryByName(userCategories, *aiCategory)
	if category == nil {
		slog.Debug(fmt.Sprintf("Category %s, guessed by AI is not found", *aiCategory))
		return nil, nil
	}

	return &category.ID, nil
}

func evaluateRule(rule model.CategoryRule, transactionKeywords []string) bool {
	if len(rule.Rule) == 0 {
		return false
	}

	for _, keyword := range rule.Rule {
		if !slices.Contains(transactionKeywords, strings.ToLower(keyword)) {
			nb, err := isNumber(keyword)
			if err == nil {
				if !slices.Contains(transactionKeywords, fmt.Sprintf("%f", nb)) {
					return false
				}

				// It was just a matter of formatting, continue because we have found the keyword
				continue
			}
			return false
		}
	}

	return true
}

func transactionKeywords(transaction *model.Transaction, enrichment *model.TransactionEnrichment) []string {
	var transactionKeywords []string
	transactionKeywords = append(transactionKeywords,
		fmt.Sprintf("%f", transaction.Amount),
		transaction.Currency,
		fmt.Sprintf("%d", transaction.AccountID),
		string(transaction.Direction),
	)

	if enrichment.CounterpartyName != nil {
		transactionKeywords = append(transactionKeywords, *enrichment.CounterpartyName)
	}

	if enrichment.Method != nil {
		transactionKeywords = append(transactionKeywords, *enrichment.Method)
	}

	if enrichment.Reference != nil {
		transactionKeywords = append(transactionKeywords, strings.Split(*enrichment.Reference, " ")...)
	} else if transaction.Reference != nil {
		transactionKeywords = append(transactionKeywords, strings.Split(*transaction.Reference, " ")...)
	}

	// Lower everything
	tmp := make([]string, len(transactionKeywords))
	for i, keyword := range transactionKeywords {
		tmp[i] = strings.ToLower(keyword)
	}
	transactionKeywords = tmp
	return transactionKeywords
}

func isNumber(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func getCategoryByName(slice []model.Category, s string) *model.Category {
	for _, e := range slice {
		if strings.EqualFold(e.Name, s) {
			return &e
		}
	}

	return nil
}

func categoryNames(categories []model.Category) []string {
	names := make([]string, len(categories))
	for i, c := range categories {
		names[i] = c.Name
	}

	return names
}
