-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP        DEFAULT NOW(),
    updated_at TIMESTAMP        DEFAULT NULL,
    deleted_at TIMESTAMP        DEFAULT NULL,
    list_id    INTEGER          DEFAULT NULL,
    name       TEXT    NOT NULL,
    bought     BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS lists
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name       TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS lists;
-- +goose StatementEnd