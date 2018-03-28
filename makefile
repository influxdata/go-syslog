.PHONY: build graph test

syslog.dot: syslog.rl
	ragel -Vp syslog.rl -o $@

graph: syslog.dot

syslog.go: syslog.rl
	ragel -Z -G2 -o $@ syslog.rl

build: syslog.go

test: build
	go test -v ./...
