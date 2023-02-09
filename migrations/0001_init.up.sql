CREATE TABLE public.users
(
    id        SERIAL PRIMARY KEY,
    tg_id     TEXT,
    is_admin  bool,
    create_at Timestamp
);

CREATE TABLE public.division
(
    id        serial primary key,
    name      text,
    active    bool,
    creat_at  timestamp,
    update_at timestamp,
    delete_at timestamp
);

CREATE TABLE public.project
(
    id serial primary key,
    division_id int not null,
    constraint division_fk foreign key (division_id) REFERENCES public.divison(id),
    name text,
    active bool,
    creat_at timestamp,
    delete_at timestamp
)
