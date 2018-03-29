package rfc5424

// Prival represents the priority value
type Prival struct {
	Facility *Facility
	Severity *Severity
	Value    int
}

type PrivalPart interface {
	Message() string
}

type Facility struct {
	Code uint8
}

func NewFacility(value int) *Facility {
	return &Facility{
		Code: uint8(value / 8),
	}
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

func (f *Facility) Message() string {
	return facilities[f.Code]
}

type Severity struct {
	Code uint8
}

func NewSeverity(value int) *Severity {
	return &Severity{
		Code: uint8(value % 8),
	}
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

func (s *Severity) Message() string {
	return severities[s.Code]
}

// NewPrival constructs a complete Prival starting from its value
//
// It assumes values is an utin8 in the range 0..191.
func NewPrival(value int) *Prival {
	return &Prival{
		Facility: NewFacility(value),
		Severity: NewSeverity(value),
		Value:    value,
	}
}
