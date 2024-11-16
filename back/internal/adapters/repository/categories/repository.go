package categories

import (
	"context"
	"database/sql"
	"encoding/json"
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
	return &Repository{db: db}
}

func (r *Repository) GetAllRulesByUserFromTransID(ctx context.Context, transID int32) ([]model.CategoryRule, error) {
	rules, err := r.db.GetAllRulesByUserFromTransactionID(ctx, transID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := make([]model.CategoryRule, len(rules))
	for i, c := range rules {
		result[i] = *toCategoryRuleDomain(&c)
	}

	return result, nil
}

func (r *Repository) GetAllCategoriesByUserFromTransactionID(ctx context.Context, transID int32) ([]model.Category, error) {
	categories, err := r.db.GetAllCategoriesByUserFromTransactionID(ctx, transID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := make([]model.Category, len(categories))
	for i, c := range categories {
		result[i] = *toCategoryDomain(&c)
	}

	return result, nil
}

func (r *Repository) GetAllCategoriesByUserID(ctx context.Context, userID int32) ([]model.Category, error) {
	categories, err := r.db.GetAllCategoriesByUserID(ctx, repository.ToSqlNullInt32(&userID))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := make([]model.Category, len(categories))
	for i, c := range categories {
		result[i] = *toCategoryDomain(&c)
	}

	return result, nil
}

func (r *Repository) DeleteCategoryByID(ctx context.Context, id int32) error {
	err := r.db.DeleteCategoryByID(ctx, id)
	if err != nil {
		return err
	}

	return r.db.DeleteCategoryRuleByCategoryID(ctx, id)
}

func (r *Repository) CreateCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	c, err := r.db.CreateCategory(ctx, database.CreateCategoryParams{
		Name:            category.Name,
		Icon:            category.Icon,
		Color:           category.Color,
		IncludeInBudget: category.IncludeInBudget,
		UserID:          repository.ToSqlNullInt32(category.UserID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("category not found")
	} else if err != nil {
		return nil, err
	}

	return toCategoryDomain(&c), nil
}

func (r *Repository) UpdateCategory(ctx context.Context, category model.Category) error {
	_, err := r.db.UpdateCategoryByIDAndUserID(ctx, database.UpdateCategoryByIDAndUserIDParams{
		ID:              category.ID,
		Name:            category.Name,
		Icon:            category.Icon,
		Color:           category.Color,
		IncludeInBudget: category.IncludeInBudget,
		UserID:          repository.ToSqlNullInt32(category.UserID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("category not found")
	} else if err != nil {
		return err
	}

	return r.db.DeleteCategoryRuleByCategoryID(ctx, category.ID)
}

func (r *Repository) CreateCategoryRule(ctx context.Context, categoryRule model.CategoryRule) (*model.CategoryRule, error) {
	jsonData, err := json.Marshal(categoryRule.Rule)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the rule: %w", err)
	}

	rule, err := r.db.CreateCategoryRule(ctx, database.CreateCategoryRuleParams{
		CategoryID: categoryRule.CategoryID,
		Rule:       string(jsonData),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("category rule not found")
	} else if err != nil {
		return nil, err
	}

	return toCategoryRuleDomain(&rule), nil
}

func (r *Repository) GetCategoryRuleByID(ctx context.Context, id int32) (*model.CategoryRule, error) {
	rule, err := r.db.GetCategoryRuleByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("category rule not found")
	} else if err != nil {
		return nil, err
	}

	return toCategoryRuleDomain(&rule), nil
}
