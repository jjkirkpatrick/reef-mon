BINARY=/usr/local/reefmon/bin/
VERSION:=$(shell git describe --always --tags)

.PHONY:pi
pi:
	env GOOS=linux GOARCH=arm go build -o "$(BINARY)reef-mon" -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml

.PHONY: run-dev
run-dev:
	go build -o "$(BINARY)reef-mon" -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml
	./"$(BINARY)/reef-mon"

.PHONY: deploy-service
deploy-service:
	sudo env GOOS=linux GOARCH=arm go build -o "$(BINARY)reef-mon" -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	sudo cp config/config.yml "$(BINARY)config.yml"
	sudo cp config/reef-mon.service /etc/systemd/system/reef-mon.service
	sudo systemctl daemon-reload
	sudo systemctl enable reef-mon
	sudo systemctl start reef-mon