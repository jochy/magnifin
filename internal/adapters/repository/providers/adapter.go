package providers

import (
	"database/sql"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

func toDomain(provider *database.Provider, cypherKey string) (*model.Provider, error) {
	accessKey, err := repository.DecryptString(sqlNullString(provider.AccessKey), cypherKey)
	if err != nil {
		return nil, err
	}

	secret, err := repository.DecryptString(sqlNullString(provider.Secret), cypherKey)
	if err != nil {
		return nil, err
	}

	return &model.Provider{
		ID:        provider.ID,
		Name:      provider.Name,
		Enabled:   provider.Enabled,
		AccessKey: accessKey,
		Secret:    secret,
	}, nil
}

func sqlNullString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}
