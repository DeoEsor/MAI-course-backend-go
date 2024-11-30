-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS pet (
    id              uuid            not null,
    status          varchar(50)     not null,
    name            varchar(128)    not null,
    passport        jsonb           not null,
    created_at      timestamp without time zone not null ,
    updated_at      timestamp without time zone not null
);

create index idx_name_pet on pet(name)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd


