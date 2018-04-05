.PHONY: generate graph test clean bench

docs/rfc5424_parser.dot: rfc5424/parser.rl
	ragel -Vp rfc5424/parser.rl -o $@

graph: 
	ragel -Z -Vp rfc5424/parser.rl -o docs/rfc5424_parser.dot

rfc5424/parser.rl: rfc5424/machine.rl

rfc5424/parser.go: rfc5424/parser.rl
	ragel -Z -G2 -o $@ rfc5424/parser.rl

generate: clean rfc5424/parser.go

test: generate
	go test -race -v ./...

bench: generate
	go test -race -bench=. ./...

clean:
	rm -f rfc5424/parser.go

