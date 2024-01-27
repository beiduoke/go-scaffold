GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

# 指定日期-默认当前日期
DATE:=$(shell date "+%Y%m%d")
# 指定时间-默认当前时间
DATETIME:=$(shell date "+%Y%m%d%H%M%S")

ifeq ($(GOHOSTOS), windows-old)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	# Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	# INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	# API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif


.PHONY: mysqldump
mysqldump:
	@echo '-------------------当前时间备份---------------------'
	@echo '--创建当前日期目录--'
	@mkdir -p ./resouces/backup/${DATE}
	@echo '--执行备份命令--'
	docker exec -i mysql57 bash -c 'exec mysqldump -uroot -p"$$MYSQL_ROOT_PASSWORD" --databases go_scaffold' > ./resouces/backup/${DATE}/go_scaffold.sql

.PHONY: mysqlimport
mysqlimport:
	@echo '------------------指定时间恢复---------------------'
	@echo '--获取日期目录--'
	@echo './resouces/backup/${DATE}'
	@echo '--执行恢复命令--'
	docker exec -i mysql57 bash -c 'exec mysql -uroot -p"$$MYSQL_ROOT_PASSWORD"' < ./resouces/backup/${DATE}/go_scaffold.sql

.PHONY: run
# generate internal proto
run:
	cd cmd/server && bee run .

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.23.1

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	buf generate
	
# generate OpenAPI v3 doc
.PHONY: openapi
openapi:
	buf generate --path proto/admin/interface/v1 --template proto/admin/interface/v1/buf.openapi.gen.yaml
	buf generate --path proto/core/service/v1 --template proto/core/service/v1/buf.openapi.gen.yaml

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;
	make mysqldump;

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
