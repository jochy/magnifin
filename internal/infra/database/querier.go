// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetProviderByID(ctx context.Context, id int32) (Provider, error)
	GetProviderByName(ctx context.Context, name string) (Provider, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByUsernameAndHashedPassword(ctx context.Context, arg GetUserByUsernameAndHashedPasswordParams) (User, error)
	ListProviders(ctx context.Context) ([]Provider, error)
	UpdateProvider(ctx context.Context, arg UpdateProviderParams) (Provider, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpsertConnector(ctx context.Context, arg UpsertConnectorParams) (Connector, error)
}

var _ Querier = (*Queries)(nil)
