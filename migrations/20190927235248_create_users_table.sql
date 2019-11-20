-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id              bigserial       PRIMARY KEY,
    name            varchar(100)    NOT NULL,
    email           varchar(255)    NOT NULL    UNIQUE,
    user_name       varchar(50)     NOT NULL,
    password        varchar(255)    NOT NULL,
    created_at      timestamp       DEFAULT     current_timestamp,
    updated_at      timestamp       DEFAULT     current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
