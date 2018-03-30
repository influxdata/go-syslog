package rfc5424

import "time"

// SyslogMessage represents a syslog message
type SyslogMessage struct {
	Header
	// StructuredData
	// Message
}

// Header represents the header of a syslog message
type Header struct {
	Pri
	Version
	Timestamp *time.Time
	// Hostname
}

// type StructuredData struct {
// }
