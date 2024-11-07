// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
insert into accounts (connection_id, provider_account_id, name, type, currency, account_number, balance,
                      bank_account_id)
values ($1, $2, $3, $4, $5, $6, $7, $8) returning id, connection_id, provider_account_id, bank_account_id, name, type, currency, account_number, balance, created_at, updated_at, deleted_at
`

type CreateAccountParams struct {
	ConnectionID      int32          `db:"connection_id"`
	ProviderAccountID string         `db:"provider_account_id"`
	Name              sql.NullString `db:"name"`
	Type              sql.NullString `db:"type"`
	Currency          sql.NullString `db:"currency"`
	AccountNumber     sql.NullString `db:"account_number"`
	Balance           string         `db:"balance"`
	BankAccountID     sql.NullString `db:"bank_account_id"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.ConnectionID,
		arg.ProviderAccountID,
		arg.Name,
		arg.Type,
		arg.Currency,
		arg.AccountNumber,
		arg.Balance,
		arg.BankAccountID,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.ConnectionID,
		&i.ProviderAccountID,
		&i.BankAccountID,
		&i.Name,
		&i.Type,
		&i.Currency,
		&i.AccountNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createConnection = `-- name: CreateConnection :one
insert into connections (provider_users_id, provider_connection_id, connector_id, status, renew_consent_before,
                         error_message, last_successful_sync)
values ($1, $2, $3, $4, $5, $6, $7) returning id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
`

type CreateConnectionParams struct {
	ProviderUsersID      int32          `db:"provider_users_id"`
	ProviderConnectionID string         `db:"provider_connection_id"`
	ConnectorID          int32          `db:"connector_id"`
	Status               string         `db:"status"`
	RenewConsentBefore   sql.NullTime   `db:"renew_consent_before"`
	ErrorMessage         sql.NullString `db:"error_message"`
	LastSuccessfulSync   sql.NullTime   `db:"last_successful_sync"`
}

func (q *Queries) CreateConnection(ctx context.Context, arg CreateConnectionParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, createConnection,
		arg.ProviderUsersID,
		arg.ProviderConnectionID,
		arg.ConnectorID,
		arg.Status,
		arg.RenewConsentBefore,
		arg.ErrorMessage,
		arg.LastSuccessfulSync,
	)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createProviderUser = `-- name: CreateProviderUser :one
insert into provider_users (provider_id, user_id, provider_user_id)
values ($1, $2, $3) returning id, provider_id, provider_user_id, user_id, created_at, updated_at, deleted_at
`

type CreateProviderUserParams struct {
	ProviderID     int32  `db:"provider_id"`
	UserID         int32  `db:"user_id"`
	ProviderUserID string `db:"provider_user_id"`
}

func (q *Queries) CreateProviderUser(ctx context.Context, arg CreateProviderUserParams) (ProviderUser, error) {
	row := q.db.QueryRowContext(ctx, createProviderUser, arg.ProviderID, arg.UserID, arg.ProviderUserID)
	var i ProviderUser
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.ProviderUserID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createTransaction = `-- name: CreateTransaction :one
insert into transactions (account_id, provider_transaction_id, bank_transaction_id, amount, currency, direction, status,
                          operation_at, counterparty_name, counterparty_account,
                          reference)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id, account_id, provider_transaction_id, bank_transaction_id, amount, currency, direction, status, operation_at, counterparty_name, counterparty_account, reference, created_at, updated_at, deleted_at
`

type CreateTransactionParams struct {
	AccountID             int32          `db:"account_id"`
	ProviderTransactionID string         `db:"provider_transaction_id"`
	BankTransactionID     sql.NullString `db:"bank_transaction_id"`
	Amount                string         `db:"amount"`
	Currency              string         `db:"currency"`
	Direction             string         `db:"direction"`
	Status                string         `db:"status"`
	OperationAt           time.Time      `db:"operation_at"`
	CounterpartyName      sql.NullString `db:"counterparty_name"`
	CounterpartyAccount   sql.NullString `db:"counterparty_account"`
	Reference             sql.NullString `db:"reference"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.AccountID,
		arg.ProviderTransactionID,
		arg.BankTransactionID,
		arg.Amount,
		arg.Currency,
		arg.Direction,
		arg.Status,
		arg.OperationAt,
		arg.CounterpartyName,
		arg.CounterpartyAccount,
		arg.Reference,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ProviderTransactionID,
		&i.BankTransactionID,
		&i.Amount,
		&i.Currency,
		&i.Direction,
		&i.Status,
		&i.OperationAt,
		&i.CounterpartyName,
		&i.CounterpartyAccount,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
insert into users (username, hashed_password)
values ($1, $2) returning id, username, hashed_password, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Username       string `db:"username"`
	HashedPassword string `db:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAccountByConnectionID = `-- name: DeleteAccountByConnectionID :exec
update accounts
set deleted_at = now()
where connection_id = $1
`

func (q *Queries) DeleteAccountByConnectionID(ctx context.Context, connectionID int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccountByConnectionID, connectionID)
	return err
}

const deleteConnectionByID = `-- name: DeleteConnectionByID :exec
update connections
set deleted_at = now()
where id = $1
`

func (q *Queries) DeleteConnectionByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteConnectionByID, id)
	return err
}

