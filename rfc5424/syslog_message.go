package rfc5424

import (
	"time"
)

// SyslogMessage represents a syslog message
type SyslogMessage struct {
	Header
	StructuredData *map[string]map[string]string
	Message        *string
}

func (sm *SyslogMessage) fromMap(data map[string]interface{}) {
	sm.Header.fromMap(data)

	if msg, ok := data["msg"].(string); ok {
		sm.Message = &msg
	}

	if els, ok := data["elements"].(map[string]map[string]string); ok {
		sm.StructuredData = &els
	}
}

// Header represents the header of a syslog message
type Header struct {
	Pri
	Version   uint16
	Timestamp *time.Time
	Hostname  *string
	Appname   *string
	ProcID    *string
	MsgID     *string
}

func (h *Header) fromMap(data map[string]interface{}) {
	if prival, ok := data["prival"].(uint8); !ok {
		panic("prival is a mandatory field")
	} else {
		h.Pri = *NewPri(prival)
	}

	if version, ok := data["version"].(uint16); !ok {
		panic("version is a mandatory field")
	} else {
		h.Version = version
	}

	if timestamp, ok := data["timestamp"].(time.Time); ok {
		h.Timestamp = &timestamp
	}

	if hostname, ok := data["hostname"].(string); ok {
		h.Hostname = &hostname
	}

	if appname, ok := data["appname"].(string); ok {
		h.Appname = &appname
	}

	if procid, ok := data["procid"].(string); ok {
		h.ProcID = &procid
	}

	if msgid, ok := data["msgid"].(string); ok {
		h.MsgID = &msgid
	}
}
