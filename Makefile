# name app
APP_NAME = server
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR = sql/schema

run:
	go run ./cmd/${APP_NAME}/main.go

dockerup:
	docker-compose up -d

dockerkill:
	docker-compose kill

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset








.PHONY: dockerup dockerkill upse downse resetse run