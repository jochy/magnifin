package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrRateLimited error = errors.New("rate limited")

type ConnectionStatus string

const (
	ConnectionStatusSynchronized   ConnectionStatus = "SYNCHRONIZED"
	ConnectionStatusSyncInProgress ConnectionStatus = "SYNC_IN_PROGRESS"
	ConnectionStatusSuspended      ConnectionStatus = "SUSPENDED"
	ConnectionStatusDeleted        ConnectionStatus = "DELETED"
	ConnectionStatusRateLimited    ConnectionStatus = "RATE_LIMITED"
)

type TransactionDirection string

const (
	TransactionDirectionCredit TransactionDirection = "CREDIT"
	TransactionDirectionDebit  TransactionDirection = "DEBIT"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "SETTLED"
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
	UserID               int32
}

type Account struct {
	ID                int32
	ConnectionID      int32
	ProviderAccountID string
	BankAccountID     *string
	Name              *string
	Type              *string
	Currency          *string
	AccountNumber     *string
	Balance           float64
}

type Transaction struct {
	ID                    int32
	AccountID             int32
	ProviderTransactionID string
	BankTransactionID     *string
	Amount                float64
	Currency              string
	Direction             TransactionDirection
	Status                TransactionStatus
	OperationAt           time.Time

	CounterpartyName    *string
	CounterpartyAccount *string
	Reference           *string

	Enrichment *TransactionEnrichment
}

type TransactionEnrichment struct {
	ID            int32
	TransactionID int32

	CounterpartyLogo     *string
	Category             *int32
	CounterpartyName     *string
	UserCounterpartyName *string
	Reference            *string
	Method               *string
}

type Image struct {
	ID          string
	Content     string
	ContentType string
}

type ConnectionWithAccounts struct {
	Connection *Connection
	Connector  *Connector
	Accounts   []Account
}

type TransactionMinAndMax struct {
	Min time.Time
	Max time.Time
}

type Category struct {
	ID              int32
	Name            string
	UserID          *int32
	Color           string
	Icon            string
	IncludeInBudget bool
}

type CategoryRule struct {
	ID         int32
	CategoryID int32
	Rule       []string
}
