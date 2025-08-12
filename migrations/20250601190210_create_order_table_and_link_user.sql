-- +goose Up
-- +goose StatementBegin
create table if not exists orders (
  id serial primary key,
  user_id int not null,
  amount_cents bigint not null,
  version bigint not null default 0,
  created_at timestamptz default now(),
  updated_at timestamptz default now()
);

CREATE TRIGGER updated_at_trigger
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE FUNCTION updated_at_trigger();

alter table orders add constraint fk_orders_user_id foreign key (user_id) references users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table orders drop constraint fk_orders_user_id;
drop table if exists orders;
-- +goose StatementEnd
