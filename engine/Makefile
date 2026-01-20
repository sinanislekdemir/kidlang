# Makefile

.PHONY: default
default: build

install:
	@go install

build:
	@go build

win64:
	@GOOS=windows GOARCH=amd64 go build -o kidlang.exe

debug-test:
	dlv test github.com/sinanislekdemir/kidlang/$(pkg) -- -test.v

test:
	@go test -v ./...

fmt:
	@go fmt ./...
