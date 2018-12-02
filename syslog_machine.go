package syslog

// Machine represent a FSM able to parse syslog messages and return them in an structured way.
type Machine interface {
	Parse(input []byte, bestEffort *bool) (Message, error)
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
	// todo(leodido) > complete
}

// Option represents the type of options setters.
type Option func(m Machine) Machine
