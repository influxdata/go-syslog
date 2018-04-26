package rfc5424

import (
	"time"
	"fmt"
)

var (
	errPrival         = "expecting a priority value in the range 1-191 or equal to 0 [col %d]"
	errPri            = "expecting a priority value within angle brackets [col %d]"
	errVersion        = "expecting a version value in the range 1-999 [col %d]"
	errTimestamp      = "expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col %d]"
	errHostname       = "expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col %d]"
	errAppname        = "expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col %d]"
	errProcid         = "expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col %d]"
	errMsgid          = "expecting a msgid (from 1 to max 32 US-ASCII characters) or a nil value [col %d]"
	errStructuredData = "expecting a structured data section containing one or more elements (`[id( key=\"value\")*]+`) or a nil value [col %d]"
	errSdID           = "expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col %d]"
	errSdIDDuplicated = "duplicate structured data element id [col %d]"
	errSdParam        = "expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col %d]"
	errMsg            = "expecting a free-form optional message in UTF-8 (starting with or without BOM) [col %d]"
	errEscape         = "expecting chars `]`, `\"`, and `\\` to be escaped within param value [col %d]"
	errParse          = "parsing error [col %d]"
)

%%{
machine rfc5424;

# unsigned alphabet
alphtype uint8;

action mark {
	m.pb = m.p
}

action markmsg {
	m.msg_at = m.p
}

action set_prival {
	output.SetPriority(uint8(unsafeUTF8DecimalCodePointsToInt(m.text())))
}

action set_version {
	output.Version = uint16(unsafeUTF8DecimalCodePointsToInt(m.text()))
}

action set_timestamp {
	if t, e := time.Parse(time.RFC3339Nano, string(m.text())); e != nil {
        m.err = fmt.Errorf("%s [col %d]", e, m.p)
		fhold;
    	fgoto fail;
    } else {
        output.Timestamp = &t
    }
}

action set_hostname {
	if hostname := string(m.text()); hostname != "-" {
		output.Hostname = &hostname
	}
}

action set_appname {
	if appname := string(m.text()); appname != "-" {
		output.Appname = &appname
	}
}

action set_procid {
	if procid := string(m.text()); procid != "-" {
		output.ProcID = &procid
	}
}

action set_msgid {
	if msgid := string(m.text()); msgid != "-" {
		output.MsgID = &msgid
	}
}


action ini_elements {
	output.StructuredData = &(map[string]map[string]string{})
}

action set_id {
	if elements, ok := interface{}(output.StructuredData).(*map[string]map[string]string); ok {
		id := string(m.text())
		if _, ok := (*elements)[id]; ok {
			// As per RFC5424 section 6.3.2 SD-ID MUST NOT exist more than once in a message
			m.err = fmt.Errorf(errSdIDDuplicated, m.p)
			fhold;
			fgoto fail;
		} else {
			(*elements)[id] = map[string]string{}
			m.currentelem = id
		}
	}
}

action ini_sdparam {
	m.backslash_at = []int{}
}

action add_slash {
	m.backslash_at = append(m.backslash_at, m.p)
}

action set_paramname {
	m.currentparam = string(m.text())
}

action set_paramvalue {
	if elements, ok := interface{}(output.StructuredData).(*map[string]map[string]string); ok {
		// (fixme) > what if SD-PARAM-NAME already exist for the current element (ie., current SD-ID)?

		// Store text
		text := m.text()
		
		// Strip backslashes only when there are ...
		if len(m.backslash_at) > 0 {
			// We need a copy here to not modify m.data
			cp := append([]byte(nil), text...)
			for _, pos := range m.backslash_at {
				at := pos - m.pb
				cp = append(cp[:at], cp[(at + 1):]...)
			}
			
			text = cp
		}
		(*elements)[m.currentelem][m.currentparam] = string(text)
	}
}

action set_msg {
	if msg := string(m.text()); msg != "" {
		output.Message = &msg
	}
}

action err_prival {
	m.err = fmt.Errorf(errPrival, m.p)
	fhold;
    fgoto fail;
}

action err_pri {
	m.err = fmt.Errorf(errPri, m.p)
	fhold;
    fgoto fail;
}

action err_version {
	m.err = fmt.Errorf(errVersion, m.p)
	fhold;
    fgoto fail;
}

action err_timestamp {
	m.err = fmt.Errorf(errTimestamp, m.p)
	fhold;
    fgoto fail;
}

action err_hostname {
	m.err = fmt.Errorf(errHostname, m.p)
	fhold;
    fgoto fail;
}

action err_appname {
	m.err = fmt.Errorf(errAppname, m.p)
	fhold;
    fgoto fail;
}

action err_procid {
	m.err = fmt.Errorf(errProcid, m.p)
	fhold;
    fgoto fail;
}

action err_msgid {
	m.err = fmt.Errorf(errMsgid, m.p)
	fhold;
    fgoto fail;
}

action err_structureddata {
	m.err = fmt.Errorf(errStructuredData, m.p)
	fhold;
    fgoto fail;
}

action err_sdid {
	delete(*output.StructuredData, m.currentelem)
	if len(*output.StructuredData) == 0 {
		output.StructuredData = nil
	}
	m.err = fmt.Errorf(errSdID, m.p)
	fhold;
    fgoto fail;
}

action err_sdparam {
	if elements, ok := interface{}(output.StructuredData).(*map[string]map[string]string); ok {
		delete((*elements)[m.currentelem], m.currentparam)
	}
	m.err = fmt.Errorf(errSdParam, m.p)
	fhold;
    fgoto fail;
}

action err_msg {
	// If error encountered within the message rule ...
	if m.msg_at > 0 {
		// Save the text until valid (m.p is where the parser has stopped)
		if trunc := string(m.data[m.msg_at:m.p]); trunc != "" {
			output.Message = &trunc
		}
	}

	m.err = fmt.Errorf(errMsg, m.p)
	fhold;
    fgoto fail;
}

action err_escape {
	m.err = fmt.Errorf(errEscape, m.p)
	fhold;
    fgoto fail;
}

action err_parse {
	m.err = fmt.Errorf(errParse, m.p)
	fhold;
    fgoto fail;
}

nilvalue = '-';

sp = ' ';

nonzerodigit = '1'..'9';

# 0..59
sexagesimal = '0'..'5' . '0'..'9';

printusascii = '!'..'~';

# 1..191
privalrange = (('1' ('9' ('0'..'1'){,1} | '0'..'8' ('0'..'9'){,1}){,1}) | ('2'..'9' ('0'..'9'){,1}));

# 1..191 or 0
prival = (privalrange | '0') >mark %from(set_prival) $err(err_prival);

pri = ('<' prival '>') @err(err_pri);

version = (nonzerodigit digit{0,2} <err(err_version)) >mark %from(set_version) %eof(set_version) @err(err_version);

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

timestamp = (nilvalue | (fulldate >mark 'T' fulltime %set_timestamp %err(set_timestamp))) <>err(err_timestamp);

hostname = nilvalue | printusascii{1,255} >mark %set_hostname $err(err_hostname);

appname = nilvalue | printusascii{1,48} >mark %set_appname $err(err_appname);

procid = nilvalue | printusascii{1,128} >mark %set_procid $err(err_procid);

msgid = nilvalue | printusascii{1,32} >mark %set_msgid $err(err_msgid);

header = (pri version sp timestamp sp hostname sp appname sp procid sp msgid) <>err(err_parse);

# rfc 3629
utf8tail = 0x80..0xBF;

utf81 = 0x00..0x7F;

utf82 = 0xC2..0xDF utf8tail;

utf83 = 0xE0 0xA0..0xBF utf8tail | 0xE1..0xEC utf8tail{2} | 0xED 0x80..0x9F utf8tail | 0xEE..0xEF utf8tail{2};

utf84 = 0xF0 0x90..0xBF utf8tail{2} | 0xF1..0xF3 utf8tail{3} | 0xF4 0x80..0x8F utf8tail{2};

utf8char = utf81 | utf82 | utf83 | utf84;

utf8octets = utf8char*;

sdname = (printusascii - ('=' | sp | ']' | '"')){1,32};

# utf8char except ", ], \
utf8charwodelims = utf8char - (0x22 | 0x5D | 0x5C);

# \", \], \\
escapes = (0x5C >add_slash (0x22 | 0x5D | 0x5C)) $err(err_escape);

# As per section 6.3.3 param value MUST NOT contain '"', '\' and ']', unless they are escaped.
# A backslash '\' followed by none of the this three characters is an invalid escape sequence.
# In this case, treat it as a regular backslash and the following character as a regular character (not altering the invalid sequence).
paramvalue = (utf8charwodelims* escapes*)+ >mark %set_paramvalue;

paramname = sdname >mark %set_paramname;

sdparam = (paramname '=' '"' paramvalue '"') >ini_sdparam $err(err_sdparam);

# (note) > finegrained semantics of section 6.3.2 not represented here since not so useful for parsing goal
sdid = sdname >mark %set_id %err(set_id) $err(err_sdid);

sdelement = ('[' sdid (sp sdparam)* ']');

structureddata = nilvalue | sdelement+ >ini_elements $err(err_structureddata);

bom = 0xEF 0xBB 0xBF;

msg = (bom? utf8octets) >mark >markmsg %set_msg $err(err_msg);

fail := (any - [\n\r])* @err{ fgoto main; };  

main := header sp structureddata (sp msg)? $err(err_parse);

}%%

