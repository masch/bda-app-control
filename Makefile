.PHONY: env-init
env-init:
	@if [ -f .env ]; then \
		echo ".env already exists, skipping creation."; \
	else \
		cp .env_template .env && echo "Created .env from .env_template"; \
	fi

.PHONY: db-init
db-init:
	podman pull docker.io/library/postgres:16-alpine && \
	podman volume create bda-pgdata && \
	podman run -d \
	--name bda-db \
	-e POSTGRES_USER=tabaquillo \
	-e POSTGRES_PASSWORD=tabaquillo \
	-e POSTGRES_DB=bosque \
	-p 5432:5432 \
	-v bda-pgdata:/var/lib/postgresql/data \
	postgres:16-alpine

.PHONY: db-start
db-start:
	podman start bda-db

.PHONY: db-stop
db-stop:
	podman stop bda-db

.PHONY: db-connect
db-connect:
	podman exec -it bda-db psql -U tabaquillo -d bosque

.PHONY: init
init:
	$(MAKE) db-start

.PHONY: run-internal
run-internal:
	go run ./cmd/web -addr="0.0.0.0:4000" -base="/bosquesdeagua/" -static="/static/"

.PHONY: run-exteral
run-exteral:
	go run ./cmd/web -addr="0.0.0.0:4000" -base="/app" -static="/app/static/"

.PHONY: run-containers-services
run-containers-services:
	podman-compose -f "docker-compose.yml" up -d grafana proxy db influxdb --build --force-recreate

.PHONY: run-containers-all
run-containers-all:
	podman-compose -f "docker-compose.yml" up -d --build --force-recreate

.PHONY: build
build:
	go build -o ./bin/app ./cmd/web

.PHONY: test
test:
	go test ./...

