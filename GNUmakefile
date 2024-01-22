TEST?=$$(go list ./... |grep -v 'vendor')
default: build

lint:
	golangci-lint run -v

lintWithFix:
	golangci-lint run --fix

install:
	golangci-lint run
	go build -o /usr/local/bin/cav .

run:
	go run .

build:
	goreleaser release --snapshot --clean

test:
	go test -coverprofile=coverage.out ./cmd/... && go tool cover -func=coverage.out

 doc:
	@echo "Generating documentation..."
	go run main.go gen-doc	

generate:
	golangci-lint run
	go generate ./...

submodules:
	@git submodule sync
	@git submodule update --init --recursive
	@git config core.hooksPath githooks
	@git config submodule.recurse true
