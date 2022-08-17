CREATE TABLE public.customers (
    id         bigserial     NOT NULL,
    first_name varchar(100)  NULL,
    last_name  varchar(100)  NULL,
    birth_date timestamp     NULL,
    gender     varchar(1)    NULL,
    email      varchar(100)  NOT NULL,
    address    varchar(200)  NOT NULL,
    CONSTRAINT customers_pkey PRIMARY KEY (id)
);