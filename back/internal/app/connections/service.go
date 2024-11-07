package connections

import (
	"context"
	"magnifin/internal/app/model"
)

type ConnectionsRepository interface {
	ListActiveByUser(ctx context.Context, user *model.User) ([]model.Connection, error)
	GetByIDAndUser(ctx context.Context, id int32, user *model.User) (*model.Connection, error)
	DeleteByID(ctx context.Context, id int32) error
}

type AccountsRepository interface {
	ListByConnection(ctx context.Context, connectionID int32) ([]model.Account, error)
	DeleteByConnectionID(ctx context.Context, connectionID int32) error
}

type TransactionsRepository interface {
	DeleteByConnectionID(ctx context.Context, connectionID int32) error
}

type ConnectorsRepository interface {
	GetByID(ctx context.Context, id int32) (*model.Connector, error)
}

type ProviderService interface {
	Delete(ctx context.Context, connection *model.Connection) error
}

type Service struct {
	ConnectionsRepository  ConnectionsRepository
	AccountsRepository     AccountsRepository
	ConnectorsRepository   ConnectorsRepository
	TransactionsRepository TransactionsRepository
	ProviderService        ProviderService
}

func NewConnectionsService(
	connections ConnectionsRepository,
	accounts AccountsRepository,
	connectors ConnectorsRepository,
	transactions TransactionsRepository,
	providerService ProviderService,
) *Service {
	return &Service{
		ConnectionsRepository:  connections,
		AccountsRepository:     accounts,
		ConnectorsRepository:   connectors,
		TransactionsRepository: transactions,
		ProviderService:        providerService,
	}
}
