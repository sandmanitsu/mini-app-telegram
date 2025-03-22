-- +goose Up
CREATE TABLE IF NOT EXISTS public.tg_user (
	id serial,
    tg_user_id bigint NOT NULL UNIQUE,
	username text NOT NULL,
	first_name text NOT NULL,
    last_name text NOT NUll,
	chat_id bigint NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);


-- +goose Down
DROP TABLE public.user;