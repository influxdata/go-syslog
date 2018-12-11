package rfc5425

import (
	"fmt"
	"io"

	"github.com/influxdata/go-syslog"
	"github.com/influxdata/go-syslog/rfc5424"
)

// Parser is capable to parse the input stream following RFC5425.
//
// Use NewParser function to instantiate one.
type Parser struct {
	msglen     int64
	s          Scanner
	internal   syslog.Machine
	last       Token
	stepback   bool // Wheter to retrieve the last token or not
	bestEffort bool // Best effort mode flag
	emit       syslog.ParserListener
}

// NewParser returns a pointer to a new instance of Parser.
func NewParser(opts ...syslog.ParserOption) syslog.Parser {
	p := &Parser{
		emit: func(*syslog.Result) { /* noop */ },
	}

	for _, opt := range opts {
		p = opt(p).(*Parser)
	}

	if p.internal == nil {
		p.internal = rfc5424.NewParser()
	}

	return p
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
func (p *Parser) HasBestEffort() bool {
	return p.bestEffort
}

// WithBestEffort sets the best effort mode on.
//
// When active the parser tries to recover as much of the syslog messages as possible.
func WithBestEffort() syslog.ParserOption {
	return func(p syslog.Parser) syslog.Parser {
		var parser = p.(*Parser)
		parser.bestEffort = true
		// Push down the best effort, too
		parser.internal = rfc5424.NewParser(rfc5424.WithBestEffort())
		return p
	}
}

// WithListener specifies the function to send the results of the parsing.
func WithListener(ln syslog.ParserListener) syslog.ParserOption {
	return func(p syslog.Parser) syslog.Parser {
		p.(*Parser).emit = ln
		return p
	}
}

// Parse parses the io.Reader incoming bytes.
//
// It stops parsing when an error regarding RFC 5425 is found.
func (p *Parser) Parse(r io.Reader) {
	p.s = *NewScanner(r)
	p.run()
}

func (p *Parser) run() {
	for {
		var tok Token

		// First token MUST be a MSGLEN
		if tok = p.scan(); tok.typ != MSGLEN {
			p.emit(&syslog.Result{
				Error: fmt.Errorf("found %s, expecting a %s", tok, MSGLEN),
			})
			break
		}

		// Next we MUST see a WS
		if tok = p.scan(); tok.typ != WS {
			p.emit(&syslog.Result{
				Error: fmt.Errorf("found %s, expecting a %s", tok, WS),
			})
			break
		}

		// Next we MUST see a SYSLOGMSG with length equal to MSGLEN
		if tok = p.scan(); tok.typ != SYSLOGMSG {
			e := fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, tok, tok.lit, SYSLOGMSG, p.s.msglen)
			// Underflow case
			if len(tok.lit) < int(p.s.msglen) && p.bestEffort {
				// Though MSGLEN was not respected, we try to parse the existing SYSLOGMSG as a RFC5424 syslog message
				result := p.parse(tok.lit)
				if result.Error == nil {
					result.Error = e
				}
				p.emit(result)
				break
			}

			p.emit(&syslog.Result{
				Error: e,
			})
			break
		}

		// Parse the SYSLOGMSG literal pretending it is a RFC5424 syslog message
		result := p.parse(tok.lit)
		if p.bestEffort || result.Error == nil {
			p.emit(result)
		}
		if !p.bestEffort && result.Error != nil {
			p.emit(&syslog.Result{Error: result.Error})
			break
		}

		// Next we MUST see an EOF otherwise the parsing we'll start again
		if tok = p.scan(); tok.typ == EOF {
			break
		} else {
			p.unscan()
		}
	}
}

func (p *Parser) parse(input []byte) *syslog.Result {
	sys, err := p.internal.Parse(input)

	return &syslog.Result{
		Message: sys,
		Error:   err,
	}
}

// scan returns the next token from the underlying scanner;
// if a token has been unscanned then read that instead.
func (p *Parser) scan() Token {
	// If we have a token on the buffer, then return it.
	if p.stepback {
		p.stepback = false
		return p.last
	}

	// Otherwise read the next token from the scanner.
	tok := p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.last = tok

	return tok
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.stepback = true
}
