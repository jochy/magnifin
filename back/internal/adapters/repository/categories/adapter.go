package categories

import (
	"encoding/json"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

func toCategoryDomain(c *database.Category) *model.Category {
	return &model.Category{
		ID:              c.ID,
		Name:            c.Name,
		UserID:          repository.FromSqlNullInt32(c.UserID),
		Color:           c.Color,
		Icon:            c.Icon,
		IncludeInBudget: c.IncludeInBudget,
	}
}

func toCategoryRuleDomain(cr *database.CategoryRule) *model.CategoryRule {
	var rules []string
	_ = json.Unmarshal([]byte(cr.Rule), &rules)
	return &model.CategoryRule{
		ID:         cr.ID,
		CategoryID: cr.CategoryID,
		Rule:       rules,
	}
}
