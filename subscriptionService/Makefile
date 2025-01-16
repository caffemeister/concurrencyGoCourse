include .env
export $(shell sed 's/=.*//' .env)

build:
	@go build -o $(BINARY_NAME) ./cmd/web
	@echo "Back end built!"

run: build
	@echo "Starting..."
	@set DSN=$(DSN) && set REDIS=$(REDIS) && $(BINARY_NAME)
	@echo "Back end started!"

clean:
	@echo "Cleaning..."
	@DEL /F /Q $(BINARY_NAME) 2>nul || true
	@go clean
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping..."
	@taskkill /IM $(BINARY_NAME) /F 2>nul || echo "Process not running"
	@echo "Stopped back end"

restart: stop start

test:
	@echo "Testing..."
	@go test -v ./...
