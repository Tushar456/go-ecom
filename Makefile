migrate-create:
	migrate create -ext sql -dir db/migration -seq init_schema
migrate-up:
	 migrate -path db/migration -database "mysql://root:bunty@123@tcp(localhost:3306)/ecom" --verbose up

migrate-down:
	 migrate -path db/migration -database "mysql://root:bunty@123@tcp(localhost:3306)/ecom" --verbose down

migrate-force:
	 migrate -path db/migration -database "mysql://root:bunty@123@tcp(localhost:3306)/ecom" --verbose force 1

sqlccodegen:
	sqlc generate

.PHONY: migrate-create migrate-up migrate-down migrate-force sql-codegen