CURRENT_OS := $(shell uname | tr "[A-Z]" "[a-z]")
CURRENT_ARCH := $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')

GO_BUILD := CGO_ENABLED=0 go build -trimpath

.PHONY: build
build: ## Build binary
	$(GO_BUILD) -o ./build/ ./...

.PHONY: build-linux-amd64
build-linux-amd64: ## Build linux amd64 binary
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./build/linux_amd64/ ./...

