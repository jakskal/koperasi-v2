run-be:
	go run cmd/api/main.go
migrate-status:
	goose -dir migrations postgres "host=localhost user=user_default database=koperasi port=4432 sslmode=disable password=password" status
migrate-all:
	goose -dir migrations postgres "host=localhost user=user_default database=koperasi port=4432 sslmode=disable password=password" up
migrate-one:
	goose -dir migrations postgres "host=localhost user=user_default database=koperasi port=4432 sslmode=disable password=password" up-by-one 
migrate-down:
	goose -dir migrations postgres "host=localhost user=user_default database=koperasi port=4432 sslmode=disable password=password" down 
