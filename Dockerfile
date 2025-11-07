# Build stage
FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./app-control ./cmd/web

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app-control .
RUN chmod +x ./app-control
COPY ui ./ui
COPY cmd/db/appliances.json ./cmd/db/appliances.json
EXPOSE 4000
CMD ["/app/app-control", "-addr=0.0.0.0:4000", "-base=/app", "-static=/app/static/"]
