-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS conversations_users (
    id                  bigserial       PRIMARY KEY,
    user_id             bigserial       NOT NULL,
    conversation_id     text            NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS conversations_users;
-- +goose StatementEnd
