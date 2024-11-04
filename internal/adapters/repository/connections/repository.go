package connections

import (
	"context"
	"database/sql"
	"errors"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"time"
)

type Repository struct {
	db database.Service
}

func NewRepository(db database.Service) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByProviderUserIDAndProviderConnectionID(ctx context.Context, providerUserID int32, providerConnectionID string) (*model.Connection, error) {
	connection, err := r.db.GetConnectionByProviderUserIDAndProviderConnectionID(ctx, database.GetConnectionByProviderUserIDAndProviderConnectionIDParams{
		ProviderUsersID:      providerUserID,
		ProviderConnectionID: providerConnectionID,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomainModel(&connection), nil
}

func (r *Repository) Create(ctx context.Context, connection *model.Connection) (*model.Connection, error) {
	saved, err := r.db.CreateConnection(ctx, database.CreateConnectionParams{
		ProviderUsersID:      connection.ProviderUserID,
		ProviderConnectionID: connection.ProviderConnectionID,
		ConnectorID:          connection.ConnectorID,
		Status:               string(connection.Status),
		RenewConsentBefore:   toSqlNullTime(connection.RenewConsentBefore),
		ErrorMessage:         toSqlNullString(connection.ErrorMessage),
		LastSuccessfulSync:   toSqlNullTime(connection.LastSuccessfulSync),
	})
	if err != nil {
		return nil, err
	}

	return toDomainModel(&saved), nil
}

func (r *Repository) Update(ctx context.Context, connection *model.Connection) (*model.Connection, error) {
	saved, err := r.db.UpdateConnection(ctx, database.UpdateConnectionParams{
		ID:                   connection.ID,
		ProviderConnectionID: connection.ProviderConnectionID,
		Status:               string(connection.Status),
		RenewConsentBefore:   toSqlNullTime(connection.RenewConsentBefore),
		ErrorMessage:         toSqlNullString(connection.ErrorMessage),
		LastSuccessfulSync:   toSqlNullTime(connection.LastSuccessfulSync),
	})
	if err != nil {
		return nil, err
	}

	return toDomainModel(&saved), nil
}

func (r *Repository) GetByID(ctx context.Context, id int32) (*model.Connection, error) {
	connection, err := r.db.GetConnectionByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomainModel(&connection), nil
}

func toSqlNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func toSqlNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{
			Time:  *t,
			Valid: true,
		}
	}
	return sql.NullTime{}
}
