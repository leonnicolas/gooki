.PHONY: generate build build-all mocks unit

ALL_ARCH := amd64 arm arm64

build:
	CGO_ENABLED=0 go build .

build-all:
	for arch in $(ALL_ARCH); do \
		GOOS=linux GOARCH=$$arch CGO_ENABLED=0 go build -o bin/linux/$$arch/gooki .; \
	done

nuki/swagger.json: nuki/patch.json
	curl -o - https://api.nuki.io/static/swagger/swagger.json | sed 's/"int"/"integer"/' | go run github.com/evanphx/json-patch/cmd/json-patch -p ./nuki/patch.json | jq > $@

generate:
	go generate ./...

generate-all: generate README.md mocks

tmp/help.txt: build
	mkdir -p tmp
	./gooki --help > $@

README.md: tmp/help.txt
	go run github.com/campoy/embedmd -w $@

unit:
	go test ./...

mocks:
	go run github.com/vektra/mockery/v2 --config ./mockery.yaml
