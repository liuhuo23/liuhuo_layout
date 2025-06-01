GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

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

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest



.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: run
# run
run:
	go run -ldflags "-X main.Version=$(VERSION)" ./cmd/{{cookiecutter.project_name}}

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy
.PHONY: client
# generate client code
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


.PHONY: template
template:
	@if [ ! -d "api/$(name)/v1" ]; then \
		mkdir -p api/$(name)/v1; \
		sed 's/#{name}/$(name)/g' api.proto.template > api/$(name)/v1/$(name).proto; \
	else \
		echo "Directory api/$(name)/v1 already exists, skip creation"; \
	fi


.PHONY: config
# generate internal proto
config:
	@echo "Generating internal code..."
	if ls internal/config/*.pb.go >/dev/null 2>&1; then \
	  rm -f internal/config/*.pb.go; \
	fi

	protoc --proto_path=./internal \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./internal \
		$(INTERNAL_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

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
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
