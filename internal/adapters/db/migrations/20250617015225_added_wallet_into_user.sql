-- +goose Up
-- +goose StatementBegin
ALTER TABLE users add column wallet numeric(20, 2) NOT NULL DEFAULT 0.00;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users drop column wallet;
-- +goose StatementEnd
