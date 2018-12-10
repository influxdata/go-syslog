package syslog

import (
	"io"
	"time"
)

// Machine represent a FSM able to parse an entire syslog message and return it in an structured way.
type Machine interface {
	Parse(input []byte) (Message, error)
}

// MachineOption represents the type of options setters.
type MachineOption func(m Machine) Machine

// BestEfforter is an interface that wraps the HasBestEffort method.
type BestEfforter interface {
	HasBestEffort() bool
}

// Parser is an interface that wraps the Parse method.
type Parser interface {
	Parse(r io.Reader)
}

// ParserOption represent an option for Parser instances.
type ParserOption func(p Parser) Parser

// Result wraps the outcomes obtained parsing a syslog message.
type Result struct {
	Message Message
	Error   error
}

// Message represent a structured representation of a syslog message.
type Message interface {
	Valid() bool
	Priority() *uint8
	Version() uint16
	Facility() *uint8
	Severity() *uint8
	FacilityMessage() *string
	FacilityLevel() *string
	SeverityMessage() *string
	SeverityLevel() *string
	SeverityShortLevel() *string
	Timestamp() *time.Time
	Hostname() *string
	ProcID() *string
	Appname() *string
	MsgID() *string
	Message() *string
	StructuredData() *map[string]map[string]string
}
