DEPLOYBINARY=/usr/local/bin/reefmon/
DEVVINARY=bin/
VERSION:=$(shell git describe --always --tags)

.PHONY:pi
pi:
	env GOOS=linux GOARCH=arm go build -o "$(DEVVINARY)reef-mon" -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	cp config/config.yml bin/config.yml

.PHONY: run-dev
run-dev:
	go build -o "$(DEVVINARY)reef-mon" -ldflags "-s -w -X main.Version=$(VERSION)"  ./commands
	sudo cp config/config.yml bin/config.yml
	./"$(DEVVINARY)/reef-mon"

.PHONY: deploy-service
deploy-service:
	sudo cp config/config.yml "$(DEPLOYBINARY)config.yml"
	sudo cp config/reef-mon.service /etc/systemd/system/reef-mon.service
	sudo systemctl daemon-reload
	sudo systemctl enable reef-mon
	sudo systemctl start reef-mon
