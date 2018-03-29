.PHONY: build graph test

docs/rfc5424_parser.dot: rfc5424/parser.rl
	ragel -Vp rfc5424/parser.rl -o $@

graph: docs/rfc5424_parser.dot

rfc5424/parser.go: rfc5424/parser.rl
	ragel -Z -G2 -o $@ rfc5424/parser.rl

generate: rfc5424/parser.go

test: generate
	go test -v ./...
