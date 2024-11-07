package redirect_sessions

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
	return &Repository{db: db}
}

func (r *Repository) SaveRedirectSession(ctx context.Context, session model.RedirectSession) error {
	return r.db.StoreRedirectSessions(ctx, database.StoreRedirectSessionsParams{
		ID:                   session.ID,
		ProviderConnectionID: sqlNullString(session.ProviderConnectionID),
		InternalConnectionID: sqlNullInt32(session.InternalConnectionID),
		UserID: sql.NullInt32{
			Int32: session.UserID,
			Valid: true,
		},
	})
}

func (r *Repository) GetRedirectSessionByID(ctx context.Context, id string) (*model.RedirectSession, error) {
	session, err := r.db.GetRedirectSessionByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &model.RedirectSession{
		ID:                   session.ID,
		ProviderConnectionID: fromSqlNullString(session.ProviderConnectionID),
		InternalConnectionID: fromSqlNullInt32(session.InternalConnectionID),
		UserID:               session.UserID.Int32,
	}, nil
}

func sqlNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func sqlNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *i,
		Valid: true,
	}
}

func fromSqlNullString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func fromSqlNullInt32(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}
	return &i.Int32
}
