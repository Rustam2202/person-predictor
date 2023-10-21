run:
	go run ./cmd/main.go -confpath=./
build:
	go build -o bin/device-manager ./cmd/main.go

swag:
	swag fmt
	swag init -g ./internal/server/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html

lint:
	golangci-lint run

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out
