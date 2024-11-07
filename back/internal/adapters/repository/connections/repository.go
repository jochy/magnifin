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

func (r *Repository) UpdateStatus(ctx context.Context, id int32, status model.ConnectionStatus) error {
	_, err := r.db.UpdateConnectionStatus(ctx, database.UpdateConnectionStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListActiveByUser(ctx context.Context, user *model.User) ([]model.Connection, error) {
	c, err := r.db.ListConnectionsByUserID(ctx, user.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	connections := make([]model.Connection, len(c))
	for i, connection := range c {
		connections[i] = *toDomainModel(&connection)
	}

	return connections, nil
}

func (r *Repository) ListConnectionsToSync(ctx context.Context) ([]model.Connection, error) {
	c, err := r.db.ListConnectionsToSync(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	connections := make([]model.Connection, len(c))
	for i, connection := range c {
		connections[i] = *toDomainModel(&connection)
	}

	return connections, nil
}

func (r *Repository) GetByIDAndUser(ctx context.Context, id int32, user *model.User) (*model.Connection, error) {
	connection, err := r.db.GetConnectionByIDAndUserID(ctx, database.GetConnectionByIDAndUserIDParams{
		ID:     id,
		UserID: user.ID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomainModel(&connection), nil
}

func (r *Repository) DeleteByID(ctx context.Context, id int32) error {
	err := r.db.DeleteConnectionByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
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
