package categories

import (
	"context"
	"magnifin/internal/app/model"
)

type Repository interface {
	GetAllCategoriesByUserID(ctx context.Context, userID int32) ([]model.Category, error)
	CreateCategoryRule(ctx context.Context, categoryRule model.CategoryRule) (*model.CategoryRule, error)
}

type Handlers struct {
	Repository Repository
}

func NewHandlers(repository Repository) *Handlers {
	return &Handlers{Repository: repository}
}
