-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS conversations (
    id              bigserial       PRIMARY KEY,
    message_id      bigserial       NOT NULL,
    created_at      timestamp       with time zone  DEFAULT     current_timestamp,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS conversations;
-- +goose StatementEnd
