package providerusers

import (
	"context"
	"database/sql"
	"errors"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

type Repository struct {
	db database.Service
}

func NewRepository(db database.Service) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByProviderIDAndUserID(ctx context.Context, providerID int32, userID int32) (*model.ProviderUser, error) {
	p, err := r.db.GetProviderUserByProviderIDAndUserID(ctx, database.GetProviderUserByProviderIDAndUserIDParams{
		ProviderID: providerID,
		UserID:     userID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(&p), nil
}

func (r *Repository) Save(ctx context.Context, providerID int32, userID int32, providerUserID string) (*model.ProviderUser, error) {
	p, err := r.db.CreateProviderUser(ctx, database.CreateProviderUserParams{
		ProviderID:     providerID,
		UserID:         userID,
		ProviderUserID: providerUserID,
	})
	if err != nil {
		return nil, err
	}

	return toDomain(&p), nil
}
