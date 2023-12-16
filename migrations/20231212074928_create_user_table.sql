-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_attributes(
   id serial PRIMARY KEY,
   member_id text,
   is_active_member BOOLEAN DEFAULT TRUE,
   name VARCHAR (100),
   join_date timestamp DEFAULT now(),
   birth timestamp,
   birth_place Varchar (50),
   address VARCHAR (300),
   profession Varchar (300),
   created_at timestamp DEFAULT now(),
   created_by INT,
   updated_at timestamp DEFAULT now(),
   updated_by INT,
   deleted_at timestamp,
   deleted_by INT
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
   id serial PRIMARY KEY,
   name VARCHAR (50),
   email VARCHAR (300) UNIQUE NOT NULL,
   password TEXT NOT NULL,
   role_id NUMERIC(2) NOT NULL,
   status NUMERIC(2) NOT NULL default 0,
   phone BIGINT,
   attribute_id int,
   created_at timestamp DEFAULT now(),
   created_by INT,
   updated_at timestamp DEFAULT now(),
   updated_by INT,
   deleted_at timestamp,
   deleted_by INT,
   FOREIGN KEY(attribute_id) REFERENCES user_attributes(id)

);
-- +goose StatementEnd




-- +goose Down

-- +goose StatementBegin
DROP TABLE member_attributes
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
