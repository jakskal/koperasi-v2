-- +goose Up
-- +goose StatementBegin
INSERT INTO users(id, name, password, role_id, status, phone, email)
VALUES(1, 'admin koperasi', '$2a$10$xNm.VNy/fb/2GjndMgzoueWmAowoTlYDLYOsLKiz8taooeXJXaYpa', 0, 1, 89999999999, 'admin@koperasi.com')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE from users where id = 1
-- +goose StatementEnd
