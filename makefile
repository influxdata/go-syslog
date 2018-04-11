SHELL := /bin/bash

rfc5424/machine.go: rfc5424/machine.go.rl
	ragel -Z -G2 -o $@ $<

.PHONY: build
build: rfc5424/machine.go

.PHONY: bench
bench: rfc5424/*_test.go rfc5424/machine.go
	go test -bench=. -benchmem -benchtime=10s ./...

.PHONY: tests
tests: rfc5424/*_test.go rfc5424/machine.go
	go test -race -coverprofile cover.out -v ./...

docs/rfc5424_parser.dot: rfc5424/machine.go.rl
	ragel -Z -Vp $< -o $@

.PHONY: graph
graph: docs/rfc5424_parser.dot

.PHONY: clean
clean: rfc5424/machine.go
	rm -f $?
