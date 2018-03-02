PORT ?= 5001

all: help

help:
	@echo ""
	@echo "  deps       - Installs dependencies"
	@echo "  dev        - Runs development server   PORT ?= $(PORT)"
	@echo "  lint       - Runs linter"
	@echo "  test       - Runs tests"
	@echo ""

deps:
	@dep ensure

dev: deps
	watcher

lint:
	@./node_modules/eslint/bin/eslint.js .

test:
	@BABEL_REACT=true NODE_PATH=./src:./src/app \
		./node_modules/jbs-fe/bin.js test --single-run
