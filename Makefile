GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: run dependency generate fmt proto wire doc

# 检查kirito脚手架
dependency:
	@which wire &>/dev/null || go get github.com/google/wire/cmd/wire
	@which kirito &>/dev/null || (go get github.com/go-kirito/cmd/kirito && kirito upgrade)

generate: dependency proto wire fmt

fmt:
	@gofmt -s -w .

# 解析protobuf
proto:
	@kirito proto client ./api

# 生成依赖注入文件
wire:
	@go mod tidy
	@kirito wire .
	@go get github.com/google/wire
	@wire ./...

# 生成doc文档
doc:
	@go run doc/main.go

# 运行程序
run:
	@go run cmd/main.go -f config/config.yaml\

all:wire run

# 编译程序
build:
	mkdir -p bin/ && go build -a -mod vendor -tags netgo -ldflags "-X main.Version=$(VERSION)" -o ./bin/app cmd/main.go

docker:
	@go mod tidy
	@go mod vendor
	docker build -t app .
	docker run --rm -p 8000:8000 -p 9000:9000 app

# 生成demo
demo:
	@kirito proto add api/helloworld/v1/helloworld.proto
	@kirito proto client api/helloworld/v1/helloworld.proto
	@mkdir -p internal/helloworld/usecase/
	@kirito proto server api/helloworld/v1/helloworld.proto -t internal/helloworld/usecase/
	@make run

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help