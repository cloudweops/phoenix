# base
PROJECT_NAME = "phoenix"
MAIN_FILE := "cmd/phoenix/main.go"
PKG = "github.com/CloudWeOps/${PROJECT_NAME}"

# build version
BUILD_BRANCH := ${shell git rev-parse --abbrev-ref HEAD}
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"

dep: ## Get the dependencies
	@go mod tidy

build: dep ## Build the binary file
	@go build -a -o build/${PROJECT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MAIN_FILE}

gen: # Generate code
	@protoc -I=. -I=/usr/local/include --go_out=. --go_opt=module=${PKG} pb/*/*.proto
	@protoc-go-inject-tag -input=pb/*/*.pb.go
	@build/phoenix generate enum -p -m pb/*/*.pb.go

