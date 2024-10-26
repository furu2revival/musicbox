BUF_VERSION=v1.42.0
SQL_BOILER_VERSION=v4.16.2
GO_LDFLAGS := -s -w -X github.com/furu2revival/musicbox/app/core/build_info.serverVersion=$(shell git describe --tags --always)

.PHONY: install-tools
install-tools:
	go install ./cmd/protoc-gen-musicbox-server
	go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/arch-go/arch-go@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/volatiletech/sqlboiler/v4@${SQL_BOILER_VERSION}
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

.PHONY: format
format:
	go fmt ./...
	golangci-lint run --issues-exit-code=0 --fix ./...
	buf format --write

BREAKING_CHANGE_BASE_BRANCH?=main
.PHONY: lint
lint:
	golangci-lint run --issues-exit-code=1 ./...
	arch-go
	buf lint

.PHONY: codegen
codegen:
	find . -type f \( -name 'wire_gen.go' \) -delete
	wire ./...
	find . -type f \( -name '*.connect.go' -or -name '*.pb.go' -or -name '*.musicbox.go' \) -delete
	buf generate
	sqlboiler psql --wipe --templates=templates/sqlboiler,$(shell go env GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@${SQL_BOILER_VERSION}/templates/main

.PHONY: codegen-apiclient
codegen-apiclient:
	cd ../pureFrontend && \
		find . -type f \( -name '*_connect.ts' -or -name '*_pb.ts' \) -delete && \
		npx buf generate ../backend/schema/protobuf --template buf.gen.yaml && \
		npm run format

.PHONY: test
test:
	mkdir -p tmp/coverage
	MUSICBOX_CONFIG_FILEPATH=$(shell pwd)/config/default.json TZ=UTC go test -p=1 -coverpkg=./... -coverprofile=tmp/coverage/cover.out ./...
	go tool cover -html=tmp/coverage/cover.out -o tmp/coverage/cover.html

.PHONY: db-migrate
db-migrate:
	docker-compose run --rm --build db-migrate

.PHONY: db-clean
db-clean:
	docker-compose down -v postgres
	docker-compose up -d postgres

.PHONY: build
build: build-api-server

.PHONE: build-api-server
build-api-server:
	CGO_ENABLED=0 go build -ldflags="$(GO_LDFLAGS)" -o tmp/build/api_server ./entrypoint/api_server

MUSICBOX_CONFIG_FILEPATH?=$(shell pwd)/config/default.json
.PHONY: run-api-server
run-api-server:
	MUSICBOX_CONFIG_FILEPATH=$(MUSICBOX_CONFIG_FILEPATH) go run -ldflags="$(GO_LDFLAGS)" ./entrypoint/api_server
