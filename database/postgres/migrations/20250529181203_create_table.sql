-- +goose Up
-- +goose StatementBegin
create table roles (
   id serial primary key,
   code varchar(32) not null
);

create table users (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    role_id integer references roles(id),
    created_at timestamp default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table roles;
-- +goose StatementEnd
