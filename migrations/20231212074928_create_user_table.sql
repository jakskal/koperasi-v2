-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
   id serial PRIMARY KEY,
   name VARCHAR (50),
   password TEXT NOT NULL,
   role_id NUMERIC(2) NOT NULL,
   phone INT,
   email VARCHAR (300) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
