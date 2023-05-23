# base
PROJECT_NAME = "phoenix"
MAIN_FILE := "cmd/phoenix/main.go"
PKG = "github.com/CloudWeOps/${PROJECT_NAME}"

# build version
BUILD_TAG := ${shell git tag}
BUILD_BRANCH := ${shell git rev-parse --abbrev-ref HEAD}
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"

dep: ## Get the dependencies
	@go mod tidy

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go build -a -o build/${PROJECT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}' -X '${VERSION_PATH}.GIT_TAG=${BUILD_TAG}'" ${MAIN_FILE}

gen: # Generate code
	@protoc -I=. -I=/usr/local/include --go_out=. --go_opt=module=${PKG} pb/*/*.proto
	@protoc-go-inject-tag -input=pb/*/*.pb.go
	@build/phoenix generate enum -p -m pb/*/*.pb.go

