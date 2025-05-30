-- +goose Up
-- +goose StatementBegin
insert into roles values (
    'admin',
    'user'
);

insert into users values (
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
