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

func toSqlNullValue(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
