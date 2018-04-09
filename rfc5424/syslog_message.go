package rfc5424

import (
	"time"
)

// SyslogMessage represents a syslog message
type SyslogMessage struct {
	Priority       uint8
	facility       uint8
	severity       uint8
	Version        uint16
	Timestamp      *time.Time
	Hostname       *string
	Appname        *string
	ProcID         *string
	MsgID          *string
	StructuredData *map[string]map[string]string
	Message        *string
}

// SetPriority set the priority values and the computed facility and severity codes accordingly.
func (sm *SyslogMessage) SetPriority(value uint8) {
	sm.Priority = value
	sm.facility = uint8(value / 8)
	sm.severity = uint8(value % 8)
}

// Facility returns the facility code.
func (sm *SyslogMessage) Facility() uint8 {
	return sm.facility
}

// Severity returns the severity code.
func (sm *SyslogMessage) Severity() uint8 {
	return sm.severity
}

// FacilityMessage returns the text message for the current facility value.
func (sm *SyslogMessage) FacilityMessage() string {
	return facilities[sm.facility]
}

// SeverityMessage returns the text message for the current severity value.
func (sm *SyslogMessage) SeverityMessage() string {
	return severities[sm.severity]
}

var severities = map[uint8]string{
	0: "Emergency: system is unusable",
	1: "Alert: action must be taken immediately",
	2: "Critical: critical conditions",
	3: "Error: error conditions",
	4: "Warning: warning conditions",
	5: "Notice: normal but significant condition",
	6: "Informational: informational messages",
	7: "Debug: debug-level messages",
}

var facilities = map[uint8]string{
	0:  "kernel messages",
	1:  "user-level messages",
	2:  "mail system",
	3:  "system daemons",
	4:  "security/authorization messages",
	5:  "messages generated internally by syslogd",
	6:  "line printer subsystem",
	7:  "network news subsystem",
	8:  "UUCP subsystem",
	9:  "clock daemon",
	10: "security/authorization messages",
	11: "FTP daemon",
	12: "NTP subsystem",
	13: "log audit",
	14: "log alert",
	15: "clock daemon (note 2)",
	16: "local use 0  (local0)",
	17: "local use 1  (local1)",
	18: "local use 2  (local2)",
	19: "local use 3  (local3)",
	20: "local use 4  (local4)",
	21: "local use 5  (local5)",
	22: "local use 6  (local6)",
	23: "local use 7  (local7)",
}
