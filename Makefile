createdb:
	docker exec -it personal-finance_postgres createdb --username=admin --owner=admin personal_finance

dropdb:
	docker exec -it personal-finance_postgres dropdb personal_finance

migrate:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate-up:
	migrate -path internal/db/migrations -database "postgresql://admin:password@localhost:5432/personal_finance?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migrations -database "postgresql://admin:password@localhost:5432/personal_finance?sslmode=disable" -verbose down