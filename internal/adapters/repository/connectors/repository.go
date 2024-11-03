package connectors

import (
	"context"
	"database/sql"
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

func (r *Repository) Upsert(ctx context.Context, connector *model.Connector) (*model.Connector, error) {
	c, err := r.db.UpsertConnector(ctx, database.UpsertConnectorParams{
		ProviderID:          connector.ProviderID,
		ProviderConnectorID: connector.ProviderConnectorID,
		Name:                connector.Name,
		LogoUrl:             toSqlNullValue(connector.LogoURL),
	})
	if err != nil {
		return nil, err
	}

	return toDomain(c), nil
}

func (r *Repository) SearchByName(ctx context.Context, name string) ([]model.Connector, error) {
	connectors, err := r.db.FuzzySearchConnectorsByName(ctx, name)
	if err != nil {
		return nil, err
	}

	result := make([]model.Connector, len(connectors))
	for i, c := range connectors {
		result[i] = *toDomain(c)
	}

	return result, nil
}

func (r *Repository) LikeSearchByName(ctx context.Context, name string) ([]model.Connector, error) {
	connectors, err := r.db.LikeSearchConnectorsByName(ctx, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	result := make([]model.Connector, len(connectors))
	for i, c := range connectors {
		result[i] = *toDomain(c)
	}

	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, id int32) (*model.Connector, error) {
	c, err := r.db.GetConnectorByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toDomain(c), nil
}

func toSqlNullValue(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
