package rfc5425

import (
	"fmt"
	"io"

	"github.com/influxdata/go-syslog/rfc5424"
)

// Parser represents a parser
type Parser struct {
	s          Scanner
	p          rfc5424.Parser
	bestEffort bool
	msglen     int64
	buf        struct {
		tok Token // last read token
		num int   // size (max = 1)
	}
}

// NewParser returns a new instance of Parser
func NewParser(r io.Reader, bestEffort bool) *Parser {
	return &Parser{
		s:          *NewScanner(r),
		p:          *rfc5424.NewParser(),
		bestEffort: bestEffort,
	}
}

// Parse parses ...
func (p *Parser) Parse() ([]rfc5424.SyslogMessage, error) {
	for {
		var tok Token

		// First token MUST be a MSGLEN
		if tok = p.scan(); tok.typ != MSGLEN {
			return nil, fmt.Errorf("found %s, expecting a %s", tok, MSGLEN)
		}
		fmt.Println(tok)

		// Next we MUST see a WS
		if tok = p.scan(); tok.typ != WS {
			return nil, fmt.Errorf("found %s, expecting a %s", tok, WS)
		}
		fmt.Println(tok)

		// Next we MUST see a SYSLOG with length equal to MSGLEN
		if tok = p.scan(); tok.typ != SYSLOGMSG {
			// (todo) > Try to parse syslogmsg literal also here?

			return nil, fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, tok, tok.lit, SYSLOGMSG, p.s.msglen)
		}
		fmt.Println(tok)
		// Parse syslogmsg literal
		sys, err := p.p.Parse(tok.lit, &p.bestEffort)
		fmt.Printf("%#v\n", sys)
		fmt.Println(err)

		// Next we MUST see an EOF otherwise the parsing we'll start again
		if tok = p.scan(); tok.typ == EOF {
			break
		} else {
			p.unscan()
		}

		fmt.Println()
	}

	return nil, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() Token {
	// If we have a token on the buffer, then return it.
	if p.buf.num != 0 {
		p.buf.num = 0
		return p.buf.tok
	}

	// Otherwise read the next token from the scanner.
	tok := p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok = tok

	return tok
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.num = 1
}
