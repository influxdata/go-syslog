package syslog

// WithListener returns a generic option for syslog parsers.
func WithListener(f ParserListener) ParserOption {
	return func(p Parser) Parser {
		p.WithListener(f)
		return p
	}
}
