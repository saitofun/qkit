
MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")
TOOLKIT_PKG = ${MODULE_NAME}/gen/cmd/toolkit

install_toolkit:
	@go install "${TOOLKIT_PKG}/..."

format:
	goimports -w -l -local "${MODULE_NAME}" ./

generate: install_toolkit format
	@cd x/misc/clone/internal/main    && go generate ./...
	@cd x/misc/must/internal/main     && go generate ./...
	@cd kit/validator/strfmt/internal && go generate ./...
	@cd conf/mqtt/                    && go generate ./...
	@cd kit/httptransport/httpx       && go generate ./...

test: generate
	@cd testutil && make test
