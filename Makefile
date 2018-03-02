PORT ?= 5001

all: help

help:
	@echo ""
	@echo "  deps       - Installs dependencies"
	@echo "  dev        - Runs development server   PORT ?= $(PORT)"
	@echo "  test       - Runs tests"
	@echo ""

deps:
	@dep ensure

dev: deps
	source .env && watcher

test:
	go test $$(go list ./... | grep -v /vendor/)
