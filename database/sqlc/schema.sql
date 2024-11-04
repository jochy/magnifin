create table users
(
    id              serial primary key,
    username        text      not null unique,
    hashed_password text      not null,

    created_at      timestamp not null default now(),
    updated_at      timestamp not null default now(),
    deleted_at      timestamp null
);

create table providers
(
    id         serial primary key,
    name       text      not null unique,
    access_key text null,
    secret     text null,
    enabled    boolean   not null default false,

    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp null
);

create table provider_users
(
    id               serial primary key,
    provider_id      integer   not null references providers (id),
    provider_user_id text      not null,
    user_id          integer   not null references users (id),

    created_at       timestamp not null default now(),
    updated_at       timestamp not null default now(),
    deleted_at       timestamp null,

    unique (provider_id, provider_user_id)
);

create table connectors
(
    id                    serial primary key,
    name                  text      not null,
    logo_url              text null,
    provider_connector_id text      not null,
    provider_id           integer   not null references providers (id),

    created_at            timestamp not null default now(),
    updated_at            timestamp not null default now(),
    deleted_at            timestamp null,

    unique (provider_id, provider_connector_id)
);

create table connections
(
    id                     serial primary key,
    provider_users_id      integer   not null references provider_users (id),
    provider_connection_id text      not null,
    connector_id           integer   not null references connectors (id),

    status                 text      not null default 'SYNC_IN_PROGRESS',
    renew_consent_before   timestamp,
    error_message          VARCHAR(1024),
    last_successful_sync   timestamp,

    created_at             timestamp not null default now(),
    updated_at             timestamp not null default now(),
    deleted_at             timestamp null,

    unique (provider_users_id, provider_connection_id)
);

create table redirect_sessions
(
    id                     text primary key,
    provider_connection_id text null,
    internal_connection_id integer null references connections (id),
    created_at             timestamp not null default now()
);

create table accounts
(
    id                  serial primary key,
    connection_id       integer   not null references connections (id),
    provider_account_id text      not null,
    name                text null,
    type                text null,
    currency            text null,
    account_number      text null,
    balance             numeric   not null default 0,

    created_at          timestamp not null default now(),
    updated_at          timestamp not null default now(),
    deleted_at          timestamp null
);
