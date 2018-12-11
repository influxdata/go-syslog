package rfc6587

import (
	syslog "github.com/influxdata/go-syslog"
	"github.com/influxdata/go-syslog/rfc5424"
	parser "github.com/leodido/ragel-machinery/parser"
	"io"
)

const rfc6587Start int = 1
const rfc6587Error int = 0

const rfc6587EnMain int = 1

type machine struct {
	trailer   byte
	candidate []byte
	internal  syslog.Machine
	emit      syslog.ParserListener
}

func (m *machine) process() {
	// res, err := m.internal.Parse(m.candidate, func(x bool) *bool { return &x }(true))
	m.emit(&syslog.Result{
		Message: (&rfc5424.SyslogMessage{}).SetMessage("momentarily empty"),
		Error:   nil,
	})
}

// Exec implements the ragel.Parser interface.
func (m *machine) Exec(s *parser.State) (int, int) {
	// Retrieve previously stored parsing variables
	cs, p, pe, eof, data := s.Get()
	// Inline FSM code here

	{
		if p == pe {
			goto _testEof
		}
		switch cs {
		case 1:
			goto stCase1
		case 0:
			goto stCase0
		case 2:
			goto stCase2
		case 3:
			goto stCase3
		}
		goto stOut
	stCase1:
		if data[p] == 60 {
			goto tr0
		}
		goto st0
	stCase0:
	st0:
		cs = 0
		goto _out
	tr0:

		if len(m.candidate) > 0 {
			// fmt.Println("CANDIDATE", m.candidate)
			m.process()
		}
		// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		// fmt.Println("INIT")
		m.candidate = make([]byte, 0)

		goto st2
	st2:
		if p++; p == pe {
			goto _testEof2
		}
	stCase2:
		if data[p] == 10 {
			goto tr3
		}
		goto st2
	tr3:

		// fmt.Println("TRAILER")
		m.candidate = append(m.candidate, data...)

		goto st3
	st3:
		if p++; p == pe {
			goto _testEof3
		}
	stCase3:
		switch data[p] {
		case 10:
			goto tr3
		case 60:
			goto tr0
		}
		goto st2
	stOut:
	_testEof2:
		cs = 2
		goto _testEof
	_testEof3:
		cs = 3
		goto _testEof

	_testEof:
		{
		}
	_out:
		{
		}
	}

	// Update parsing variables
	s.Set(cs, p, pe, eof)
	return p, pe
}

func (m *machine) OnErr() {
	// fmt.Println("OnErr")
}

func (m *machine) OnEOF() {
	// fmt.Println("OnEOF")
}

func (m *machine) OnCompletion() {
	// fmt.Println("OnCompletion")
	if len(m.candidate) > 0 {
		// fmt.Println("CANDIDATE", m.candidate)
		m.process()
	}
}

// todo(leodido) > make trailer byte configurable, e.g. WithTrailer(TrailerType)
// todo(leodido) > make best effort option for internal parsing configurable, e.g. WithBestEffort(bool)

// NewParser ...
func NewParser(options ...syslog.ParserOption) syslog.Parser {
	m := &machine{
		internal: rfc5424.NewMachine(),
		trailer:  10,
		emit:     func(*syslog.Result) { /* noop */ },
	}

	for _, opt := range options {
		m = opt(m).(*machine)
	}

	return m
}

// WithListener ....
func WithListener(f syslog.ParserListener) syslog.ParserOption {
	return func(m syslog.Parser) syslog.Parser {
		machine := m.(*machine)
		machine.emit = f
		return machine
	}
}

// Parse ...
func (m *machine) Parse(reader io.Reader) {
	r := parser.ArbitraryReader(reader, m.trailer)
	parser.New(r, m, parser.WithStart(1)).Parse()
}
