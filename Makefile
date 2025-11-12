.PHONY: run build test

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

db-start:
	podman start bda-db

db-stop:
	podman stop bda-db

db-connect:
	podman exec -it bda-db psql -U tabaquillo -d bosque

init:
	$(MAKE) db-start

run-internal:
	go run ./cmd/web -addr="0.0.0.0:4000" -base="/bosquesdeagua/" -static="/static/"

run-exteral:
	go run ./cmd/web -addr="0.0.0.0:4000" -base="/app" -static="/app/static/"

run-containers-services:
	podman-compose -f "docker-compose.yml" up -d grafana proxy db influxdb --build --force-recreate

run-containers-all:
	podman-compose -f "docker-compose.yml" up -d --build --force-recreate

build:
	go build -o ./bin/app ./cmd/web

test:
	go test ./...
