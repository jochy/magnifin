create table redirect_sessions
(
    id                     text primary key,
    provider_connection_id text null,
    internal_connection_id integer null references connections (id),
    created_at             timestamp not null default now()
);