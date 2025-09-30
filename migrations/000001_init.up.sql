CREATE TABLE IF NOT EXISTS public.users
(
    id text COLLATE pg_catalog."default" NOT NULL,
    role text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT "1" UNIQUE (username)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

-- Table: public.profiles

-- DROP TABLE IF EXISTS public.profiles;

CREATE TABLE IF NOT EXISTS public.profiles
(
    id text COLLATE pg_catalog."default" NOT NULL,
    user_id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    status text COLLATE pg_catalog."default",
    photo text COLLATE pg_catalog."default",
    CONSTRAINT id PRIMARY KEY (id),
    CONSTRAINT unique_profile UNIQUE (name)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.profiles
    OWNER to postgres;

-- Table: public.posts

-- DROP TABLE IF EXISTS public.posts;

CREATE TABLE IF NOT EXISTS public.posts
(
    id text COLLATE pg_catalog."default" NOT NULL,
    profile_id text COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    date timestamp with time zone NOT NULL,
    likes_count integer,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.posts
    OWNER to postgres;

-- Table: public.likes

-- DROP TABLE IF EXISTS public.likes;

CREATE TABLE IF NOT EXISTS public.likes
(
    profile_id text COLLATE pg_catalog."default" NOT NULL,
    post_id text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT unique_like UNIQUE (profile_id, post_id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.likes
    OWNER to postgres;