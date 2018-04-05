package rfc5424

import "time"

// SyslogMessage represents a syslog message
type SyslogMessage struct {
	Header
	StructuredData
	Message string
}

// Header represents the header of a syslog message
type Header struct {
	Pri
	Version
	Timestamp *time.Time
	Hostname  string
	Appname   string
	ProcID    string
	MsgID     string
}

// StructuredData representes the element described at https://tools.ietf.org/html/rfc5424#section-6.3
type StructuredData struct {
	Name   string
	Params map[string]string
}
