package providers

import (
	"context"
	"database/sql"
	"errors"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

type Repository struct {
	db        database.Service
	CypherKey string
}

func NewRepository(db database.Service, cypherKey string) *Repository {
	return &Repository{
		db:        db,
		CypherKey: cypherKey,
	}
}

func (r *Repository) List(ctx context.Context) ([]model.Provider, error) {
	res, err := r.db.ListProviders(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	providers := make([]model.Provider, len(res))
	for i, p := range res {
		provider, err := toDomain(&p, r.CypherKey)
		if err != nil {
			return nil, err
		}
		providers[i] = *provider
	}

	return providers, nil
}

func (r *Repository) Create(ctx context.Context, provider *model.Provider) (*model.Provider, error) {
	accessKey, err := repository.EncryptString(provider.AccessKey, r.CypherKey)
	if err != nil {
		return nil, err
	}

	secret, err := repository.EncryptString(provider.Secret, r.CypherKey)
	if err != nil {
		return nil, err
	}

	p, err := r.db.CreateProvider(ctx, database.CreateProviderParams{
		Name:      provider.Name,
		Enabled:   provider.Enabled,
		AccessKey: toSqlNullValue(accessKey),
		Secret:    toSqlNullValue(secret),
	})
	if err != nil {
		return nil, err
	}

	return toDomain(&p, r.CypherKey)
}

func (r *Repository) Update(ctx context.Context, provider *model.Provider) (*model.Provider, error) {
	accessKey, err := repository.EncryptString(provider.AccessKey, r.CypherKey)
	if err != nil {
		return nil, err
	}

	secret, err := repository.EncryptString(provider.Secret, r.CypherKey)
	if err != nil {
		return nil, err
	}

	p, err := r.db.UpdateProvider(ctx, database.UpdateProviderParams{
		ID:        provider.ID,
		Name:      provider.Name,
		Enabled:   provider.Enabled,
		AccessKey: toSqlNullValue(accessKey),
		Secret:    toSqlNullValue(secret),
	})
	if err != nil {
		return nil, err
	}

	return toDomain(&p, r.CypherKey)
}

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Provider, error) {
	p, err := r.db.GetProviderByName(ctx, name)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(&p, r.CypherKey)
}

func (r *Repository) GetByID(ctx context.Context, id int32) (*model.Provider, error) {
	p, err := r.db.GetProviderByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(&p, r.CypherKey)
}

func toSqlNullValue(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: *s}
}
