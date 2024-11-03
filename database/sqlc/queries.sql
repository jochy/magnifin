-- name: GetUserByUsernameAndHashedPassword :one
select *
from users
where username = $1
  and hashed_password = $2
  and deleted_at is null;

-- name: CreateUser :one
insert into users (username, hashed_password)
values ($1, $2)
returning *;

-- name: UpdateUser :one
update users
set hashed_password = $2,
    updated_at      = now()
where id = $1
returning *;

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
where id = $1
returning *;

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
values ($1, $2, $3, $4)
on conflict (provider_id, provider_connector_id) do update
    set name     = excluded.name,
        logo_url = excluded.logo_url
returning *;

-- name: FuzzySearchConnectorsByName :many
select *
from connectors
where name % $1
  and deleted_at is null;

-- name: LikeSearchConnectorsByName :many
select *
from connectors
where name ilike $1
  and deleted_at is null;

-- name: GetConnectorByID :one
select *
from connectors
where id = $1
  and deleted_at is null;
