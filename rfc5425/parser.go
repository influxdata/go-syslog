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

// Result represent the resulting syslog message and (eventually) error occured during parsing
type Result struct {
	Message      *rfc5424.SyslogMessage
	MessageError error
	Error        error
}

// ResultHook is a function the user can use to specify what to do with every `Result` instance
type ResultHook func(result *Result)

// Parse parses the incoming bytes accumulating the results
func (p *Parser) Parse() []Result {
	results := []Result{}

	acc := func(result *Result) {
		results = append(results, *result)
	}

	p.ParseExecuting(acc)

	return results
}

// ParseExecuting parses the incoming bytes in the parser executing the `hook` function to each `Result`
//
// It stops parsing when an error regarding RFC 5425 is found.
func (p *Parser) ParseExecuting(hook ResultHook) {
	for {
		var tok Token

		// First token MUST be a MSGLEN
		if tok = p.scan(); tok.typ != MSGLEN {
			hook(&Result{
				Error: fmt.Errorf("found %s, expecting a %s", tok, MSGLEN),
			})
			break
		}

		// Next we MUST see a WS
		if tok = p.scan(); tok.typ != WS {
			hook(&Result{
				Error: fmt.Errorf("found %s, expecting a %s", tok, WS),
			})
			break
		}

		// Next we MUST see a SYSLOGMSG with length equal to MSGLEN
		if tok = p.scan(); tok.typ != SYSLOGMSG {
			e := fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, tok, tok.lit, SYSLOGMSG, p.s.msglen)

			if len(tok.lit) < int(p.s.msglen) && p.bestEffort {
				// Though MSGLEN was not respected, we try to parse the existing SYSLOGMSG as a RFC5424 syslog message
				result := p.parse(tok.lit)
				result.Error = e
				hook(result)
				break
			}

			hook(&Result{
				Error: e,
			})
			break
		}

		// Parse the SYSLOGMSG literal pretending it is a RFC5424 syslog message
		hook(p.parse(tok.lit))

		// Next we MUST see an EOF otherwise the parsing we'll start again
		if tok = p.scan(); tok.typ == EOF {
			break
		} else {
			p.unscan()
		}
	}
}

func (p *Parser) parse(input []byte) *Result {
	sys, err := p.p.Parse(input, &p.bestEffort)

	return &Result{
		Message:      sys,
		MessageError: err,
	}
}

// scan returns the next token from the underlying scanner;
// if a token has been unscanned then read that instead.
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
