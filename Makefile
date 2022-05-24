NAME=GameOfLife

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

#.SILENT:

define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## dependencies - fetch all dependencies for scripts
dependencies:
	${call colored,dependensies is running...}
	./scripts/get-dependencies.sh

## lint project
lint:
	${call colored,lint is running...}
	./scripts/linters.sh
.PHONY: lint

## Test all packages.
test:
	${call colored,test is running...}
	./scripts/tests.sh
.PHONY: test

## Test coverage.
test-cover:
	${call colored,test-cover is running...}
	go test -race -coverpkg=./... -v -coverprofile .testCoverage.out ./...
	gocov convert .testCoverage.out | gocov report
.PHONY: test-cover

## ------------------------------------------------- Common commands: --------------------------------------------------
## Formats the code.
format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...

## Fix-imports order.
fix-imports:
	${call colored,fixing imports...}
	./scripts/fix-imports-order.sh

