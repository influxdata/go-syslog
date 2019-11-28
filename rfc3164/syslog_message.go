package rfc3164

import (
	"time"

	"github.com/influxdata/go-syslog/v2/common"
)

type syslogMessage struct {
	prioritySet    bool // We explictly flag the setting of priority since its zero value is a valid priority by RFC 3164
	timestampSet   bool // We explictly flag the setting of timestamp since its zero value is a valid timestamp by RFC 3164
	hasElements    bool
	priority       uint8
	timestamp      time.Time
	hostname       string
	tag            string
	content        string
	msgID          string
	structuredData map[string]map[string]string
	message        string
}

func (sm *syslogMessage) minimal() bool {
	return sm.prioritySet && common.ValidPriority(sm.priority)
}

// export is meant to be called on minimally-valid messages
// thus it presumes priority and version values exists and are correct
func (sm *syslogMessage) export() *SyslogMessage {
	out := &SyslogMessage{}
	out.setPriority(sm.priority)

	if sm.timestampSet {
		out.timestamp = &sm.timestamp
	}

	return out
}

// SyslogMessage represents a RFC3164 syslog message.
type SyslogMessage struct {
	priority       *uint8
	facility       *uint8
	severity       *uint8
	timestamp      *time.Time
	hostname       *string
	appname        *string
	procID         *string
	msgID          *string
	structuredData *map[string]map[string]string
	message        *string
}

func (sm *SyslogMessage) setPriority(value uint8) {
	sm.priority = &value
	facility := uint8(value / 8)
	severity := uint8(value % 8)
	sm.facility = &facility
	sm.severity = &severity
}

// Valid tells whether the receiving SyslogMessage is well-formed or not.
//
// A minimally well-formed RFC3164 syslog message contains at least the priority ([1, 191] or 0).
func (sm *SyslogMessage) Valid() bool {
	// A nil priority or a 0 version means that the message is not valid
	return sm.priority != nil && common.ValidPriority(*sm.priority)
}

// Priority returns the syslog priority or nil when not set
func (sm *SyslogMessage) Priority() *uint8 {
	return sm.priority
}

// Facility returns the facility code.
func (sm *SyslogMessage) Facility() *uint8 {
	return sm.facility
}

// Severity returns the severity code.
func (sm *SyslogMessage) Severity() *uint8 {
	return sm.severity
}

// Timestamp returns the syslog timestamp or nil when not set
func (sm *SyslogMessage) Timestamp() *time.Time {
	return sm.timestamp
}
