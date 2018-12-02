package rfc6587

import (
	"github.com/davecgh/go-spew/spew"
	"fmt"
	"github.com/influxdata/go-syslog/rfc5424"
	"github.com/influxdata/go-syslog"
)

type Trailer uint8

const (
	// LF trailer represents the line feed trailer (the default one as per https://tools.ietf.org/html/rfc6587#section-3.4.2).
	LF Trailer = iota
	// Null trailers have been seen sometimes.
	Null
	// NewLine (CR + LF) trailers are possible too.
	NewLine
)

%%{
machine rfc6587;

# unsigned alphabet
alphtype uint8;

action mark {
    m.pb = m.p
}

action message {
	m.current = m.bytes()
}

action trailer {
	fmt.Println(m.current, string(m.current))
	if mex, err := m.internal.Parse(m.current, func(x bool) *bool { return &x }(true)); true {
		spew.Dump(mex)
		fmt.Println(err)
	}
	fmt.Println("TRAILER")
}

action lf {
	fmt.Println("LF")
}

action null {
	fmt.Println("NULL")
}

action crlf {
	fmt.Println("CRLF")
}

action ciao1 {
	fmt.Println("C1", m.current, string(m.current))
}

action ciao2 {
	fmt.Println("C2", m.current, string(m.current))
}

t = 10 when { m.trailer == 0 } |
    00 when { m.trailer == 2 } | 
    (13 . 10) when { m.trailer == 1 };

main := 
	start: (
		'<' >mark (any)* -> trailer
	),
	trailer: (
		t @message %trailer -> final |
		t %trailer -> start 
	);

}%%

%% write data noerror noprefix;

type machine struct {
	data         []byte
	cs           int
	p, pe, eof   int
	pb           int
	err          error

	current 	 []byte

	internal     syslog.Machine
	bestEffort   bool
	trailer      uint8
}

// NewMachine creates a new FSM able to parse syslog messages transported as per RFC6587.
func NewMachine(options ...syslog.Option) syslog.Machine {
	m := &machine{}

	for _, opt := range options {
		m = opt(m).(*machine)
	}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;

	return m
}

// Err returns the error that occurred on the last call to Parse.
//
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

func (m *machine) bytes() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) getData(in []byte) {
	lastOne := in[len(in)-1]
	lastButOne := in[len(in)-2]
	m.data = in
	switch m.trailer {
		case 0:
			if lastOne != 10 {
				m.data = append(m.data, 10)
			}
			break
		case 1:
			if lastButOne != 13 && lastOne != 10 {
				m.data = append(m.data, 13, 10)
			}
			break
		case 2:
			if lastOne != 0 {
				m.data = append(m.data, 0)
			}
			break
	}
}

// Parse parses the input byte array as a RFC5424 syslog message.
//
// When a valid RFC5424 syslog message is given it outputs its structured representation.
// If the parsing detects an error it returns it with the position where the error occurred.
//
// It can also partially parse input messages returning a partially valid structured representation
// and the error that stopped the parsing.
func (m *machine) Parse(input []byte, bestEffort *bool) (syslog.Message, error) {
	m.getData(input)
	m.p = 0
	m.pb = 0
	m.pe = len(m.data)
	m.eof = len(m.data)
	m.err = nil
	m.internal = rfc5424.NewMachine()

    %% write init;
    %% write exec;

	if m.cs < first_final {
		return nil, m.err
	}

	return &rfc5424.SyslogMessage{}, nil
}

// WithTrailer allows the user to specifiy the trailer to use for non-transparent framing.
func WithTrailer(t Trailer) syslog.Option {
	return func(m syslog.Machine) syslog.Machine {
		machine := m.(*machine)
		switch t {
			case Null:
				machine.trailer = 2
				break
			case NewLine:
				machine.trailer = 1
			case LF:
				fallthrough
			default: 
				machine.trailer = 0
		}
		
		return machine
	}
}

// WithBestEffort sets the best effort mode on.
//
// When active the parser tries to recover as much of the syslog messages as possible.
func WithBestEffort() syslog.Option {
	return func(m syslog.Machine) syslog.Machine {
		machine := m.(*machine)
		machine.bestEffort = true
		return machine
	}
}