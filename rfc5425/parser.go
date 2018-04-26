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

type Result struct {
	Message *rfc5424.SyslogMessage
	Error   error
}

// Parse parses the bytes coming in the parser emitting rfc5424.SyslogMessage instances
func (p *Parser) Parse() chan Result {
	c := make(chan Result)

	go func() {
		defer close(c)
		for {
			var tok Token

			// First token MUST be a MSGLEN
			if tok = p.scan(); tok.typ != MSGLEN {
				c <- Result{
					Error: fmt.Errorf("found %s, expecting a %s", tok, MSGLEN),
				}
				break
			}
			fmt.Println(tok)

			// Next we MUST see a WS
			if tok = p.scan(); tok.typ != WS {
				c <- Result{
					Error: fmt.Errorf("found %s, expecting a %s", tok, WS),
				}
				break
			}
			fmt.Println(tok)

			// Next we MUST see a SYSLOGMSG with length equal to MSGLEN
			if tok = p.scan(); tok.typ != SYSLOGMSG {
				if len(tok.lit) < int(p.s.msglen) && p.bestEffort {
					// (todo)
					// underflow case
					// try besteffort
				}
				c <- Result{
					Error: fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, tok, tok.lit, SYSLOGMSG, p.s.msglen),
				}
				break
			}
			fmt.Println(tok)

			// Parse the SYSLOGMSG literal pretending it is a RFC5424 syslog message
			sys, err := p.p.Parse(tok.lit, &p.bestEffort)
			c <- Result{
				Message: sys,
				Error:   err,
			}

			// Next we MUST see an EOF otherwise the parsing we'll start again
			if tok = p.scan(); tok.typ == EOF {
				break
			} else {
				p.unscan()
			}

			fmt.Println()
		}
	}()

	return c
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
