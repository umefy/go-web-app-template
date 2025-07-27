-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    email varchar(255) unique not null,
    age int not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

CREATE TRIGGER updated_at_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION updated_at_trigger();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
