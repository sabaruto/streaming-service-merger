-- +goose Up
-- +goose StatementBegin
ALTER TABLE stream_creds
    ADD auth_code VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stream_creds
    DROP COLUMN auth_code;
-- +goose StatementEnd
