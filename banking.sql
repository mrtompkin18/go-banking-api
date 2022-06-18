create table if not exists account
(
    account_id   int auto_increment
        primary key,
    customer_id  int                                   not null,
    opening_date timestamp default current_timestamp() null,
    account_type varchar(10)                           null,
    balance      decimal   default 0                   not null,
    status       int       default 0                   not null
);

create table if not exists customer
(
    customer_id   int auto_increment
        primary key,
    name          varchar(255)                          null,
    date_of_birth timestamp default current_timestamp() null,
    city          varchar(255)                          null,
    zipcode       varchar(255)                          null,
    status        int       default 0                   null
);

