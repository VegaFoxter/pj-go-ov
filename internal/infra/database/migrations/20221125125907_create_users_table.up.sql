CREATE TABLE IF NOT EXISTS public.users
(
    id              serial PRIMARY KEY,
    email           varchar(50) NOT NULL,
    password        varchar(100) NOT NULL,
    first_name      varchar(50) NOT NULL,
    second_name     varchar(50) NOT NULL,
    "role"          varchar(50) NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);
