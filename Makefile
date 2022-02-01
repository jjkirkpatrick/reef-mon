BINARY=bin/reef-pi
VERSION:=$(shell git describe --always --tags)

.PHONY:go
go:
	go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml

.PHONY:pi
pi:
	env GOOS=linux GOARCH=arm go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml

.PHONY: pi-zero
pi-zero:
	env GOARM=6 GOOS=linux GOARCH=arm go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml


.PHONY: run-dev
run-dev:
	go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml
	./$(BINARY)