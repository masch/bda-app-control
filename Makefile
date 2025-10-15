.PHONY: run build test

db-init:
	podman volume create bda-pgdata && \
	podman run -d \
	--name bda-db \
	-e POSTGRES_USER=tabaquillo \
	-e POSTGRES_PASSWORD=tabaquillo \
	-e POSTGRES_DB=bosque \
	-p 5432:5432 \
	-v bda-pgdata:/var/lib/postgresql/data \
	postgres:16

db-start:
	podman start bda-db

db-stop:
	podman stop bda-db

init:
	$(MAKE) db-start

run:
	go run ./cmd/web

build:
	go build -o ./bin/app ./cmd/web

test:
	go test ./...
