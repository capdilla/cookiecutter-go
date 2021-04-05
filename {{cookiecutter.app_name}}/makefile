NAME={{cookiecutter.app_name}}
VERSION=0.0.1
DB_NAME=duna-{{cookiecutter.app_name}}
DB_PORT={{cookiecutter.db_dev_port}}
MIGRATION_DIR=storage/migrations
## get extra arguments and filter out commands from args
args = $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps-save
## deps-save: Godeps save
deps-save:
	@godep save

.PHONY: deps
## deps: Godeps restore
deps:
	@godep restore

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: migration
migration:
	goose -dir ${MIGRATION_DIR} create $(call args,defaultstring) sql

migrate-status:
	goose -dir ${MIGRATION_DIR} postgres "user=postgres dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" status

migrate-up:
	goose -dir ${MIGRATION_DIR} postgres "user=postgres dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" up

migrate-down:
	goose -dir ${MIGRATION_DIR} postgres "user=postgres dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" down

migrate-rollback:
	goose -dir ${MIGRATION_DIR} postgres "user=postgres dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" reset

migrate-reset:
	goose -dir ${MIGRATION_DIR} postgres "user=postgres dbname=${DB_NAME} port=${DB_PORT} sslmode=disable" reset

generate-mocks:
	mockgen -source services/crypto.go -destination mocks/mock_services/mock_crypto_service.go
	mockgen -source repository/merchant_private_key.go -destination mocks/mock_repositories/mock_merchant_respository.go
	mockgen -source repository/user.go -destination mocks/mock_repositories/mock_user_respository.go

## default that allows accepting extra args
%:
    @:

swagger:
	swagger generate spec --input=="main.go" --output="swagger.yml"

serve-swagger:
	swagger serve swagger.yml --flavor="swagger" --doc-url="docs"