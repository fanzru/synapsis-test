include .env

run:
	go run ./cmd/main.go

setup:
	go mod download

# Database Migration
# create migration file
migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)

# migration up (craete all table)
migrate-up:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migrations up

# migration down (drop all table)
migrate-down:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migrations down -all

# rollback migration
migrate-rollback:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migrations down $(N)

# migration force with version (craete all table)
migrate-force:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migrations force $(VERSION)

mock-usecase:
	mockgen -package=mock -source=app/${SUBDOMAIN}/usecase/${FILE}.go -destination=app/${SUBDOMAIN}/usecase/usecase_mock/${FILE}.go

mock-repo:
	mockgen -package=mock -source=app/${SUBDOMAIN}/repo/${FILE}.go  -destination=app/${SUBDOMAIN}/repo/repo_mock/${FILE}.go