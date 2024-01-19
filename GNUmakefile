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
	go test -v  -coverprofile=coverage.out ./cmd/... && go tool cover -func=coverage.out

# doc:
# 	func init() {
# 		@echo "Generating docs"
# 		err := doc.GenMarkdownTree(cmd, "/tmp")
# 		if err != nil {
# 			log.Fatal(err)
# 		}
# 	}
	

generate:
	golangci-lint run 
	go test -v  -coverprofile=coverage.out ./cmd/... && go tool cover -func=coverage.out
	goreleaser release --snapshot --clean

submodules:
	@git submodule sync
	@git submodule update --init --recursive
	@git config core.hooksPath githooks
	@git config submodule.recurse true
