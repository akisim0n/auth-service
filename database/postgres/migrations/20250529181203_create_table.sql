-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    role int references roles(id)
);

create table roles (
    id serial primary key,
    code varchar(32) not null,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
drop table roles;
-- +goose StatementEnd
