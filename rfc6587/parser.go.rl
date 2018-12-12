package rfc6587

import (
    "io"
    parser "github.com/leodido/ragel-machinery/parser"
	syslog "github.com/influxdata/go-syslog"
	"github.com/influxdata/go-syslog/rfc5424"
)

%%{
machine rfc6587;

# unsigned alphabet
alphtype uint8;

action on_trailer {
	m.candidate = append(m.candidate, data...)
}

action on_init {
	if len(m.candidate) > 0 {
		m.process()
	}
	m.candidate = make([]byte, 0)
}

t = 10 when { m.trailertyp == LF } |
    00 when { m.trailertyp == NUL };

main := 
	start: (
		'<' >on_init (any)* -> trailer
	),
	trailer: (
		t >on_trailer -> final |
		t >on_trailer -> start 
	);

}%%

%% write data nofinal;

type machine struct{
	trailertyp TrailerType // default is 0 thus TrailerType(LF)
	trailer    byte
	candidate  []byte
	bestEffort bool
	internal   syslog.Machine
	emit       syslog.ParserListener
}

// Exec implements the ragel.Parser interface.
func (m *machine) Exec(s *parser.State) (int, int) {
    // Retrieve previously stored parsing variables
    cs, p, pe, eof, data := s.Get()
    // Inline FSM code here
    %% write exec;
    // Update parsing variables
	s.Set(cs, p, pe, eof)
	return p, pe
}

func (m *machine) OnErr() {
}

func (m *machine) OnEOF() {
}

func (m *machine) OnCompletion() {
	if len(m.candidate) > 0 {
		m.process()		
	}
}

// NewParser returns a syslog.Parser suitable to parse syslog messages sent with non-transparent framing - ie. RFC 6587.
func NewParser(options ...syslog.ParserOption) syslog.Parser {
	m := &machine{
		emit: func(*syslog.Result) { /* noop */ },
	}

	for _, opt := range options {
		m = opt(m).(*machine)
	}

	// No error can happens since during its setting we check the trailer type passed in
	trailer, _ := m.trailertyp.Value()
	m.trailer = byte(trailer)

	if m.internal == nil {
		m.internal = rfc5424.NewMachine()
	}

	return m
}

// HasBestEffort tells whether the receiving parser has best effort mode on or off.
func (m *machine) HasBestEffort() bool {
	return m.bestEffort
}

// WithTrailer ... todo(leodido)
func WithTrailer(t TrailerType) syslog.ParserOption {
	return func(m syslog.Parser) syslog.Parser {
		if val, err := t.Value(); err == nil {
			m.(*machine).trailer = byte(val)
			m.(*machine).trailertyp = t
		}
		
		return m
	}
}

// WithBestEffort sets the best effort mode on.
//
// When active the parser tries to recover as much of the syslog messages as possible.
func WithBestEffort(f syslog.ParserListener) syslog.ParserOption {
	return func(m syslog.Parser) syslog.Parser {
		var p = m.(*machine)
		p.bestEffort = true
		// Push down the best effort, too
		p.internal = rfc5424.NewParser(rfc5424.WithBestEffort())
		return p
	}
}

// WithListener specifies the function to send the results of the parsing.
func WithListener(f syslog.ParserListener) syslog.ParserOption {
	return func(m syslog.Parser) syslog.Parser {
		machine := m.(*machine)
		machine.emit = f
		return machine
	}
}

// Parse parses the io.Reader incoming bytes.
//
// It stops parsing when an error regarding RFC 6587 is found.
func (m *machine) Parse(reader io.Reader) {	
	r := parser.ArbitraryReader(reader, m.trailer)
	parser.New(r, m, parser.WithStart(%%{ write start; }%%)).Parse()
}

func (m *machine) process() {
	lastOne := len(m.candidate) - 1
	if m.candidate[lastOne] == m.trailer {
		m.candidate = m.candidate[:lastOne]
	}
	res, err := m.internal.Parse(m.candidate)
	m.emit(&syslog.Result{
		Message: res,
		Error: err,
	})
}

// todo(leodido) > error management.