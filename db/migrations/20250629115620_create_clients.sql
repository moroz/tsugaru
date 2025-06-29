-- +goose Up
-- +goose StatementBegin
create table clients (
  id uuid primary key default uuid7(),
  name text not null,
  owner_id uuid not null references users (id) on delete cascade,
  redirect_urls text[] not null default '{}',
  inserted_at timestamp not null default (now() at time zone 'utc'),
  updated_at timestamp not null default (now() at time zone 'utc')
);

SELECT manage_updated_at('clients');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table clients;
-- +goose StatementEnd
