package rfc6587

import (
	"fmt"
	"github.com/influxdata/go-syslog"
	"github.com/influxdata/go-syslog/rfc5424"
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

const start int = 1
const first_final int = 4

const en_main int = 1

type machine struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	err        error

	current []byte

	internal syslog.Machine
	trailer  uint8
}

// NewMachine creates a new FSM able to parse syslog messages transported as per RFC6587.
func NewMachine(options ...syslog.Option) syslog.Machine {
	m := &machine{}

	for _, opt := range options {
		m = opt(m).(*machine)
	}

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

	fmt.Println(m.data)

	m.p = 0
	m.pb = 0
	m.pe = len(m.data)
	m.eof = len(m.data)
	m.err = nil
	m.internal = rfc5424.NewMachine()

	{
		m.cs = start
	}

	{
		var _widec int16
		if (m.p) == (m.pe) {
			goto _test_eof
		}
		switch m.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 4:
			goto st_case_4
		case 3:
			goto st_case_3
		}
		goto st_out
	st_case_1:
		if (m.data)[(m.p)] == 60 {
			goto tr0
		}
		goto st0
	st_case_0:
	st0:
		m.cs = 0
		goto _out
	tr0:

		m.pb = m.p

		goto st2
	tr5:

		m.pb = m.p

		fmt.Println(m.current, string(m.current))
		fmt.Println(m.internal.Parse(m.current, func(x bool) *bool { return &x }(true)))
		fmt.Println("TRAILER")

		goto st2
	st2:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof2
		}
	st_case_2:
		_widec = int16((m.data)[(m.p)])
		switch {
		case (m.data)[(m.p)] < 10:
			if (m.data)[(m.p)] <= 0 {
				_widec = 768 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 2 {
					_widec += 256
				}
			}
		case (m.data)[(m.p)] > 10:
			if 13 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				_widec = 1280 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 1 {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16((m.data)[(m.p)]) - 0)
			if m.trailer == 0 {
				_widec += 256
			}
		}
		switch _widec {
		case 266:
			goto st2
		case 522:
			goto tr3
		case 768:
			goto st2
		case 1024:
			goto tr3
		case 1293:
			goto st2
		case 1549:
			goto st3
		}
		switch {
		case _widec < 11:
			if 1 <= _widec && _widec <= 9 {
				goto st2
			}
		case _widec > 12:
			if 14 <= _widec {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	tr3:

		m.current = m.bytes()

		goto st4
	st4:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof4
		}
	st_case_4:
		_widec = int16((m.data)[(m.p)])
		switch {
		case (m.data)[(m.p)] < 10:
			if (m.data)[(m.p)] <= 0 {
				_widec = 768 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 2 {
					_widec += 256
				}
			}
		case (m.data)[(m.p)] > 10:
			if 13 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				_widec = 1280 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 1 {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16((m.data)[(m.p)]) - 0)
			if m.trailer == 0 {
				_widec += 256
			}
		}
		switch _widec {
		case 60:
			goto tr5
		case 266:
			goto st2
		case 522:
			goto tr3
		case 768:
			goto st2
		case 1024:
			goto tr3
		case 1293:
			goto st2
		case 1549:
			goto st3
		}
		switch {
		case _widec < 11:
			if 1 <= _widec && _widec <= 9 {
				goto st2
			}
		case _widec > 12:
			if 14 <= _widec {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st3:
		if (m.p)++; (m.p) == (m.pe) {
			goto _test_eof3
		}
	st_case_3:
		_widec = int16((m.data)[(m.p)])
		switch {
		case (m.data)[(m.p)] < 10:
			if (m.data)[(m.p)] <= 0 {
				_widec = 768 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 2 {
					_widec += 256
				}
			}
		case (m.data)[(m.p)] > 10:
			if 13 <= (m.data)[(m.p)] && (m.data)[(m.p)] <= 13 {
				_widec = 1280 + (int16((m.data)[(m.p)]) - 0)
				if m.trailer == 1 {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16((m.data)[(m.p)]) - 0)
			if m.trailer == 0 {
				_widec += 256
			}
			if m.trailer == 1 {
				_widec += 512
			}
		}
		switch _widec {
		case 768:
			goto st2
		case 1024:
			goto tr3
		case 1293:
			goto st2
		case 1549:
			goto st3
		case 1802:
			goto st2
		case 2058:
			goto tr3
		case 2314:
			goto tr3
		case 2570:
			goto tr3
		}
		switch {
		case _widec < 11:
			if 1 <= _widec && _widec <= 9 {
				goto st2
			}
		case _widec > 12:
			if 14 <= _widec {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st_out:
	_test_eof2:
		m.cs = 2
		goto _test_eof
	_test_eof4:
		m.cs = 4
		goto _test_eof
	_test_eof3:
		m.cs = 3
		goto _test_eof

	_test_eof:
		{
		}
		if (m.p) == (m.eof) {
			switch m.cs {
			case 4:

				fmt.Println(m.current, string(m.current))
				fmt.Println(m.internal.Parse(m.current, func(x bool) *bool { return &x }(true)))
				fmt.Println("TRAILER")

			}
		}

	_out:
		{
		}
	}

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