const deleteTransactionsByConnectionID = `-- name: DeleteTransactionsByConnectionID :exec
update transactions
set deleted_at = now()
where account_id in (select id from accounts where connection_id = $1)
`

func (q *Queries) DeleteTransactionsByConnectionID(ctx context.Context, connectionID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTransactionsByConnectionID, connectionID)
	return err
}

const deleteTransactionsEnrichmentsByConnectionID = `-- name: DeleteTransactionsEnrichmentsByConnectionID :exec
update transaction_enrichments
set deleted_at = now()
where transaction_id in
      (select id from transactions where account_id in (select id from accounts where connection_id = $1))
`

func (q *Queries) DeleteTransactionsEnrichmentsByConnectionID(ctx context.Context, connectionID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTransactionsEnrichmentsByConnectionID, connectionID)
	return err
}

const findTransactionByAccountIDAndProviderTransactionID = `-- name: FindTransactionByAccountIDAndProviderTransactionID :one
select id, account_id, provider_transaction_id, bank_transaction_id, amount, currency, direction, status, operation_at, counterparty_name, counterparty_account, reference, created_at, updated_at, deleted_at
from transactions
where account_id = $1
  and provider_transaction_id = $2
  and deleted_at is null
`

type FindTransactionByAccountIDAndProviderTransactionIDParams struct {
	AccountID             int32  `db:"account_id"`
	ProviderTransactionID string `db:"provider_transaction_id"`
}

func (q *Queries) FindTransactionByAccountIDAndProviderTransactionID(ctx context.Context, arg FindTransactionByAccountIDAndProviderTransactionIDParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, findTransactionByAccountIDAndProviderTransactionID, arg.AccountID, arg.ProviderTransactionID)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ProviderTransactionID,
		&i.BankTransactionID,
		&i.Amount,
		&i.Currency,
		&i.Direction,
		&i.Status,
		&i.OperationAt,
		&i.CounterpartyName,
		&i.CounterpartyAccount,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const fuzzySearchConnectorsByName = `-- name: FuzzySearchConnectorsByName :many
select connectors.id, connectors.name, connectors.logo_url, connectors.provider_connector_id, connectors.provider_id, connectors.created_at, connectors.updated_at, connectors.deleted_at
from connectors
         inner join providers on connectors.provider_id = providers.id
where connectors.name % $1
  and connectors.deleted_at is null
  and providers.deleted_at is null
  and providers.enabled = true
`

func (q *Queries) FuzzySearchConnectorsByName(ctx context.Context, name string) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, fuzzySearchConnectorsByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Connector{}
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LogoUrl,
			&i.ProviderConnectorID,
			&i.ProviderID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountByConnectionIDAndProviderAccountID = `-- name: GetAccountByConnectionIDAndProviderAccountID :one
select id, connection_id, provider_account_id, bank_account_id, name, type, currency, account_number, balance, created_at, updated_at, deleted_at
from accounts
where connection_id = $1
  and provider_account_id = $2
  and deleted_at is null
`