%% write data noerror noprefix;

type machine struct {
	data         []byte
	cs           int
	p, pe, eof   int
	pb           int
	err          error
	currentelem  string
	currentparam string
	msg_at       int
	backslash_at []int
}

// NewMachine creates a new FSM able to parse RFC5424 syslog messages.
func NewMachine() *machine {
	m := &machine{}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;

	return m
}

// Err returns the error that occurred on the last call to Parse.
//
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
	return m.err
}

func (m *machine) text() []byte {
	return m.data[m.pb:m.p]
}

// Parse parses the input byte array as a RFC5424 syslog message.
//
// When a valid RFC5424 syslog message is given it outputs its structured representation.
// If the parsing detects an error it returns it with the position where the error occurred.
//
// It can also partially parse input messages returning a partially valid structured representation
// and the error that stopped the parsing.
func (m *machine) Parse(input []byte, bestEffort *bool) (*SyslogMessage, error) {
	m.data = input
	m.p = 0
	m.pb = 0
	m.msg_at = 0
	m.backslash_at = []int{}
	m.pe = len(input)
	m.eof = len(input)
	m.err = nil
	output := &SyslogMessage{}

    %% write init;
    %% write exec;

	if m.cs < first_final || m.cs == en_fail {
		if bestEffort != nil && *bestEffort != false && output.Valid() {
			// An error occurred but partial parsing is on and partial message is minimally valid
			return output, m.err
		}
		return nil, m.err
	}

	return output, nil
}