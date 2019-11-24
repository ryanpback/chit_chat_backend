-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS conversations_messages (
    id                  bigserial       PRIMARY KEY,
    conversation_id     bigserial       NOT NULL,
    message_id          bigserial       NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS conversations_messages;
-- +goose StatementEnd
