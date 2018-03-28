.PHONY: build test
build:
	pigeon syslog.peg > syslog.go

test: build
	go test -v ./...
