-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    role
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
