package rfc5424

import (
    "time"
    "fmt"
)

%%{

machine builder;

include rfc5424 "rfc5424.rl";

action mark {
    pb = p
}

action set_timestamp {
    fmt.Println("set_timestamp")
	if t, e := time.Parse(time.RFC3339Nano, string(data[pb:p])); e == nil {
        sm.Timestamp = &t
    }
}

action set_hostname {
    fmt.Println("set_hostname")
    if s := string(data[pb:p]); s != "-" {
        sm.Hostname = &s
    }
}

action set_appname {
    fmt.Println("set_appname")
    if s := string(data[pb:p]); s != "-" {
        sm.Appname = &s
    }
}

action set_procid {
    fmt.Println("set_procid")
    if s := string(data[pb:p]); s != "-" {
        sm.ProcID = &s
    }
}

action set_msgid {
    fmt.Println("set_msgid")
    if s := string(data[pb:p]); s != "-" {
        sm.MsgID = &s
    }
}

action set_sdid {

}

action set_sdpn {
    
}

action set_msg {
    fmt.Println("set_msg")
    if s := string(data[pb:p]); s != "" {
        sm.Message = &s
    }
}

timestamp := (fulldate 'T' fulltime) >mark %set_timestamp;

hostname := hostnamerange >mark %set_hostname;

appname := appnamerange >mark %set_appname;

procid := procidrange >mark %set_procid;

msgid := msgidrange >mark %set_msgid;

sdid := sdname >mark %set_sdid;

sdpn := sdname >mark %set_sdpn;

msg := (bom? utf8octets) >mark %set_msg;

write data noerror nofinal;
}%%

type entrypoint int

const (
	timestamp entrypoint = iota
	hostname
	appname
	procid
	msgid
    sdid
    sdpn
    msg
)

func (e entrypoint) translate() int {
    switch e {
    case timestamp:
        return builder_en_timestamp
    case hostname:
        return builder_en_hostname
    case appname:
        return builder_en_appname
    case procid:
        return builder_en_procid
    case msgid:
        return builder_en_msgid
    case sdid:
        return builder_en_sdid
    case sdpn:
        return builder_en_sdpn
    case msg:
        return builder_en_msg
    default:
        return builder_start
    }
}

func (sm *SyslogMessage) set(from entrypoint, value string) *SyslogMessage {
    data := []byte(value)
	p := 0
	pb := 0
	pe := len(data)
	eof := len(data)
    cs := from.translate()
    %% write exec;

    return sm
}

// SetPriority set the priority value and the computed facility and severity codes accordingly.
//
// It ignores incorrect priority values (range [0, 191]).
func (sm *SyslogMessage) SetPriority(value uint8) *SyslogMessage {
	if value >= 0 && value <= 191 {
		sm.setPriority(value)
	}

	return sm
}

// SetVersion set the version value.
//
// It ignores incorrect version values (range ]0, 999]).
func (sm *SyslogMessage) SetVersion(value uint16) *SyslogMessage {
	if value > 0 && value <= 999 {
		sm.Version = value
	}

	return sm
}

// SetTimestamp set the timestamp value.
func (sm *SyslogMessage) SetTimestamp(value string) *SyslogMessage {
    return sm.set(timestamp, value)
}

// SetHostname set the hostname value.
func (sm *SyslogMessage) SetHostname(value string) *SyslogMessage {
    return sm.set(hostname, value)
}

// SetAppname set the appname value.
func (sm *SyslogMessage) SetAppname(value string) *SyslogMessage {
    return sm.set(appname, value)
}

// SetProcID set the procid value.
func (sm *SyslogMessage) SetProcID(value string) *SyslogMessage {
    return sm.set(procid, value)
}

// SetMsgID set the msgid value.
func (sm *SyslogMessage) SetMsgID(value string) *SyslogMessage {
    return sm.set(msgid, value)
}

// (todo) > setters for structured data elements (id + parameters)
// func (sm *SyslogMessage) SetElementID(value string) *SyslogMessage {
//     return sm.set(sdid, value) // (todo) > ignore incoming duplicates? per essere coerenti col design si ...
// }

// func (sm *SyslogMessage) SetParameter(element string, name string, value string) *SyslogMessage {
//     // (todo) > se chiave non esiste gia sm.set(sdid, element)
//     sm.set(sdpn, name)
//     // (todo) > se id okay sm.set(sdpv, value) 
// }

// SetMessage set the message value.
func (sm *SyslogMessage) SetMessage(value string) *SyslogMessage {
	return sm.set(msg, value)
}

// textmarshaler or string method?