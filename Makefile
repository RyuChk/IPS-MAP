GO ?= go
GO_TAGS ?= go_json
GO_FILES = $(shell find . -name \*.go)
GO_BUILDFLAG ?= -v trimpath$(if $(GO_TAGS), -tags $(GO_TAGS),)
GO_TESTFLAGS ?= -tags "test$(if $)"
GOTESTFLAGS ?= -tags "test$(if $(GO_TAGS),$(comma)$(GO_TAGS),)" -short -race
GOPRIVATE=*.cie-ips.com
GOINSECURE=*
# HTTP_PROXY=http://go-athens-athens-proxy.go-athens.svc.cluster.local
# HTTPS_PROXY=https://go-proxy.cie-ips.com
GOPROXY=https://go-proxy.cie-ips.com

all: generate build

preparebeta:
	@printf \\e[1m"Create beta env file"\\e[0m\\n
	@$(GO) mod vendor
	@cp -n .env.example .env.beta || true
	@printf \\e[1m"-------------------Finish init beta environment-------------------"\\e[0m\\n

initbeta: preparebeta
	@printf \\e[1m"Start mongodb replicaset instance via docker"\\e[0m\\n
	@chmod +x ./scripts/rs-init.sh
	@docker compose up -d
	@sleep 5
	@docker exec mongo1 /scripts/rs-init.sh
	@printf \\e[1m"Success startup mongodb"\\e[0m\\n

generate: pregenerate
	@printf \\e[1m"Generate"\\e[0m\\n
	@$(GO) generate ./...
	@cd proto && $(GO) generate

pregenerate:
	@printf \\e[1m"Install dependency"\\e[0m\\n
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@$(GO) get github.com/google/wire/cmd/wire@v0.6.0
	@git config --global http.sslverify false
	@git submodule update --remote

test:
	@printf \\e[1m"Run test"\\e[0m\\n
	@ENV=unittest $(GO) test $(GOTESTFLAGS) ./...

build: .bin/map-grpc .bin/user-tracking-grpc

go.sum:
	@printf \\e[1m"go mod tidy"\\e[0m\\n
	@$(GO) mod tidy

.bin/map-grpc: go.mod go.sum $(GO_FILES)
	@printf \\e[1m"Build .bin/map-grpc"\\e[0m\\n
	@cd cmd/map-grpc && $(GO) build -o ../../.bin/map-grpc .

.bin/user-tracking-grpc: go.mod go.sum $(GO_FILES)
	@printf \\e[1m"Build .bin/user-tracking-grpc"\\e[0m\\n
	@cd cmd/user-tracking-grpc && $(GO) build -o ../../.bin/user-tracking-grpc .