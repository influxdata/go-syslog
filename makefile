SHELL := /bin/bash

syslog/machine.go: syslog/machine.go.rl
	ragel -Z -G2 -o $@ $<

.PHONY: build
build: syslog/machine.go

.PHONY: bench
bench: syslog/*_test.go syslog/machine.go
	go test -bench=. -benchmem -benchtime=10s ./...

.PHONY: tests
tests: syslog/*_test.go syslog/machine.go
	go test -v ./... # (todo) > test race conditions

.PHONY: graph
graph: syslog/machine.go.rl
	ragel -Z -Vp $< -o docs/rfc5424_parser.dot

.PHONY: clean
clean: syslog/machine.go
	rm -f $?