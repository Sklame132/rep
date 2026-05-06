include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d rep-postgres

env-down:
	@docker compose down rep-postgres

env-cleanup:
	@read -p "Опасность утери данных. Вы точно хотите удалить все volume файлы окружения? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down rep-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена."; \
	fi;

env-port-forward:
	@docker compose up -d port-forwarder;

env-port-close: 
	@docker compose down port-forwarder;


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутвует параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \

	docker compose run --rm rep-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутвует параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \

	docker compose run --rm rep-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@rep-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

rep-run: 
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/rep/main.go

rep-deploy:
	@docker compose up -d --build rep

ps:
	@docker compose ps