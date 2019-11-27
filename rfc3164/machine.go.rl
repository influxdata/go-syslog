package rfc3164

import (
    "fmt"
    "time"
)

func unsafeUTF8DecimalCodePointsToInt(chars []uint8) int {
    out := 0
    ord := 1
    for i := len(chars) - 1; i >= 0; i-- {
        curchar := int(chars[i])
        out += (curchar - '0') * ord
        ord *= 10
    }
    return out
}

var (
    errPrival         = "expecting a priority value in the range 1-191 or equal to 0 [col %d]"
    errPri            = "expecting a priority value within angle brackets [col %d]"
    errTimestamp      = "expecting a Stamp timestamp [col %d]"
    errHostname       = "expecting an hostname (from 1 to max 255 US-ASCII characters) [col %d]"
    errTag            = "expecting an alphanumeric tag (max 32 characters) [col %d]"
    errContentStart   = "expecting a content part starting with a non-alphanumeric character [col %d]"
    errContent        = "expecting a content part composed by visible characters only [col %d]"
    errParse          = "parsing error [col %d]"
)

%%{
machine rfc3164;

include common "common.rl";

# unsigned alphabet
alphtype uint8;

action mark {
    m.pb = m.p
}

action set_prival {
    output.priority = uint8(unsafeUTF8DecimalCodePointsToInt(m.text()))
    output.prioritySet = true
}

action set_timestamp {
    if t, e := time.Parse(time.Stamp, string(m.text())); e != nil {
        m.err = fmt.Errorf("%s [col %d]", e, m.p)
        fhold;
        fgoto fail;
    } else {
        output.timestamp = t
        output.timestampSet = true
    }
}

action set_hostname {
    output.hostname = string(m.text())
}

action set_tag {
    fmt.Println(">", string(m.text()))
    output.tag += string(m.text())
}

action set_content {
    output.content = string(m.text())
}

action set_message {
    output.message = string(m.text())
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

action err_tag {
    m.err = fmt.Errorf(errTag, m.p)
	fhold;
    fgoto fail;
}

action err_contentstart {
    m.err = fmt.Errorf(errContentStart, m.p)
	fhold;
    fgoto fail;
}

action err_content {
    m.err = fmt.Errorf(errContent, m.p)
	fhold;
    fgoto fail;
}

pri = ('<' prival >mark %from(set_prival) $err(err_prival) '>') @err(err_pri);

timestamp = (datemmm sp datemday sp hhmmss) >mark %set_timestamp @err(err_timestamp);

# (todo) > RFC3164 says something about its maximum length?
hostname = hostnamerange >mark %set_hostname $err(err_hostname);

# Section 4.1.3
tag = alnum{1,32} >mark %set_tag @err(err_tag);

visible = print | 0x80..0xFF;

# The first not alphanumeric character start the content part of the message part
content = !alnum @err(err_contentstart) >mark print* %set_content @err(err_content);

mex = visible+ >mark %set_message;

msg = tag content? mex;

fail := (any - [\n\r])* @err{ fgoto main; };

main := pri timestamp sp hostname sp msg;

}%%

%% write data noerror noprefix;

type machine struct {
    data         []byte
    cs           int
    p, pe, eof   int
    pb           int
    err          error
    bestEffort 	 bool
}

// NewMachine creates a new FSM able to parse RFC3164 syslog messages.
func NewMachine(options ...syslog.MachineOption) syslog.Machine {
	m := &machine{}

	%% access m.;
	%% variable p m.p;
	%% variable pe m.pe;
	%% variable eof m.eof;
	%% variable data m.data;

	return m
}

// WithBestEffort enables best effort mode.
func (m *machine) WithBestEffort() {
	m.bestEffort = true
}

// HasBestEffort tells whether the receiving machine has best effort mode on or off.
func (m *machine) HasBestEffort() bool {
	return m.bestEffort
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

// Parse parses the input byte array as a RFC3164 syslog message.
func (m *machine) Parse(input []byte) (syslog.Message, error) {
    m.data = input
    m.p = 0
    m.pb = 0
    m.pe = len(input)
    m.eof = len(input)
    m.err = nil
    output := &syslogMessage{}

    %% write init;
    %% write exec;

    if m.cs < first_final || m.cs == en_fail {
    	if m.bestEffprt != nil && output.valid() {
    		// An error occurred but partial parsing is on and partial message is minimally valid
    		return output.export(), m.err
    	}
    	return nil, m.err
    }

    return output.export(), nil
}

