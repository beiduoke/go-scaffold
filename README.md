# Kratos Project Template

## Install Kratos

```shell
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Create a service

```shell
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## Generate other auxiliary files by Makefile

```shell
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Automated Initialization (wire)

```shell
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Locainze i18n [Link](https://github.com/nicksnyder/go-i18n/blob/main/.github/README.zh-Hans.md)

```shell
go get -u github.com/nicksnyder/go-i18n/v2/goi18n 

# generate code i18n
goi18n extract # to update active.en.toml with the new message

```

## Docker

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```
