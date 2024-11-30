-- +goose Up
-- +goose StatementBegin
create unique index
    if not exists
    idx_uniq_id_pet on pet(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
