include config/.env 

init:
	docker compose --env-file config/.env up db -d
	sleep 2
	$(MAKE) migrateup
	docker compose --env-file config/.env up --build -d

start:
	docker compose --env-file config/.env up -d

migrateup:
	migrate -path migrations -database 'postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable' -verbose up

migratedown:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

send:
	docker compose --env-file config/.env up producer

generate:
	sqlc -f config/sqlc.yml generate
	swag init -g cmd/order/main.go 

lint:
	golangci-lint --color auto -v run --fix 

cloc:
	gocloc .

stop:
	docker compose down
	docker compose down producer