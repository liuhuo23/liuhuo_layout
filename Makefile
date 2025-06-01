.PHONY: help config
.DEFAULT_GOAL := help


ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif


help:
	@echo "Available targets:"
	@echo "  help     - Show this help message (default)"
	@echo "  config   - Generate protobuf code"
	@echo "  wire     - Generate wire code"


.PHONY: config
config:
	@echo "Configuring project..."
	if [ -f ./internal/conf/config.pb.go ]; then cd ./internal/conf && rm -f config.pb.go; fi
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/conf/config.proto
	go mod tidy

.PHONY:init
init:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: wire
wire:
	@echo "Generating wire code..."
	cd ./cmd/liuos && wire


.PHONY: run
run:
	@echo "Running project..."
	go run ./cmd/liuos

.PHONY: build
build:
	@echo "Building project..."
	wire gen ./cmd/liuos && go build -o ./bin ./...  


.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	      --proto_path=./third_party \
 	      --go_out=paths=source_relative:./api \
 	      --go-http_out=paths=source_relative:./api \
	      --openapi_out=fq_schema_naming=true,default_response=false:. \
	      $(API_PROTO_FILES)

.PHONY: client
# generate client code
#  如果生生成 grpc 客户端代码，需要执行以下命令：
#  则在下面添加 --go-grpc_out=paths=source_relative:./api
client:
	@echo "Generating client code..."
	if ls api/${name}/v1/*.go >/dev/null 2>&1; then \
	  rm -f api/${name}/v1/*.go; \
	fi
	protoc --proto_path=./api \
	      --proto_path=./third_party \
	      --go_out=paths=source_relative:./api \
		  --go-http_out=paths=source_relative:./api \
	      api/${name}/v1/*.proto
	go mod tidy



	

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy


.PHONY: all
all:
	make config
	make wire
	make build

