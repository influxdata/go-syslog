package rfc5424

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeParse(layout, value string) *time.Time {
	t, _ := time.Parse(layout, value)
	return &t
}

func getStringAddress(str string) *string {
	return &str
}

func createName(x []byte) string {
	str := string(x)
	lim := 60
	if len(str) > lim {
		return string(str[0:lim])
	}
	return str
}

// []byte("<22>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [id1 a=\"b\" c=\"d\" e=\"\"][id2 z=\"w\"]")

type testCase struct {
	input       []byte
	valid       bool
	value       *SyslogMessage
	errorString string
}

var testCases = []testCase{
	// Malformed pri
	{
		[]byte("(190>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
	},
	// (fixme) > Malformed pri outputs wrong error (inner error regarding prival, not the outer one regarding pri)
	// {
	// 	[]byte("<190)122 2018-11-22"),
	// 	false,
	// 	nil,
	// 	"expecting a priority value within angle brackets [col 4]",
	// },
	// Missing pri
	{
		[]byte("122 2018-11-22"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
	},
	// Missing prival
	{
		[]byte("<>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
	},
	// Prival with too much digits
	{
		[]byte("<19000021>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 4]",
	},
	// Prival too high
	{
		[]byte("<192>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 3]",
	},
	// 0 starting prival
	{
		[]byte("<002>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 2]",
	},
	// Non numeric prival
	{
		[]byte("<aaa>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
	},
	// Missing version
	{
		[]byte("<100> 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
	},
	// 0 version
	{
		[]byte("<100>0 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
	},
	// Out of range version
	{
		[]byte("<100>1000 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 8]",
	},
	// Non numeric version
	{
		[]byte("<100>abc 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
	},
	// Wrong year
	{
		[]byte("<101>122 201-11-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 12]",
	},
	// Wrong month
	{
		[]byte("<101>122 2018-112-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 16]",
	},
	// Wrong day
	{
		[]byte("<191>123 2018-02-32"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 18]",
	},
	// Wrong hour
	{
		[]byte("<191>124 2018-02-01:25:15Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 19]",
	},
	// Wrong minutes
	{
		[]byte("<191>125 2003-09-29T22:99:16Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 23]",
	},
	// Wrong seconds
	{
		[]byte("<191>126 2003-09-29T22:09:99Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 26]",
	},
	// Wrong sec fraction
	{
		[]byte("<191>127 2003-09-29T22:09:01.000000000009Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 35]",
	},
	{
		[]byte("<191>128 2003-09-29T22:09:01.Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 29]",
	},
	// Wrong time offset
	{
		[]byte("<191>129 2003-09-29T22:09:01A"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 28]",
	},
	{
		[]byte("<191>129 2003-08-24T05:14:15.000003-24:00"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 37]",
	},
	{
		[]byte("<191>129 2003-08-24T05:14:15.000003-60:00"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 36]",
	},
	{
		[]byte("<191>129 2003-08-24T05:14:15.000003-07:61"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 39]",
	},
	// Non existing date
	// (fixme) > removing last space we get nil, not the parse timestamp error
	{
		[]byte("<34>11 2003-09-31T22:14:15.003Z "),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:15.003Z\": day out of range",
	},
	{
		[]byte("<34>12 2003-09-31T22:14:16Z "),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:16Z\": day out of range",
	},
	// Too long hostname
	{
		[]byte("<1>1 - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 262]",
	},
	// Too long appname
	{
		[]byte("<1>1 - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 57]",
	},
	// Too long procid
	{
		[]byte("<1>1 - - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabX - -"),
		false,
		nil,
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 139]",
	},
	// Too long msgid
	{
		[]byte("<1>1 - - - - abcdefghilmnopqrstuvzabcdefghilmX -"),
		false,
		nil,
		"expecting a msgid (from 1 to max 32 US-ASCII characters) [col 45]",
	},
	// Not print US-ASCII chars for hostname, appname, procid, and msgid
	{
		[]byte("<1>1 -   - - - -"),
		false,
		nil,
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 7]",
	},
	{
		[]byte("<1>1 - -   - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 9]",
	},
	{
		[]byte("<1>1 - - -   - -"),
		false,
		nil,
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 11]",
	},
	{
		[]byte("<1>1 - - - -   -"),
		false,
		nil,
		"expecting a msgid (from 1 to max 32 US-ASCII characters) [col 13]",
	},
	// Invalid, with empty structured data
	{
		[]byte("<1>1 - - - - - []"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
	},
	// Invalid, with structured data id containing space
	{
		[]byte("<1>1 - - - - - [ ]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
	},
	// Invalid, with structured data id containing =
	{
		[]byte("<1>1 - - - - - [=]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
	},
	// Invalid, with structured data id containing ]
	{
		[]byte("<1>1 - - - - - []]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
	},
	// Invalid, with structured data id containing "
	{
		[]byte(`<1>1 - - - - - ["]`),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
	},
	// Invalid, too long structured data id
	{
		[]byte(`<1>1 - - - - - [abcdefghilmnopqrstuvzabcdefghilmX]`),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 48]",
	},
	// Invalid, too long structured data param key
	{
		[]byte(`<1>1 - - - - - [id abcdefghilmnopqrstuvzabcdefghilmX=val]`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
	},
	// Valid
	{
		[]byte("<1>1 - - - - - -"),
		true,
		&SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
		"",
	},
	// Valid, w/o structure data, w/0 procid
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       4,
			severity:       2,
			Priority:       34,
			Version:        1,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:       getStringAddress("mymachine.example.com"),
			Appname:        getStringAddress("su"),
			ProcID:         nil,
			MsgID:          getStringAddress("ID47"),
			StructuredData: nil,
			Message:        getStringAddress("BOM'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
	},
	// Valid, w/o structure data, w/o timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       23,
			severity:       3,
			Priority:       187,
			Version:        222,
			Timestamp:      nil,
			Hostname:       getStringAddress("mymachine.example.com"),
			Appname:        getStringAddress("su"),
			ProcID:         nil,
			MsgID:          getStringAddress("ID47"),
			StructuredData: nil,
			Message:        getStringAddress("'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
	},
	// Valid, w/o structure data, w/o msgid
	{
		[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% Time to make the do-nuts."),
		true,
		&SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        1,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       getStringAddress("192.0.2.1"),
			Appname:        getStringAddress("myproc"),
			ProcID:         getStringAddress("8710"),
			MsgID:          nil,
			StructuredData: nil,
			Message:        getStringAddress("%% Time to make the do-nuts."),
		},
		"",
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, w/o msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - -"),
		true,
		&SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        2,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       nil,
			Appname:        nil,
			ProcID:         nil,
			MsgID:          nil,
			StructuredData: nil,
			Message:        nil,
		},
		"",
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, empty msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - - "),
		true,
		&SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        2,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       nil,
			Appname:        nil,
			ProcID:         nil,
			MsgID:          nil,
			StructuredData: nil,
			Message:        nil,
		},
		"",
	},
	// Valid, with structured data is, w/o structured data params
	{
		[]byte("<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid] some_message"),
		true,
		&SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  getStringAddress("host1"),
			Appname:   getStringAddress("CROND"),
			ProcID:    getStringAddress("10391"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"sdid": map[string]string{},
			},
			Message: getStringAddress("some_message"),
		},
		"",
	},
	// Valid, with structured data is, with structured data params
	{
		[]byte(`<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="⌘"] some_message`),
		true,
		&SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  getStringAddress("host1"),
			Appname:   getStringAddress("CROND"),
			ProcID:    getStringAddress("10391"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"sdid": map[string]string{
					"x": "⌘",
				},
			},
			Message: getStringAddress("some_message"),
		},
		"",
	},
	// Valid, with structured data is, wit structured data params
	{
		[]byte(`<78>2 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="hey \\u2318 hey"] some_message`),
		true,
		&SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  getStringAddress("host1"),
			Appname:   getStringAddress("CROND"),
			ProcID:    getStringAddress("10391"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"sdid": map[string]string{
					"x": `hey \u2318 hey`,
				},
			},
			Message: getStringAddress("some_message"),
		},
		"",
	},
	// Valid, with (escaped) backslash within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="\\valid"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"es": `\valid`,
				},
			},
			Message: getStringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
	},
	// Valid, with control characters within structured data param value
	{
		[]byte("<29>50 2016-01-15T01:00:43Z hn S - - [meta es=\"\b5Ὂg̀9! ℃ᾭG\"] 127.0.0.1 - - 1452819643 \"GET\""),
		true,
		&SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"es": "\b5Ὂg̀9! ℃ᾭG",
				},
			},
			Message: getStringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
	},
	// Valid, with utf8 within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="κόσμε" xx="ñ"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"es": "κόσμε",
					"xx": "ñ",
				},
			},
			Message: getStringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
	},
	// Valid, with structured data, w/o msg
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"),
		true,
		&SyslogMessage{
			facility:  20,
			severity:  5,
			Priority:  165,
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  getStringAddress("example.com"),
			Appname:   getStringAddress("evnts"),
			ProcID:    nil,
			MsgID:     getStringAddress("ID27"),
			StructuredData: &map[string]map[string]string{
				"exampleSDID@32473": map[string]string{
					"iut":         "3",
					"eventSource": "Application",
					"eventID":     "1011",
				},
				"examplePriority@32473": map[string]string{
					"class": "high",
				},
			},
			Message: nil,
		},
		"",
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [id1][id1]"),
		false,
		nil,
		"duplicate structured data element id [col 66]",
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [dupe e=\"1\"][id1][dupe class=\"l\"]"),
		false,
		nil,
		"duplicate structured data element id [col 79]",
	},
	// Valid, with structured data w/o msg
	{
		[]byte("<165>4 2003-10-11T22:14:15.003Z mymachine.it e - 1 [ex@32473 iut=\"3\" eventSource=\"A\"] An application event log entry..."),
		true,
		&SyslogMessage{
			facility:  20,
			severity:  5,
			Priority:  165,
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  getStringAddress("mymachine.it"),
			Appname:   getStringAddress("e"),
			ProcID:    nil,
			MsgID:     getStringAddress("1"),
			StructuredData: &map[string]map[string]string{
				"ex@32473": map[string]string{
					"iut":         "3",
					"eventSource": "A",
				},
			},
			Message: getStringAddress("An application event log entry..."),
		},
		"",
	},
	// Valid, with double quotes in the message
	{
		[]byte(`<29>1 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [origin x-service="svcname"][meta sequenceId="1"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("some-host-name"),
			Appname:   getStringAddress("SEKRETPROGRAM"),
			ProcID:    getStringAddress("prg"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"origin": map[string]string{
					"x-service": "svcname",
				},
				"meta": map[string]string{
					"sequenceId": "1",
				},
			},
			Message: getStringAddress("127.0.0.1 - - 1452819643 \"GET\""),
		},
		"",
	},
	// Valid, with double quotes in the message and escaped character within param
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("some-host-name"),
			Appname:   getStringAddress("SEKRETPROGRAM"),
			ProcID:    getStringAddress("prg"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"escape": "\\]", // (todo) > verify output is correct to the input param value
				},
			},
			Message: getStringAddress("127.0.0.1 - - 1452819643 \"GET\""),
		},
		"",
	},
	// Invalid, param value can not contain closing square bracket - ie., ]
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
	},
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]q"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
	},
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="p]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
	},
	// Invalid, param value can not contain doublequote char - ie., ""
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="a""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 52]",
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape=""b"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
	},
	// Invalid, param value can not contain backslash - ie., \
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 52]",
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="a\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 53]",
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="\n"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
	},
	// Valid, message starting with byte order mark (BOM, \uFEFF)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("\ufeff"),
		},
		"",
	},
	// Valid, greek
	{
		[]byte("<1>1 - - - - - - κόσμε"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("κόσμε"),
		},
		"",
	},
	// Valid, 2 octet sequence
	{
		[]byte("<1>1 - - - - - - "),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress(""),
		},
		"",
	},
	// Valid, spanish (2 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xc3\xb1"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("ñ"),
		},
		"",
	},
	// Valid, colon currency sign (3 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xe2\x82\xa1"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("₡"),
		},
		"",
	},
	// Valid, gothic letter (4 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF \xf0\x90\x8c\xbc"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("\ufeff 𐌼"),
		},
		"",
	},
	// Valid, 5 octet sequence
	{
		[]byte("<1>1 - - - - - - \xC8\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("Ȁ000"),
		},
		"",
	},
	// Valid, 6 octet sequence
	{
		[]byte("<1>1 - - - - - - \xE4\x80\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("䀀000"),
		},
		"",
	},
	// Valid, UTF-8 boundary conditions
	{
		[]byte("<1>1 - - - - - - \xC4\x90\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("Đ000"),
		},
		"",
	},
	{
		[]byte("<1>1 - - - - - - \x0D\x37\x46\x46"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("\r7FF"),
		},
		"",
	},
	// Valid, Tamil poetry of Subramaniya Bharathiyar
	{
		[]byte("<1>1 - - - - - - யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		},
		"",
	},
	// Valid, I Can Eat Glass (Milanese)
	{
		[]byte("<1>1 - - - - - - Sôn bôn de magnà el véder, el me fa minga mal."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("Sôn bôn de magnà el véder, el me fa minga mal."),
		},
		"",
	},
	// Valid, I Can Eat Glass (Romano)
	{
		[]byte("<1>1 - - - - - - Me posso magna' er vetro, e nun me fa male."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("Me posso magna' er vetro, e nun me fa male."),
		},
		"",
	},
	// Valid, I Can Eat Glass (Braille)
	{
		[]byte("<1>1 - - - - - - ⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		},
		"",
	},
	// Valid, I Can Eat Glass (Sanskrit)
	{
		[]byte("<1>1 - - - - - - काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		},
		"",
	},
	// Valid, I Can Eat Glass (Urdu)
	{
		[]byte("<1>1 - - - - - - میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		},
		"",
	},
	// Valid, I Can Eat Glass (Yiddish)
	{
		[]byte("<1>1 - - - - - - איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		},
		"",
	},
	// Valid, I Can Eat Glass (Polish)
	{
		[]byte("<1>1 - - - - - - Mogę jeść szkło, i mi nie szkodzi."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("Mogę jeść szkło, i mi nie szkodzi."),
		},
		"",
	},
	// Valid, I Can Eat Glass (Japanese)
	{
		[]byte("<1>1 - - - - - - 私はガラスを食べられます。それは私を傷つけません。"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("私はガラスを食べられます。それは私を傷つけません。"),
		},
		"",
	},
	// Valid, I Can Eat Glass (Arabic)
	{
		[]byte("<1>1 - - - - - - أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		},
		"",
	},
	// Valid, russian alphabet
	{
		[]byte("<1>1 - - - - - - абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		},
		"",
	},
	// Valid, armenian letters
	{
		[]byte("<1>1 - - - - - - ԰ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ՗՘ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆևֈ։֊֋֌֍֎֏"),
		true,
		&SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  getStringAddress("\u0530ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ\u0557\u0558ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆև\u0588։֊\u058b\u058c֍֎֏"),
		},
		"",
	},
	// Invalid, out of range code within message
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xC1"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
	},
	{
		[]byte("<1>1 - - - - - - \xC1"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
	},
	{
		[]byte("<1>1 - - - - - - \xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
	},
	{
		[]byte("<1>1 - - - - - - \xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	// Invalid, impossible bytes
	{
		[]byte("<1>1 - - - - - - \xfe\xfe\xff\xff"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xfe"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xff"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	// Invalid, overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x80\x80\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x80\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xc0\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	// Invalid, maximum overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x83\xbf\xbf\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x87\xbf\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x8f\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x9f\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xc1\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
	},
	// Invalid, illegal code positions, single utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xad\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xae\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xaf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xb0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbe\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},
	// Invalid, illegal code positions, paired utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80\xed\xb0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
	},

	// (fixme) > evaluate non characters for UTF-8 security concerns, eg. \xef\xbf\xbe
	// (fixme) > "<1>1 -letters- - - - -" gives nil (correct, since it's invalid, missing space) object but also no error
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(createName(tc.input), func(t *testing.T) {
			t.Parallel()

			// (fixme) > instantiating a new machine each time since otherwise there are racing conditions
			msg, err := NewMachine().Parse(tc.input)

			if !tc.valid {
				assert.Nil(t, msg)
				assert.Error(t, err)
				assert.EqualError(t, err, tc.errorString)
			}
			if tc.valid {
				assert.Nil(t, err)
				assert.NotEmpty(t, msg)
			}
			assert.Equal(t, tc.value, msg)
		})
	}
}

// This is here to avoid compiler optimizations that
// could remove the actual call we are benchmarking
// during benchmarks
var benchParseResult *SyslogMessage

func BenchmarkParse(b *testing.B) {
	parser := NewMachine()
	for _, tc := range testCases {
		tc := tc
		b.Run(createName(tc.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchParseResult, _ = parser.Parse(tc.input)
			}
		})
	}
}
