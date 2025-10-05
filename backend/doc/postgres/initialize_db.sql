CREATE TABLE public.app_logs
(
    id         bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    timestamp  timestamp with time zone            NOT NULL DEFAULT now(),
    level      character varying                   NOT NULL,
    message    text                                NOT NULL,
    attributes jsonb,
    CONSTRAINT app_logs_pkey PRIMARY KEY (id)
);
CREATE TABLE public.currencies
(
    code character varying NOT NULL,
    CONSTRAINT currencies_pkey PRIMARY KEY (code)
);
CREATE TABLE public.transaction_categories
(
    id          bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    name        character varying                   NOT NULL,
    description character varying                   NOT NULL,
    CONSTRAINT transaction_categories_pkey PRIMARY KEY (id)
);
CREATE TABLE public.transaction_status
(
    id   bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    name character varying                   NOT NULL,
    CONSTRAINT transaction_status_pkey PRIMARY KEY (id)
);
CREATE TABLE public.transaction_sub_categories
(
    id              bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    parent_category bigint                              NOT NULL,
    name            character varying                   NOT NULL,
    CONSTRAINT transaction_sub_categories_pkey PRIMARY KEY (id),
    CONSTRAINT transaction_sub_categories_parent_category_fkey FOREIGN KEY (parent_category) REFERENCES public.transaction_categories (id)
);
CREATE TABLE public.transaction_types
(
    id   bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    name character varying                   NOT NULL,
    CONSTRAINT transaction_types_pkey PRIMARY KEY (id)
);
CREATE TABLE public.users
(
    id                   bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    username             BYTEA                               NOT NULL,
    username_search_hash BYTEA                               NOT NULL,
    hashed_password      text                                NOT NULL,
    created_at           timestamp with time zone            NOT NULL DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username_hash_unique UNIQUE (username_search_hash)
);
CREATE TABLE public.transactions
(
    id           bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
    user_id      bigint                              NOT NULL, -- Added user_id
    amount       BYTEA                               NOT NULL,
    category     bigint                              NOT NULL,
    sub_category bigint,
    date         date                                NOT NULL,
    description  BYTEA,
    status       bigint                              NOT NULL,
    currency     character varying                   NOT NULL,
    type         bigint                              NOT NULL,
    essential    boolean                             NOT NULL,
    CONSTRAINT transactions_pkey PRIMARY KEY (id),
    CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users (id),
    CONSTRAINT transactions_currency_fkey FOREIGN KEY (currency) REFERENCES public.currencies (code),
    CONSTRAINT transactions_category_fkey FOREIGN KEY (category) REFERENCES public.transaction_categories (id),
    CONSTRAINT transactions_sub_category_fkey FOREIGN KEY (sub_category) REFERENCES public.transaction_sub_categories (id),
    CONSTRAINT transactions_status_fkey FOREIGN KEY (status) REFERENCES public.transaction_status (id),
    CONSTRAINT transactions_type_fkey FOREIGN KEY (type) REFERENCES public.transaction_types (id)
);