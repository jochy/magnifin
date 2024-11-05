package connections

import (
	"database/sql"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"time"
)

func toDomainModel(connection *database.Connection) *model.Connection {
	return &model.Connection{
		ID:                   connection.ID,
		ProviderUserID:       connection.ProviderUsersID,
		ConnectorID:          connection.ConnectorID,
		ProviderConnectionID: connection.ProviderConnectionID,
		Status:               model.ConnectionStatus(connection.Status),
		RenewConsentBefore:   sqlNullTime(connection.RenewConsentBefore),
		ErrorMessage:         sqlNullString(connection.ErrorMessage),
		LastSuccessfulSync:   sqlNullTime(connection.LastSuccessfulSync),
	}
}

func sqlNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func sqlNullTime(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}
