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
    fmt.Println("set_sdid")
    if sm.StructuredData == nil {
        sm.StructuredData = &(map[string]map[string]string{})
    }
    
    id := string(data[pb:p])
    elements := *sm.StructuredData
    if _, ok := elements[id]; !ok {
        elements[id] = map[string]string{}
    }
}

action set_sdpn {
    fmt.Println("set_sdpn")
    // Assuming SD map already exists, contains currentid key (set from outside)
    elements := *sm.StructuredData
    elements[currentid][string(data[pb:p])] = ""
}

action set_sdpv {
    fmt.Println("set_sdpv")
    // Assuming SD map already exists, contains currentid key and currentparamname key (set from outside)
    elements := *sm.StructuredData
    elements[currentid][currentparamname] = string(data[pb:p])
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

# (fixme) > this is temporary for testing purposes only - must use its real rule
sdpv := sdname >mark %set_sdpv;

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
    sdpv
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
    case sdpv:
        return builder_en_sdpv
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

// SetElementID set a structured data id.
//
// When the provided id already exists the operation is discarded.
func (sm *SyslogMessage) SetElementID(value string) *SyslogMessage {
    return sm.set(sdid, value)
}

// SetParameter set a structured data parameter belonging to the given element.
//
// If the element does not exist it creates one with the given element id.
// When a parameter with the given name already exists for the given element the operation is discarded.
func (sm *SyslogMessage) SetParameter(id string, name string, value string) *SyslogMessage {
    // Create an element with the given id (or re-use the existing one)
    sm.set(sdid, id)
    
    // We can create parameter iff the given element id exists
    if sm.StructuredData != nil {
        elements := *sm.StructuredData
        if _, ok := elements[id]; ok {
            currentid = id
            sm.set(sdpn, name)
            // We can assign parameter value iff the given parameter key exists
            if _, ok := elements[id][name]; ok {
                currentparamname = name
                sm.set(sdpv, value)
            }
        }
    }
    
    return sm
}

// SetMessage set the message value.
func (sm *SyslogMessage) SetMessage(value string) *SyslogMessage {
	return sm.set(msg, value)
}

func (sm *SyslogMessage) String() (string, error) {
    if !sm.Valid() {
        return "", fmt.Errorf("invalid syslog")
    }

    return "", nil // (todo)
}

var currentid string
var currentparamname string