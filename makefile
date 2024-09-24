include .env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Запуск миграций вверх
migrate-up:
	migrate -database $(DB_URL) -path migrations up

# Откатить последнюю миграцию
migrate-down:
	migrate -database $(DB_URL) -path migrations down

# Откатить все миграции
migrate-down-all:
	migrate -database $(DB_URL) -path migrations down -all

migrate-to:
	migrate -database $(DB_URL) -path migrations goto $(version)

create-migration:
	migrate create -ext sql -dir migrations -seq $(name)
