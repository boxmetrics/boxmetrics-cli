.PHONY: run test coverage coverage-html prebuild build clean

prebuild:
	@echo "Building dependencies"
	mkdir -p bin
	cd bin && go build -v ../internal/app/server.go
	chmod +x scripts/*.sh

run:
	@echo "Starting module"
	make prebuild
	go run main.go

test:
	@echo "Running module test"
	go test -v ./...

coverage:
	@echo "Running coverage"
	go test -v -cover ./...

coverage-html:
	@echo "Running coverage html"
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html

build:
	@echo "Building module to ./bin"
	make prebuild
	cd bin && go build -v ../.

clean:
	@echo "Cleaning workspace"
	rm bin/*
	go mod tidy
	go clean ./...
