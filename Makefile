.SILENT:
PACKAGE=sidecar
export GOPATH:=$(HOME)/.gopath:$(PWD)
VERSION=`cat VERSION`
.PHONY: format lint test watch run

## help: this help file
help:
	@( echo "" && echo "Makefile targets..." && echo "" )
	@( cat Makefile | grep '^##' | sed -e 's/##/ -/' | sort && echo "" )

## build: compile the primary application
build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/$(PACKAGE) src/main.go )

## build-linux: compile the primary application for linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/$(PACKAGE) src/main.go

## install-deps: fetch the go dependencies
install-deps:
	go get github.com/coreos/etcd/client
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get -u github.com/darrylwest/cassava-logger/logger

## format: format the source files
format:
	( gofmt -s -w src/*.go src/*/*.go test/*/*.go )

## lint: lint the source files
lint:
	@( golint src/... && golint test/... )

## test: run the tests
test:
	@( go vet src/*/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

## run: run the app locally
run:
	go run src/main.go

## watch: watch the source files and build/compile with changes
watch:
	go-watcher --loglevel 3

## edit: format then edit the source files
edit:
	make format
	vi -O3 src/*/*.go test/unit/*.go src/main.go

