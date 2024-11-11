create table transactions
(
    id                      serial primary key,
    account_id              integer   not null references accounts (id),
    provider_transaction_id text      not null,
    bank_transaction_id     text null,

    amount                  numeric   not null,
    currency                text      not null,
    direction               text      not null,
    status                  text      not null,
    operation_at            timestamp not null,

    counterparty_name       text null,
    counterparty_account    text null,

    reference               text null,

    created_at              timestamp not null default now(),
    updated_at              timestamp not null default now(),
    deleted_at              timestamp null
);

create table transaction_enrichments
(
    id                    serial primary key,
    transaction_id        integer not null references transactions (id),

    category              integer null references categories (id),
    reference             text null,
    counterparty_name     text null,
    counterparty_logo_url text null,

    deleted_at            timestamp null
);