-- name: GetUserByUsername :one
select *
from users
where username = $1
  and deleted_at is null;

-- name: CreateUser :one
insert into users (username, hashed_password)
values ($1, $2) returning *;

-- name: UpdateUser :one
update users
set hashed_password = $2,
    updated_at      = now()
where id = $1 returning *;

-- name: GetUserByID :one
select *
from users
where id = $1
  and deleted_at is null;

-- name: ListProviders :many
select *
from providers
where deleted_at is null;

-- name: UpdateProvider :one
update providers
set name       = $2,
    access_key = $3,
    secret     = $4,
    enabled    = $5,
    updated_at = now()
where id = $1 returning *;

-- name: GetProviderByID :one
select *
from providers
where id = $1
  and deleted_at is null;

-- name: GetProviderByName :one
select *
from providers
where name = $1
  and deleted_at is null;

-- name: UpsertConnector :one
insert into connectors (name, logo_url, provider_connector_id, provider_id)
values ($1, $2, $3, $4) on conflict (provider_id, provider_connector_id) do
update
    set name = excluded.name,
    logo_url = excluded.logo_url
    returning *;

-- name: FuzzySearchConnectorsByName :many
select connectors.*
from connectors
         inner join providers on connectors.provider_id = providers.id
where connectors.name % $1
  and connectors.deleted_at is null
  and providers.deleted_at is null
  and providers.enabled = true;

-- name: LikeSearchConnectorsByName :many
select connectors.*
from connectors
         inner join providers on connectors.provider_id = providers.id
where connectors.name ilike $1
  and connectors.deleted_at is null
  and providers.deleted_at is null
  and providers.enabled = true;

-- name: GetConnectorByID :one
select *
from connectors
where id = $1
  and deleted_at is null;

-- name: GetProviderUserByProviderIDAndUserID :one
select *
from provider_users
where provider_id = $1
  and user_id = $2
  and deleted_at is null;

-- name: CreateProviderUser :one
insert into provider_users (provider_id, user_id, provider_user_id)
values ($1, $2, $3) returning *;

-- name: GetConnectionByProviderUserIDAndProviderConnectionID :one
select *
from connections
where provider_users_id = $1
  and provider_connection_id = $2
  and deleted_at is null;

-- name: GetProviderUserByID :one
select *
from provider_users
where id = $1
  and deleted_at is null;

-- name: CreateConnection :one
insert into connections (provider_users_id, provider_connection_id, connector_id, status, renew_consent_before,
                         error_message, last_successful_sync)
values ($1, $2, $3, $4, $5, $6, $7) returning *;

-- name: UpdateConnection :one
update connections
set status                 = $2,
    renew_consent_before   = $3,
    error_message          = $4,
    last_successful_sync   = $5,
    provider_connection_id = $6,
    updated_at             = now()
where id = $1 returning *;

-- name: GetConnectionByID :one
select *
from connections
where id = $1
  and deleted_at is null;

-- name: StoreRedirectSessions :exec
insert into redirect_sessions (id, provider_connection_id, internal_connection_id, user_id)
values ($1, $2, $3, $4);

-- name: GetRedirectSessionByID :one
select *
from redirect_sessions
where id = $1;

-- name: GetAccountByConnectionIDAndProviderAccountID :one
select *
from accounts
where connection_id = $1
  and provider_account_id = $2
  and deleted_at is null;

-- name: CreateAccount :one
insert into accounts (connection_id, provider_account_id, name, type, currency, account_number, balance,
                      bank_account_id)
values ($1, $2, $3, $4, $5, $6, $7, $8) returning *;

-- name: UpdateAccount :one
update accounts
set name                = $2,
    type                = $3,
    currency            = $4,
    account_number      = $5,
    balance             = $6,
    provider_account_id = $7,
    bank_account_id     = $8,
    updated_at          = now()
where id = $1 returning *;

-- name: FindTransactionByAccountIDAndProviderTransactionID :one
select *
from transactions
where account_id = $1
  and provider_transaction_id = $2
  and deleted_at is null;

-- name: CreateTransaction :one
insert into transactions (account_id, provider_transaction_id, bank_transaction_id, amount, currency, direction, status,
                          operation_at, counterparty_name, counterparty_account,
                          reference)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning *;

-- name: UpdateTransaction :one
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
where id = $1 returning *;

-- name: UpdateConnectionStatus :one
update connections
set status = $2
where id = $1 returning *;

-- name: ListConnectionsToSync :many
select *
from connections
where ((last_successful_sync is null and created_at < now() - interval '1 hour') or
       (last_successful_sync < now() - interval '11 hours'))
  and deleted_at is null;

-- name: ListConnectionsByUserID :many
select *
from connections
where provider_users_id in (select id from provider_users where user_id = $1 and deleted_at is null)
  and deleted_at is null;

-- name: ListAccountsByConnectionID :many
select *
from accounts
where connection_id = $1
  and deleted_at is null;

-- name: GetConnectionByIDAndUserID :one
select *
from connections
where connections.id = $1
  and provider_users_id in (select provider_users.id from provider_users where user_id = $2 and deleted_at is null)
  and deleted_at is null;

-- name: DeleteConnectionByID :exec
update connections
set deleted_at = now()
where id = $1;

-- name: DeleteAccountByConnectionID :exec
update accounts
set deleted_at = now()
where connection_id = $1;

-- name: DeleteTransactionsByConnectionID :exec
update transactions
set deleted_at = now()
where account_id in (select id from accounts where connection_id = $1);

-- name: DeleteTransactionsEnrichmentsByConnectionID :exec
update transaction_enrichments
set deleted_at = now()
where transaction_id in
      (select id from transactions where account_id in (select id from accounts where connection_id = $1));