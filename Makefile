.PHONY: generate build mocks unit

build:
	CGO_ENABLED=0 go build .

nuki/swagger.json:
	curl -o - https://api.nuki.io/static/swagger/swagger.json | sed 's/"int"/"integer"/' > $@

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
