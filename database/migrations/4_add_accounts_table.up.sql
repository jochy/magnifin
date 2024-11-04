create table accounts
(
    id             serial primary key,
    connection_id  integer   not null references connections (id),
    name           text      null,
    type           text      null,
    currency       text      null,
    account_number text      null,
    balance        numeric   not null default 0,

    created_at     timestamp not null default now(),
    updated_at     timestamp not null default now(),
    deleted_at     timestamp null
);