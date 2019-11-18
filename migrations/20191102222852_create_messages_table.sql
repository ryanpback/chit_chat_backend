-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages (
    id              bigserial       PRIMARY KEY,
    sender_id       bigserial       NOT NULL,
    receiver_id     bigserial       NOT NULL,
    message         text            NOT NULL,
    created_at      timestamp       with time zone DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd