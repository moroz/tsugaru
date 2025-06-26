-- +goose Up
-- +goose StatementBegin
create extension if not exists citext;

create table users (
  id uuid primary key default uuid7(),
  email citext not null unique,
  password_hash text,
  inserted_at timestamp not null default (now() at time zone 'utc'),
  updated_at timestamp not null default (now() at time zone 'utc')
);

SELECT manage_updated_at('users');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
