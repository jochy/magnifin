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
