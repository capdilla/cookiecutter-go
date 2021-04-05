-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS auth;
create table auth
(
    id SERIAL,
    password varchar(255) NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table auth;