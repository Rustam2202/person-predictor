run:
	go run ./cmd/main.go -confpath=./

swag:
	swag fmt
	swag init -g ./internal/server/server.go

lint:
	golangci-lint run
