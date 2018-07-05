PACKAGE=sidecar
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/$(PACKAGE) src/main.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/$(PACKAGE) src/main.go

install-deps:
	go get github.com/coreos/etcd/client
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get -u github.com/darrylwest/cassava-logger/logger

format:
	( gofmt -s -w src/*.go src/$(PACKAGE)/*.go test/*/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/*/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

run:
	go run src/main.go --service router

watch:
	go-watcher --loglevel 3

edit:
	make format
	vi -O3 src/*/*.go test/unit/*.go src/main.go

.PHONY: format lint test watch run
.SILENT:
