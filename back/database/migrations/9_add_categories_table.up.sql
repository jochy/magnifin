create table categories
(
    id                serial primary key,
    name              text not null,
    user_id           integer null references users (id),
    color             text not null,
    icon              text not null,
    include_in_budget boolean not null default true,

    deleted_at        timestamp null
);

create table category_rules
(
    id          serial primary key,
    category_id integer not null references categories (id),
    rule        text    not null,

    created_at  timestamp not null default now(),
    deleted_at  timestamp null
);

insert into categories (name, color, icon)
values ('Auto & Transport', '#FFC107', 'directions_car'),
       ('Subscriptions and Bills', '#FF5722', 'receipt'),
       ('Cash & Checks', '#795548', 'attach_money'),
       ('Business & Work', '#607D8B', 'work'),
       ('Food & Drink', '#4CAF50', 'restaurant'),
       ('Investment', '#9C27B0', 'trending_up'),
       ('Health', '#E91E63', 'local_hospital'),
       ('Loan Repayment', '#FF9800', 'credit_card'),
       ('Income', '#FF5722', 'money_off'),
       ('Taxes', '#FFC107', 'monet'),
       ('Transfers', '#2196F3', 'swap_horiz'),
       ('Essential Needs', '#FFEB3B', 'local_grocery_store');
