package rfc5425

//go:generate stringer -type TokenType $GOFILE

// Token represents a lexical token
type Token struct {
	typ TokenType
	lit string
}

// TokenType represents a lexical token type
type TokenType int

// Tokens
const (
	ILLEGAL TokenType = iota
	EOF
	WS
	MSGLEN
	SYSLOGMSG
)
