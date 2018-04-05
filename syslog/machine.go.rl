package syslog

import (
    "errors"
	"time"
    "github.com/davecgh/go-spew/spew"
    "github.com/influxdata/go-syslog/chars"
)

var (
    errNilValue = errors.New("expected nilvalue>")
	errTimestamp = errors.New("expected <timestamp>")
)

%%{
machine rfc5424;

# unsigned alphabet
alphtype uint8;

action mark {
	m.pb = m.p
}

action set_prival {
	m.repository["prival"] = chars.UnsafeUTF8DecimalCodePointsToInt(m.text())
}

action set_version {
	m.repository["version"] = chars.UnsafeUTF8DecimalCodePointsToInt(m.text())
}

action set_timestamp {
	if t, e := time.Parse(time.RFC3339Nano, string(m.text())); e != nil {
        m.err = e
		fhold;
    	fgoto line;
    	fbreak;
    } else {
        m.repository["timestamp"] = t
    }
}

action set_hostname {
	if hostname := string(m.text()); hostname != "-" {
		m.repository["hostname"] = hostname
	}
}

action set_appname {
	if appname := string(m.text()); appname != "-" {
		m.repository["appname"] = appname
	}
}

action set_procid {
	if procid := string(m.text()); procid != "-" {
		m.repository["procid"] = procid
	}
}

action set_msgid {
	if msgid := string(m.text()); msgid != "-" {
		m.repository["msgid"] = msgid
	}
}

action ini_elements {
	m.repository["elements"] = make(map[string]map[string]string)
}

action set_id {
	if elements, ok := m.repository["elements"].(map[string]map[string]string); ok {
		id := string(m.text())
		elements[id] = map[string]string{}
		m.currentelem = id
	}
}

action set_paramname {
	m.currentparam = string(m.text())
}

action set_paramvalue {
	if elements, ok := m.repository["elements"].(map[string]map[string]string); ok {
		elements[m.currentelem][m.currentparam] = string(m.text())
	}
}

action set_msg {
	m.repository["msg"] = string(m.text())
}

action err_timestamp {
	m.err = errTimestamp
	fhold;
    fgoto line;
    fbreak;
}

action err_nilvalue {
    m.err = errNilValue
    fhold;
    fgoto line;
    fbreak;
}


nilvalue = '-';

sp = ' ';

nonzerodigit = '1'..'9';

# 0..59
sexagesimal = '0'..'5' . '0'..'9';

printusascii = '!'..'~';

# 1..191 or 0
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

pri = '<'  prival >mark %set_prival '>';

version = (nonzerodigit digit{0,2}) >mark %set_version;

datemday = ('0' . '1'..'9' | '1'..'2' . '0'..'9' | '3' . '0'..'1');

datemonth = ('0' . '1'..'9' | '1' . '0'..'2');

datefullyear = digit{4};

fulldate = datefullyear '-' datemonth '-' datemday;

timehour = ('0'..'1' . '0'..'9' | '2' . '0'..'3');

timeminute = sexagesimal;

timesecond = sexagesimal;

timesecfrac = '.' digit{1,6};

timenumoffset = ('+' | '-') timehour ':' timeminute;

timeoffset = 'Z' | timenumoffset;

partialtime = timehour ':' timeminute ':' timesecond . timesecfrac?;

fulltime = partialtime . timeoffset;

timestamp = nilvalue | (fulldate >mark 'T' fulltime %set_timestamp) $err(err_timestamp); 

hostname = nilvalue | printusascii{1,255} >mark %set_hostname; 

appname = nilvalue | printusascii{1,48} >mark %set_appname;

procid = nilvalue | printusascii{1,128} >mark %set_procid;

msgid = nilvalue | printusascii{1,32} >mark %set_msgid;

header = pri version sp timestamp sp hostname sp appname sp procid sp msgid;

utf8octets = any*; # (todo) > substitute with correct ranges/rules (see below, rfc 3629)

#utf8octets = utf8char*;

#utf8char = utf81 | utf82 | utf83 | utf84;

#utf81 = 0x00..0x7F;

#utf82 = 0xC2..0xDF utf8tail;

#utf83 = 0xE0 0xA0..0xBF utf8tail | 0xE1..0xEC utf8tail{2} | 0xED 0x80..0x9F utf8tail | 0xEE..0xEF utf8tail{2};

#utf84 = 0xF0 0x90..0xBF utf8tail{2} | 0xF1..0xF3 utf8tail{3} | 0xF4 0x80..0x8F utf8tail{2};

#utf8tail = 0x80..0xBF;

sdname = printusascii{1,32} -- ('=' | sp | ']' | '"');

paramvalue = utf8octets >mark %set_paramvalue; # (todo) > characters '"', '\' and ']' must be escaped

paramname = sdname >mark %set_paramname;

sdparam = paramname '=' '"' paramvalue '"';

sdid = sdname >mark %set_id;

sdelement = '[' sdid (sp sdparam)* ']'; 

structureddata = nilvalue | sdelement+ >ini_elements;

bom = 0xEF 0xBB 0xBF;

msgutf8 = bom utf8octets;

msgany = any*;

msg = (msgany | msgutf8) >mark %set_msg;

line := (any - [\n\r])* @{ fgoto main; };

main := header sp structureddata (sp msg)?;

}%%

%% write data;

type machine struct {
	data       		[]byte
	cs         		int
	p, pe, eof 		int
	pb         		int
	err        		error
	repository  	map[string]interface{}
	currentelem		string
	currentparam	string
}

func NewMachine() *machine {
	m := &machine{
		repository: make(map[string]interface{}, 0),
	}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;
	%% write init;

	return m
}

// Err returns the error that occurred on the last call to Parse.
// 
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

// Position returns the current position into the input.
func (m *machine) Position() int {
	return m.p
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

func (m *machine) Parse(input []byte) (bool, error) {
    m.data = input
	m.p = 0
	m.pb = 0
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil

    %% write init;
    %% write exec;

	spew.Dump(m)

    // m.cs == rfc5424_error
    if m.cs < rfc5424_first_final {
        return false, m.err
    }

    return true, nil
}