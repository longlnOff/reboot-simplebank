DB_URL_MIGRATE = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
PATH_SCHEMA = sql/schema

postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:16-alpine

start_postgres:
	docker start postgres

stop_postgres:
	docker stop postgres

create_db:
	docker exec -it postgres createdb --user=root --owner=root simple_bank

drop_db:
	docker exec -it postgres dropdb simple_bank

1create_migrate:
	migrate create -ext sql -dir sql/schema -seq create_table

up_migrate:
	migrate -path $(PATH_SCHEMA) -database $(DB_URL_MIGRATE) -verbose up

one_up_migrate:
	migrate -path $(PATH_SCHEMA) -database $(DB_URL_MIGRATE) -verbose up 1

down_migrate:
	migrate -path $(PATH_SCHEMA) -database $(DB_URL_MIGRATE) -verbose down

one_down_migrate:
	migrate -path $(PATH_SCHEMA) -database $(DB_URL_MIGRATE) -verbose down 1