type GetAccountByConnectionIDAndProviderAccountIDParams struct {
	ConnectionID      int32  `db:"connection_id"`
	ProviderAccountID string `db:"provider_account_id"`
}

func (q *Queries) GetAccountByConnectionIDAndProviderAccountID(ctx context.Context, arg GetAccountByConnectionIDAndProviderAccountIDParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByConnectionIDAndProviderAccountID, arg.ConnectionID, arg.ProviderAccountID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.ConnectionID,
		&i.ProviderAccountID,
		&i.BankAccountID,
		&i.Name,
		&i.Type,
		&i.Currency,
		&i.AccountNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getConnectionByID = `-- name: GetConnectionByID :one
select id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
from connections
where id = $1
  and deleted_at is null
`

func (q *Queries) GetConnectionByID(ctx context.Context, id int32) (Connection, error) {
	row := q.db.QueryRowContext(ctx, getConnectionByID, id)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getConnectionByIDAndUserID = `-- name: GetConnectionByIDAndUserID :one
select id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
from connections
where connections.id = $1
  and provider_users_id in (select provider_users.id from provider_users where user_id = $2 and deleted_at is null)
  and deleted_at is null
`

type GetConnectionByIDAndUserIDParams struct {
	ID     int32 `db:"id"`
	UserID int32 `db:"user_id"`
}

func (q *Queries) GetConnectionByIDAndUserID(ctx context.Context, arg GetConnectionByIDAndUserIDParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, getConnectionByIDAndUserID, arg.ID, arg.UserID)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getConnectionByProviderUserIDAndProviderConnectionID = `-- name: GetConnectionByProviderUserIDAndProviderConnectionID :one
select id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
from connections
where provider_users_id = $1
  and provider_connection_id = $2
  and deleted_at is null
`

type GetConnectionByProviderUserIDAndProviderConnectionIDParams struct {
	ProviderUsersID      int32  `db:"provider_users_id"`
	ProviderConnectionID string `db:"provider_connection_id"`
}

func (q *Queries) GetConnectionByProviderUserIDAndProviderConnectionID(ctx context.Context, arg GetConnectionByProviderUserIDAndProviderConnectionIDParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, getConnectionByProviderUserIDAndProviderConnectionID, arg.ProviderUsersID, arg.ProviderConnectionID)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getConnectorByID = `-- name: GetConnectorByID :one
select id, name, logo_url, provider_connector_id, provider_id, created_at, updated_at, deleted_at
from connectors
where id = $1
  and deleted_at is null
`

func (q *Queries) GetConnectorByID(ctx context.Context, id int32) (Connector, error) {
	row := q.db.QueryRowContext(ctx, getConnectorByID, id)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LogoUrl,
		&i.ProviderConnectorID,
		&i.ProviderID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProviderByID = `-- name: GetProviderByID :one
select id, name, access_key, secret, enabled, created_at, updated_at, deleted_at
from providers
where id = $1
  and deleted_at is null
`

func (q *Queries) GetProviderByID(ctx context.Context, id int32) (Provider, error) {
	row := q.db.QueryRowContext(ctx, getProviderByID, id)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccessKey,
		&i.Secret,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProviderByName = `-- name: GetProviderByName :one
select id, name, access_key, secret, enabled, created_at, updated_at, deleted_at
from providers
where name = $1
  and deleted_at is null
`

func (q *Queries) GetProviderByName(ctx context.Context, name string) (Provider, error) {
	row := q.db.QueryRowContext(ctx, getProviderByName, name)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccessKey,
		&i.Secret,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProviderUserByID = `-- name: GetProviderUserByID :one
select id, provider_id, provider_user_id, user_id, created_at, updated_at, deleted_at
from provider_users
where id = $1
  and deleted_at is null
`

func (q *Queries) GetProviderUserByID(ctx context.Context, id int32) (ProviderUser, error) {
	row := q.db.QueryRowContext(ctx, getProviderUserByID, id)
	var i ProviderUser
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.ProviderUserID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProviderUserByProviderIDAndUserID = `-- name: GetProviderUserByProviderIDAndUserID :one
select id, provider_id, provider_user_id, user_id, created_at, updated_at, deleted_at
from provider_users
where provider_id = $1
  and user_id = $2
  and deleted_at is null
`

type GetProviderUserByProviderIDAndUserIDParams struct {
	ProviderID int32 `db:"provider_id"`
	UserID     int32 `db:"user_id"`
}

func (q *Queries) GetProviderUserByProviderIDAndUserID(ctx context.Context, arg GetProviderUserByProviderIDAndUserIDParams) (ProviderUser, error) {
	row := q.db.QueryRowContext(ctx, getProviderUserByProviderIDAndUserID, arg.ProviderID, arg.UserID)
	var i ProviderUser
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.ProviderUserID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRedirectSessionByID = `-- name: GetRedirectSessionByID :one
select id, provider_connection_id, internal_connection_id, user_id, created_at
from redirect_sessions
where id = $1
`

func (q *Queries) GetRedirectSessionByID(ctx context.Context, id string) (RedirectSession, error) {
	row := q.db.QueryRowContext(ctx, getRedirectSessionByID, id)
	var i RedirectSession
	err := row.Scan(
		&i.ID,
		&i.ProviderConnectionID,
		&i.InternalConnectionID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id, username, hashed_password, created_at, updated_at, deleted_at
from users
where id = $1
  and deleted_at is null
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select id, username, hashed_password, created_at, updated_at, deleted_at
from users
where username = $1
  and deleted_at is null
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const likeSearchConnectorsByName = `-- name: LikeSearchConnectorsByName :many
select connectors.id, connectors.name, connectors.logo_url, connectors.provider_connector_id, connectors.provider_id, connectors.created_at, connectors.updated_at, connectors.deleted_at
from connectors
         inner join providers on connectors.provider_id = providers.id
where connectors.name ilike $1
  and connectors.deleted_at is null
  and providers.deleted_at is null
  and providers.enabled = true
`

func (q *Queries) LikeSearchConnectorsByName(ctx context.Context, name string) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, likeSearchConnectorsByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Connector{}
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LogoUrl,
			&i.ProviderConnectorID,
			&i.ProviderID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAccountsByConnectionID = `-- name: ListAccountsByConnectionID :many
select id, connection_id, provider_account_id, bank_account_id, name, type, currency, account_number, balance, created_at, updated_at, deleted_at
from accounts
where connection_id = $1
  and deleted_at is null
`

func (q *Queries) ListAccountsByConnectionID(ctx context.Context, connectionID int32) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccountsByConnectionID, connectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.ConnectionID,
			&i.ProviderAccountID,
			&i.BankAccountID,
			&i.Name,
			&i.Type,
			&i.Currency,
			&i.AccountNumber,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listConnectionsByUserID = `-- name: ListConnectionsByUserID :many
select id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
from connections
where provider_users_id in (select id from provider_users where user_id = $1 and deleted_at is null)
  and deleted_at is null
`

func (q *Queries) ListConnectionsByUserID(ctx context.Context, userID int32) ([]Connection, error) {
	rows, err := q.db.QueryContext(ctx, listConnectionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Connection{}
	for rows.Next() {
		var i Connection
		if err := rows.Scan(
			&i.ID,
			&i.ProviderUsersID,
			&i.ProviderConnectionID,
			&i.ConnectorID,
			&i.Status,
			&i.RenewConsentBefore,
			&i.ErrorMessage,
			&i.LastSuccessfulSync,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listConnectionsToSync = `-- name: ListConnectionsToSync :many
select id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
from connections
where ((last_successful_sync is null and created_at < now() - interval '1 hour') or
       (last_successful_sync < now() - interval '11 hours'))
  and deleted_at is null
`

func (q *Queries) ListConnectionsToSync(ctx context.Context) ([]Connection, error) {
	rows, err := q.db.QueryContext(ctx, listConnectionsToSync)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Connection{}
	for rows.Next() {
		var i Connection
		if err := rows.Scan(
			&i.ID,
			&i.ProviderUsersID,
			&i.ProviderConnectionID,
			&i.ConnectorID,
			&i.Status,
			&i.RenewConsentBefore,
			&i.ErrorMessage,
			&i.LastSuccessfulSync,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProviders = `-- name: ListProviders :many
select id, name, access_key, secret, enabled, created_at, updated_at, deleted_at
from providers
where deleted_at is null
`

func (q *Queries) ListProviders(ctx context.Context) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, listProviders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AccessKey,
			&i.Secret,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const storeRedirectSessions = `-- name: StoreRedirectSessions :exec
insert into redirect_sessions (id, provider_connection_id, internal_connection_id, user_id)
values ($1, $2, $3, $4)
`

type StoreRedirectSessionsParams struct {
	ID                   string         `db:"id"`
	ProviderConnectionID sql.NullString `db:"provider_connection_id"`
	InternalConnectionID sql.NullInt32  `db:"internal_connection_id"`
	UserID               sql.NullInt32  `db:"user_id"`
}

func (q *Queries) StoreRedirectSessions(ctx context.Context, arg StoreRedirectSessionsParams) error {
	_, err := q.db.ExecContext(ctx, storeRedirectSessions,
		arg.ID,
		arg.ProviderConnectionID,
		arg.InternalConnectionID,
		arg.UserID,
	)
	return err
}

const updateAccount = `-- name: UpdateAccount :one
update accounts
set name                = $2,
    type                = $3,
    currency            = $4,
    account_number      = $5,
    balance             = $6,
    provider_account_id = $7,
    bank_account_id     = $8,
    updated_at          = now()
where id = $1 returning id, connection_id, provider_account_id, bank_account_id, name, type, currency, account_number, balance, created_at, updated_at, deleted_at
`

type UpdateAccountParams struct {
	ID                int32          `db:"id"`
	Name              sql.NullString `db:"name"`
	Type              sql.NullString `db:"type"`
	Currency          sql.NullString `db:"currency"`
	AccountNumber     sql.NullString `db:"account_number"`
	Balance           string         `db:"balance"`
	ProviderAccountID string         `db:"provider_account_id"`
	BankAccountID     sql.NullString `db:"bank_account_id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.Currency,
		arg.AccountNumber,
		arg.Balance,
		arg.ProviderAccountID,
		arg.BankAccountID,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.ConnectionID,
		&i.ProviderAccountID,
		&i.BankAccountID,
		&i.Name,
		&i.Type,
		&i.Currency,
		&i.AccountNumber,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateConnection = `-- name: UpdateConnection :one
update connections
set status                 = $2,
    renew_consent_before   = $3,
    error_message          = $4,
    last_successful_sync   = $5,
    provider_connection_id = $6,
    updated_at             = now()
where id = $1 returning id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
`

type UpdateConnectionParams struct {
	ID                   int32          `db:"id"`
	Status               string         `db:"status"`
	RenewConsentBefore   sql.NullTime   `db:"renew_consent_before"`
	ErrorMessage         sql.NullString `db:"error_message"`
	LastSuccessfulSync   sql.NullTime   `db:"last_successful_sync"`
	ProviderConnectionID string         `db:"provider_connection_id"`
}

func (q *Queries) UpdateConnection(ctx context.Context, arg UpdateConnectionParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, updateConnection,
		arg.ID,
		arg.Status,
		arg.RenewConsentBefore,
		arg.ErrorMessage,
		arg.LastSuccessfulSync,
		arg.ProviderConnectionID,
	)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateConnectionStatus = `-- name: UpdateConnectionStatus :one
update connections
set status = $2
where id = $1 returning id, provider_users_id, provider_connection_id, connector_id, status, renew_consent_before, error_message, last_successful_sync, created_at, updated_at, deleted_at
`

type UpdateConnectionStatusParams struct {
	ID     int32  `db:"id"`
	Status string `db:"status"`
}

func (q *Queries) UpdateConnectionStatus(ctx context.Context, arg UpdateConnectionStatusParams) (Connection, error) {
	row := q.db.QueryRowContext(ctx, updateConnectionStatus, arg.ID, arg.Status)
	var i Connection
	err := row.Scan(
		&i.ID,
		&i.ProviderUsersID,
		&i.ProviderConnectionID,
		&i.ConnectorID,
		&i.Status,
		&i.RenewConsentBefore,
		&i.ErrorMessage,
		&i.LastSuccessfulSync,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateProvider = `-- name: UpdateProvider :one
update providers
set name       = $2,
    access_key = $3,
    secret     = $4,
    enabled    = $5,
    updated_at = now()
where id = $1 returning id, name, access_key, secret, enabled, created_at, updated_at, deleted_at
`

type UpdateProviderParams struct {
	ID        int32          `db:"id"`
	Name      string         `db:"name"`
	AccessKey sql.NullString `db:"access_key"`
	Secret    sql.NullString `db:"secret"`
	Enabled   bool           `db:"enabled"`
}

func (q *Queries) UpdateProvider(ctx context.Context, arg UpdateProviderParams) (Provider, error) {
	row := q.db.QueryRowContext(ctx, updateProvider,
		arg.ID,
		arg.Name,
		arg.AccessKey,
		arg.Secret,
		arg.Enabled,
	)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AccessKey,
		&i.Secret,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :one
update transactions
set bank_transaction_id     = $2,
    amount                  = $3,
    currency                = $4,
    direction               = $5,
    status                  = $6,
    operation_at            = $7,
    counterparty_name       = $8,
    counterparty_account    = $9,
    reference               = $10,
    provider_transaction_id = $11,
    updated_at              = now()
where id = $1 returning id, account_id, provider_transaction_id, bank_transaction_id, amount, currency, direction, status, operation_at, counterparty_name, counterparty_account, reference, created_at, updated_at, deleted_at
`

type UpdateTransactionParams struct {
	ID                    int32          `db:"id"`
	BankTransactionID     sql.NullString `db:"bank_transaction_id"`
	Amount                string         `db:"amount"`
	Currency              string         `db:"currency"`
	Direction             string         `db:"direction"`
	Status                string         `db:"status"`
	OperationAt           time.Time      `db:"operation_at"`
	CounterpartyName      sql.NullString `db:"counterparty_name"`
	CounterpartyAccount   sql.NullString `db:"counterparty_account"`
	Reference             sql.NullString `db:"reference"`
	ProviderTransactionID string         `db:"provider_transaction_id"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, updateTransaction,
		arg.ID,
		arg.BankTransactionID,
		arg.Amount,
		arg.Currency,
		arg.Direction,
		arg.Status,
		arg.OperationAt,
		arg.CounterpartyName,
		arg.CounterpartyAccount,
		arg.Reference,
		arg.ProviderTransactionID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ProviderTransactionID,
		&i.BankTransactionID,
		&i.Amount,
		&i.Currency,
		&i.Direction,
		&i.Status,
		&i.OperationAt,
		&i.CounterpartyName,
		&i.CounterpartyAccount,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
update users
set hashed_password = $2,
    updated_at      = now()
where id = $1 returning id, username, hashed_password, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	ID             int32  `db:"id"`
	HashedPassword string `db:"hashed_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const upsertConnector = `-- name: UpsertConnector :one
insert into connectors (name, logo_url, provider_connector_id, provider_id)
values ($1, $2, $3, $4) on conflict (provider_id, provider_connector_id) do
update
    set name = excluded.name,
    logo_url = excluded.logo_url
    returning id, name, logo_url, provider_connector_id, provider_id, created_at, updated_at, deleted_at
`

type UpsertConnectorParams struct {
	Name                string         `db:"name"`
	LogoUrl             sql.NullString `db:"logo_url"`
	ProviderConnectorID string         `db:"provider_connector_id"`
	ProviderID          int32          `db:"provider_id"`
}

func (q *Queries) UpsertConnector(ctx context.Context, arg UpsertConnectorParams) (Connector, error) {
	row := q.db.QueryRowContext(ctx, upsertConnector,
		arg.Name,
		arg.LogoUrl,
		arg.ProviderConnectorID,
		arg.ProviderID,
	)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LogoUrl,
		&i.ProviderConnectorID,
		&i.ProviderID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
