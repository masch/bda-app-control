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

run:
	go run ./cmd/web -addr="0.0.0.0:4000" -base="/bosquesdeagua"

build:
	go build -o ./bin/app ./cmd/web

test:
	go test ./...
