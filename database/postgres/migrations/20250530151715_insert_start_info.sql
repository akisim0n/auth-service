-- +goose Up
-- +goose StatementBegin
insert into roles(code) values
    (
    'admin'
    ),
    (
    'user'
    );

insert into users(name,email,password,role_id) values (
    'Daniil',
    'dak@gmail.com',
    'Egalam47',
    (select id from roles where code = 'admin')
), (
    'Jojo',
    'jojo@gmail.com',
    '123345678',
    (select id from roles where code = 'user')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from users;
delete from roles;
-- +goose StatementEnd
