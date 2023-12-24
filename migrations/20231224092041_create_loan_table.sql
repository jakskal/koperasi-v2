-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_types(
    id SERIAL PRIMARY KEY,
    name text UNIQUE NOT NULL,
    ratio_percentage NUMERIC(3, 1),
    created_at timestamp DEFAULT now(),
    created_by INT,
    updated_at timestamp DEFAULT now(),
    updated_by INT,
    deleted_at timestamp,
    deleted_by INT
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loans(
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    loan_type_id INT NOT NULL,
    name varchar(64),
    amount DOUBLE PRECISION NOT NULL DEFAULT 0,
    ratio_percentage NUMERIC(3, 1),
    total_ratio_amount DOUBLE PRECISION,
    installment_qty_target int,
    notes text,
    created_at timestamp DEFAULT now(),
    created_by INT,
    updated_at timestamp DEFAULT now(),
    updated_by INT,
    deleted_at timestamp,
    deleted_by INT,
    CONSTRAINT loan_type
        FOREIGN KEY(loan_type_id) 
            REFERENCES loan_types(id),
    CONSTRAINT users
        FOREIGN KEY(user_id) 
            REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_changes(
    id serial PRIMARY KEY,
    loan_id BIGINT NOT NULL,
    name varchar(64),
    amount DOUBLE PRECISION NOT NULL DEFAULT 0,
    ratio_percentage NUMERIC(3, 1),
    total_ratio_amount DOUBLE PRECISION,
    installment_qty_target int,
    notes text,
    changes_notes text,
    created_at timestamp DEFAULT now(),
    created_by int NOT NULL,
    CONSTRAINT loans
        FOREIGN KEY(loan_id) 
            REFERENCES loans(id),
    CONSTRAINT users
        FOREIGN KEY(created_by) 
            REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_installments(
    id serial PRIMARY KEY,
    loan_id BIGINT NOT NULL,
    transaction_type_id INT NOT NULL,
    payment_date timestamp NOT NULL DEFAULT now(),
    principal_amount DOUBLE PRECISION,
    interest_amount DOUBLE PRECISION,
    notes text,
    created_at timestamp DEFAULT now(),
    created_by INT,
    updated_at timestamp DEFAULT now(),
    updated_by INT,
    deleted_at timestamp,
    deleted_by INT,
    CONSTRAINT loans
        FOREIGN KEY(loan_id) 
            REFERENCES loans(id)
)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS loan_installment_changes(
    id serial PRIMARY KEY,
    loan_installment_id BIGINT NOT NULL,
    transaction_type_id INT NOT NULL,
    payment_date timestamp NOT NULL DEFAULT now(),
    principal_amount DOUBLE PRECISION,
    interest_amount DOUBLE PRECISION,
    notes text,
    changes_notes text,
    created_at timestamp DEFAULT now(),
    created_by INT,
    CONSTRAINT loan_installments
        FOREIGN KEY(loan_installment_id) 
            REFERENCES loan_installments(id),
    CONSTRAINT users
        FOREIGN KEY(created_by) 
            REFERENCES users(id)
)
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE loan_installment_changes
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE loan_installments
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE loan_changes
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE loans
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE loan_types
-- +goose StatementEnd