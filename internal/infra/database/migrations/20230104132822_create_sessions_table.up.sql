CREATE TABLE IF NOT EXISTS public.sessions
(
    user_id int         NOT NULL,
    uuid    varchar(50) NOT NULL,
    CONSTRAINT auths_pkey PRIMARY KEY (user_id, uuid)
);

