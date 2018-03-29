package rfc5424

//go:generate ragel -Z -G2 -o parser.go parser.rl

func Parse(data string) Message {
	return parse(data)
}
