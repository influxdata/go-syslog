package rfc5425

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"unicode/utf8"
)

// eof represents a marker rune for the end of the reader
var eof = rune(0)

// isDigit returns true if the rune is in [0,9]
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// isNonZeroDigit returns true if the rune is in ]0,9]
func isNonZeroDigit(ch rune) bool {
	return (ch >= '1' && ch <= '9')
}

// isWhitespace returns true if the rune is a space
func isWhitespace(ch rune) bool {
	return ch == ' '
}

// Scanner represents a lexical scanner
type Scanner struct {
	r      *bufio.Reader
	msglen uint64
	ready  bool
}

// NewScanner returns a new instance of Scanner
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

// read reads the next rune from the buffered reader
// it returns the rune(0) if an error occurs (or io.EOF is returned)
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// Scan returns the next token
func (s *Scanner) Scan() (tok Token) {
	// Read the next rune.
	r := s.read()

	if isNonZeroDigit(r) {
		s.unread()
		s.ready = false
		return s.scanMsgLen()
	}

	// Otherwise read the individual character
	switch r {
	case eof:
		s.ready = false
		return Token{
			typ: EOF,
		}
	case ' ':
		s.ready = true
		return Token{
			typ: WS,
			lit: " ",
		}
	default:
		if s.msglen > 0 && s.ready {
			s.unread()
			return s.scanSyslogMsg()
		}
		s.ready = false
	}

	// (todo) > verify it is reachable
	return Token{
		typ: ILLEGAL,
		lit: string(r),
	}
}

func (s *Scanner) scanMsgLen() Token {
	// Create a buffer and read the current character into it
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent digit character into the buffer
	// Non-digit characters and EOF will cause the loop to exit
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	msglen := buf.String()
	s.msglen, _ = strconv.ParseUint(msglen, 10, 64)

	return Token{
		typ: MSGLEN,
		lit: msglen,
	}
}

func (s *Scanner) scanSyslogMsg() Token {
	// Create a buffer and read the current character into it
	buf := make([]rune, 0, s.msglen)

	for i := uint64(0); i < s.msglen; {
		ch := s.read()

		if ch == eof {
			return Token{
				typ: EOF,
				lit: string(buf),
			}
		}

		buf = append(buf, ch)
		i += uint64(utf8.RuneLen(ch))
	}

	return Token{
		typ: SYSLOGMSG,
		lit: string(buf),
	}
}
