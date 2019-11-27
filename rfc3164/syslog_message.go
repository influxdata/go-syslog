package rfc3164

import (
	"time"

	"github.com/influxdata/go-syslog/v2/rfc5424"
)

type syslogMessage struct {
	prioritySet    bool // We explictly flag the setting of priority since its zero value is a valid priority by RFC 3164
	timestampSet   bool // We explictly flag the setting of timestamp since its zero value is a valid timestamp by RFC 3164
	hasElements    bool
	priority       uint8
	version        uint16
	timestamp      time.Time
	hostname       string
	tag            string
	content        string
	msgID          string
	structuredData map[string]map[string]string
	message        string
}

func (sm *syslogMessage) valid() bool {
	if sm.prioritySet {
		return true
	}

	return false
}

func (sm *syslogMessage) export() *rfc5424.SyslogMessage {
	// todo(leodido) > to be implemented
	return &rfc5424.SyslogMessage{}
}
