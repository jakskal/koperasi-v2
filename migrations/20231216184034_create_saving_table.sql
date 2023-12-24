-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS saving_types(
   id SERIAL PRIMARY KEY,
   name text UNIQUE NOT NULL,
   created_at timestamp DEFAULT now(),
   created_by INT,
   updated_at timestamp DEFAULT now(),
   updated_by INT,
   deleted_at timestamp,
   deleted_by INT
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS savings(
   id serial PRIMARY KEY,
   user_id INT NOT NULL,
   saving_type_id INT NOT NULL,
   transaction_type_id INT NOT NULL,
   amount BIGINT NOT NULL DEFAULT 0,
   notes text,
   created_at timestamp DEFAULT now(),
   created_by INT,
   updated_at timestamp DEFAULT now(),
   updated_by INT,
   deleted_at timestamp,
   deleted_by INT,
   CONSTRAINT saving_type
      FOREIGN KEY(saving_type_id) 
         REFERENCES saving_types(id),
   CONSTRAINT users
      FOREIGN KEY(user_id) 
         REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS saving_changes(
   id serial PRIMARY KEY,
   saving_id BIGINT NOT NULL,
   transaction_type_id INT NOT NULL,
   notes text,
   amount BIGINT NOT NULL DEFAULT 0,
   changes_notes text,
   created_at timestamp DEFAULT now(),
   created_by int NOT NULL,
   CONSTRAINT savings
      FOREIGN KEY(saving_id) 
         REFERENCES savings(id),
   CONSTRAINT users
      FOREIGN KEY(created_by) 
         REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE saving_changes
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE savings
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE saving_types
-- +goose StatementEnd