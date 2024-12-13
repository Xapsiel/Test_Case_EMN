FROM golang:1.22.8 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o test_case ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/test_case ./

COPY configs/config.yml ./configs/config.yml
COPY schema/* ./schema/
COPY .env ./
ENTRYPOINT ["./test_case"]
