package model

import "time"

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
	ID             int32
	ProviderUserID int32
	ConnectorID    int32

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
	SuccessURL string
	ErrorURL   string
}
