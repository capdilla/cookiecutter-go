-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS hello;
create table hello
(
    id SERIAL,
    name varchar(255),
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table hello;