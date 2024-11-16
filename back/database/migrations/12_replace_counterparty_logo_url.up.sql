alter table transaction_enrichments
    drop column if exists counterparty_logo_url;
alter table transaction_enrichments
    add column counterparty_logo text null references images (id);