alter table transaction_enrichments
    create column counterparty_logo_url text null;

alter table transaction_enrichments
    drop column if exists counterparty_logo;
