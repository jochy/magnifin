package model

import (
	"time"

	"github.com/google/uuid"
)

type ConnectionStatus string

const (
	ConnectionStatusSynchronized   ConnectionStatus = "SYNCHRONIZED"
	ConnectionStatusSyncInProgress ConnectionStatus = "SYNC_IN_PROGRESS"
	ConnectionStatusSuspended      ConnectionStatus = "SUSPENDED"
	ConnectionStatusDeleted        ConnectionStatus = "DELETED"
)

type User struct {
	ID       int32
	Username string
}

type Provider struct {
	ID        int32
	Name      string
	AccessKey *string
	Secret    *string
	Enabled   bool
}

type ProviderUser struct {
	ID             int32
	ProviderID     int32
	UserID         int32
	ProviderUserID string
}

type Connector struct {
	ID                  int32
	Name                string
	LogoURL             string
	ProviderID          int32
	ProviderConnectorID string
}

type Connection struct {
	ID                   int32
	ProviderUserID       int32
	ConnectorID          int32
	ProviderConnectionID string

	Status             ConnectionStatus
	RenewConsentBefore *time.Time
	ErrorMessage       *string
	LastSuccessfulSync *time.Time
}

type ConnectInstruction struct {
	ID          string
	RedirectURL string
}

type ConnectParams struct {
	SID        uuid.UUID
	SuccessURL string
	ErrorURL   string
}

type RedirectSession struct {
	ID                   string
	ProviderConnectionID *string
	InternalConnectionID *int32
}

type Account struct {
	ID                int32
	ConnectionID      int32
	ProviderAccountID string
	Name              *string
	Type              *string
	Currency          *string
	AccountNumber     *string
	Balance           float64
}
