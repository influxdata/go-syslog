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
	// fmt.Println("TRAILER")
	m.candidate = append(m.candidate, data...)
}

action on_init {
	if len(m.candidate) > 0 {
		// fmt.Println("CANDIDATE", m.candidate)
		m.process()
	}
	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// fmt.Println("INIT")
	m.candidate = make([]byte, 0)
}

t = 10; # todo(leodido) > handle other possible trailers with semantic conditioning

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
	trailer   byte
	candidate []byte
	internal  syslog.Machine
	emitter   func(syslog.Result)
}

func (m *machine) process() {
	// res, err := m.internal.Parse(m.candidate, func(x bool) *bool { return &x }(true))
	m.emitter(syslog.Result{
		Message: (&rfc5424.SyslogMessage{}).SetMessage("momentarily empty"),
		Error: nil,
	})
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

// New ...
func New(options ...syslog.ParserOption) syslog.Parser {
	m := &machine{
		internal: rfc5424.NewMachine(),
		trailer: 10,
		emitter: func(syslog.Result) { /* noop */ },
	}

	for _, opt := range options {
		m = opt(m).(*machine)
	}

	return m
}

// WithListener ....
func WithListener(f func(syslog.Result)) syslog.ParserOption {
	return func(m syslog.Parser) syslog.Parser {
		machine := m.(*machine)
		machine.emitter = f
		return machine
	}
}

// Parse ...
func (m *machine) Parse(reader io.Reader) {
	r := parser.ArbitraryReader(reader, m.trailer)
	parser.New(r, m, parser.WithStart(%%{ write start; }%%)).Parse()
}
