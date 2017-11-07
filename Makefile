.PHONY: run build test

run:
	@go run $$(ls -1 src/*.go | grep -v _test.go)

build:
	@go build $$(ls -1 src/*.go | grep -v _test.go)

test:
	@cd src && go test
