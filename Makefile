.PHONY: test build run clean help

# Help target
help:
	@echo "Available targets:"
	@echo "  make test       - Run all tests"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make test-v     - Run tests with verbose output"
	@echo "  make test-cover - Run tests with coverage report"

test:
	cd demo/web/cmd && go test ./...

test-v:
	cd demo/web/cmd && go test -v ./...


test-cover:
	cd demo/web/cmd && go test -v -cover ./...

build:
	cd demo/web/cmd && go build -o app

run: build
	./demo/web/cmd/app

clean:
	cd demo/web/cmd && rm -f app
	go clean

test-html:
	cd demo/web/cmd && go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
