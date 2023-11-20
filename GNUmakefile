TEST?=$$(go list ./... |grep -v 'vendor')
default: build

lint:
	golangci-lint run

lintWithFix:
	golangci-lint run --fix

build: lint
	install 

install:
	go install .

submodules:
	@git submodule sync
	@git submodule update --init --recursive
	@git config core.hooksPath githooks
	@git config submodule.recurse true
