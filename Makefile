.PHONY: generate build

build:
	CGO_ENABLED=0 go build .

nuki/swagger.json:
	curl -o - https://api.nuki.io/static/swagger/swagger.json | sed 's/"int"/"integer"/' > $@

generate:
	go generate ./...

generate-all: generate README.md

tmp/help.txt: build
	mkdir -p tmp
	./gooki --help > $@

README.md: tmp/help.txt
	go run github.com/campoy/embedmd -w $@
