package rfc5424

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeParse(layout, value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}

func getStringAddress(str string) *string {
	return &str
}

func getUint8Address(x uint8) *uint8 {
	return &x
}

func rxpad(str string, lim int) string {
	str = str + strings.Repeat(" ", lim)
	return str[:lim]
}

type testCase struct {
	input        []byte
	valid        bool
	value        SyslogMessage
	errorString  string
	partialValue SyslogMessage
}

var testCases = []testCase{
	// Invalid, empty input
	{
		[]byte(""),
		false,
		SyslogMessage{},
		"expecting a priority value within angle brackets [col 0]",
		SyslogMessage{},
	},
	// Invalid, multiple syslog messages on multiple lines
	{
		[]byte(`<1>1 - - - - - -
	<2>1 - - - - - -`),
		false,
		SyslogMessage{},
		"parsing error [col 16]",
		SyslogMessage{
			Priority: 1,
			severity: 1,
			facility: 0,
			Version:  1,
		},
	},
	// Invalid, new lines allowed only within message part
	{
		[]byte("<1>1 - \nhostname - - - -"),
		false,
		SyslogMessage{},
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 7]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - host\x0Aname - - - -"),
		false,
		SyslogMessage{},
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 11]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - \nan - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 9]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - a\x0An - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 10]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - \npid - -"),
		false,
		SyslogMessage{},
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 11]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - p\x0Aid - -"),
		false,
		SyslogMessage{},
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 12]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - \nmid -"),
		false,
		SyslogMessage{},
		"expecting a msgid (from 1 to max 32 US-ASCII characters) or a nil value [col 13]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - m\x0Aid -"),
		false,
		SyslogMessage{},
		"expecting a msgid (from 1 to max 32 US-ASCII characters) or a nil value [col 14]",
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
	},
	// Invalid, malformed pri
	{
		[]byte("(190>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value within angle brackets [col 0]",
		SyslogMessage{},
	},
	// (fixme) > Malformed pri outputs wrong error (inner error regarding prival, not the outer one regarding pri)
	// {
	// 	[]byte("<87]123 -"),
	// 	false,
	// 	nil,
	// 	"expecting a priority value within angle brackets [col 3]",
	// 	SyslogMessage{}, // nil since cannot reach version
	// },
	// Invalid, missing pri
	{
		[]byte("122 - - - - - -"),
		false,
		SyslogMessage{},
		"expecting a priority value within angle brackets [col 0]",
		SyslogMessage{},
	},
	// Invalid, missing prival
	{
		[]byte("<>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
		SyslogMessage{},
	},
	// Invalid, prival with too much digits
	{
		[]byte("<19000021>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value in the range 1-191 or equal to 0 [col 4]",
		SyslogMessage{}, // no valid partial message since was not able to reach and extract version (which is mandatory for a valid message)
	},
	// Invalid, prival too high
	{
		[]byte("<192>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value in the range 1-191 or equal to 0 [col 3]",
		SyslogMessage{},
	},
	// Invalid, 0 starting prival
	{
		[]byte("<002>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value in the range 1-191 or equal to 0 [col 2]",
		SyslogMessage{},
	},
	// Invalid, non numeric prival
	{
		[]byte("<aaa>122 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
		SyslogMessage{},
	},
	// Invalid, missing version
	{
		[]byte("<100> 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a version value in the range 1-999 [col 5]",
		SyslogMessage{},
	},
	// Invalid, 0 version
	{
		[]byte("<103>0 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a version value in the range 1-999 [col 5]",
		SyslogMessage{},
	},
	// Invalid, out of range version
	{
		[]byte("<101>1000 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a version value in the range 1-999 [col 8]",
		SyslogMessage{},
	},
	// Invalid, non numeric version
	{
		[]byte("<102>abc 2018-11-22"),
		false,
		SyslogMessage{},
		"expecting a version value in the range 1-999 [col 5]",
		SyslogMessage{},
	},
	// Invalid, wrong year
	{
		[]byte("<101>122 201-11-22"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 12]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  122,
		},
	},
	{
		[]byte("<101>189 0-11-22"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 10]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  189,
		},
	},
	// Invalid, wrong month
	{
		[]byte("<101>122 2018-112-22"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 16]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  122,
		},
	},
	// Invalid, wrong day
	{
		[]byte("<101>123 2018-02-32"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 18]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  123,
		},
	},
	// Invalid, wrong hour
	{
		[]byte("<101>124 2018-02-01:25:15Z"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 19]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  124,
		},
	},
	// Invalid, wrong minutes
	{
		[]byte("<101>125 2003-09-29T22:99:16Z"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 23]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  125,
		},
	},
	// Invalid, wrong seconds
	{
		[]byte("<101>126 2003-09-29T22:09:99Z"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 26]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  126,
		},
	},
	// Invalid, wrong sec fraction
	{
		[]byte("<101>127 2003-09-29T22:09:01.000000000009Z"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 35]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  127,
		},
	},
	{
		[]byte("<101>128 2003-09-29T22:09:01.Z"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 29]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  128,
		},
	},
	{
		[]byte("<101>28 2003-09-29T22:09:01."),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 28]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  28,
		},
	},
	// Invalid, wrong time offset
	{
		[]byte("<101>129 2003-09-29T22:09:01A"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 28]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  129,
		},
	},
	{
		[]byte("<101>130 2003-08-24T05:14:15.000003-24:00"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 37]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  130,
		},
	},
	{
		[]byte("<101>131 2003-08-24T05:14:15.000003-60:00"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 36]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  131,
		},
	},
	{
		[]byte("<101>132 2003-08-24T05:14:15.000003-07:61"),
		false,
		SyslogMessage{},
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 39]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  132,
		},
	},
	// Invalid, non existing dates
	{
		[]byte("<101>11 2003-09-31T22:14:15.003Z"),
		false,
		SyslogMessage{},
		"parsing time \"2003-09-31T22:14:15.003Z\": day out of range [col 32]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  11,
		},
	},
	{
		[]byte("<101>12 2003-09-31T22:14:16Z"),
		false,
		SyslogMessage{},
		"parsing time \"2003-09-31T22:14:16Z\": day out of range [col 28]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  12,
		},
	},
	{
		[]byte("<101>12 2018-02-29T22:14:16+01:00"),
		false,
		SyslogMessage{},
		"parsing time \"2018-02-29T22:14:16+01:00\": day out of range [col 33]",
		SyslogMessage{
			Priority: 101,
			facility: 12,
			severity: 5,
			Version:  12,
		},
	},
	// Invalid, hostname too long
	{
		[]byte("<1>1 - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		SyslogMessage{},
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 262]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		SyslogMessage{},
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 281]",
		SyslogMessage{
			Priority:  1,
			facility:  0,
			severity:  1,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
		},
	},
	// Invalid, appname too long
	{
		[]byte("<1>1 - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 57]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 76]",
		SyslogMessage{
			Priority:  1,
			facility:  0,
			severity:  1,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
		},
	},
	{
		[]byte("<1>1 - host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 60]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Hostname: "host",
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 79]",
		SyslogMessage{
			Priority:  1,
			facility:  0,
			severity:  1,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
			Hostname:  "host",
		},
	},
	// Invalid, procid too long
	{
		[]byte("<1>1 - - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabX - -"),
		false,
		SyslogMessage{},
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 139]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, msgid too long
	{
		[]byte("<1>1 - - - - abcdefghilmnopqrstuvzabcdefghilmX -"),
		false,
		SyslogMessage{},
		"expecting a msgid (from 1 to max 32 US-ASCII characters) or a nil value [col 45]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Not print US-ASCII chars for hostname, appname, procid, and msgid
	{
		[]byte("<1>1 -   - - - -"),
		false,
		SyslogMessage{},
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 7]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - -   - - -"),
		false,
		SyslogMessage{},
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 9]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - -   - -"),
		false,
		SyslogMessage{},
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 11]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - -   -"),
		false,
		SyslogMessage{},
		"expecting a msgid (from 1 to max 32 US-ASCII characters) or a nil value [col 13]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, with malformed structured data
	{
		[]byte("<1>1 - - - - - X"),
		false,
		SyslogMessage{},
		"expecting a structured data section containing one or more elements (`[id( key=\"value\")*]+`) or a nil value [col 15]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, with empty structured data
	{
		[]byte("<1>1 - - - - - []"),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing space
	{
		[]byte("<1>1 - - - - - [ ]"),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing =
	{
		[]byte("<1>1 - - - - - [=]"),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing ]
	{
		[]byte("<1>1 - - - - - []]"),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing "
	{
		[]byte(`<1>1 - - - - - ["]`),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, too long structured data id
	{
		[]byte(`<1>1 - - - - - [abcdefghilmnopqrstuvzabcdefghilmX]`),
		false,
		SyslogMessage{},
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 48]",
		SyslogMessage{
			Priority:       1,
			facility:       0,
			severity:       1,
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, too long structured data param key
	{
		[]byte(`<1>1 - - - - - [id abcdefghilmnopqrstuvzabcdefghilmX="val"]`),
		false,
		SyslogMessage{},
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			StructuredData: map[string]map[string]string{
				"id": map[string]string{},
			},
		},
	},
	// Valid, minimal
	{
		[]byte("<1>1 - - - - - -"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte("<0>1 - - - - - -"),
		true,
		SyslogMessage{
			Priority: 0,
			facility: 0,
			severity: 0,
			Version:  1,
		},
		"",
		SyslogMessage{},
	},
	// Valid, average message
	{
		[]byte(`<29>1 2016-02-21T04:32:57+00:00 web1 someservice - - [origin x-service="someservice"][meta sequenceId="14125553"] 127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`),
		true,
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-02-21T04:32:57+00:00"),
			Hostname:  "web1",
			Appname:   "someservice",
			StructuredData: map[string]map[string]string{
				"origin": map[string]string{
					"x-service": "someservice",
				},
				"meta": map[string]string{
					"sequenceId": "14125553",
				},
			},
			Message: `127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, hostname, appname, procid, msgid can contain dashes
	{
		[]byte("<1>100 - host-name - - - -"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  100,
			Hostname: "host-name",
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte("<1>101 - host-name app-name - - -"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  101,
			Hostname: "host-name",
			Appname:  "app-name",
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte("<1>102 - host-name app-name proc-id - -"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  102,
			Hostname: "host-name",
			Appname:  "app-name",
			ProcID:   "proc-id",
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte("<1>103 - host-name app-name proc-id msg-id -"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  103,
			Hostname: "host-name",
			Appname:  "app-name",
			ProcID:   "proc-id",
			MsgID:    "msg-id",
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/0 structured data and w/o message, with other fields all max length
	{
		[]byte("<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm -"),
		true,
		SyslogMessage{
			Priority:  191,
			facility:  23,
			severity:  7,
			Version:   999,
			Timestamp: timeParse(time.RFC3339Nano, "2018-12-31T23:59:59.999999-23:59"),
			Hostname:  "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc",
			Appname:   "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef",
			ProcID:    "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab",
			MsgID:     "abcdefghilmnopqrstuvzabcdefghilm",
		},
		"",
		SyslogMessage{},
	},
	// Valid, all fields max length, with structured data and message
	{
		[]byte(`<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm [an@id key1="val1" key2="val2"][another@id key1="val1"] Some message "GET"`),
		true,
		SyslogMessage{
			Priority:  191,
			facility:  23,
			severity:  7,
			Version:   999,
			Timestamp: timeParse(time.RFC3339Nano, "2018-12-31T23:59:59.999999-23:59"),
			Hostname:  "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc",
			Appname:   "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef",
			ProcID:    "abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab",
			MsgID:     "abcdefghilmnopqrstuvzabcdefghilm",
			StructuredData: map[string]map[string]string{
				"an@id": map[string]string{
					"key1": "val1",
					"key2": "val2",
				},
				"another@id": map[string]string{
					"key1": "val1",
				},
			},
			Message: `Some message "GET"`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/o structure data, w/0 procid
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		SyslogMessage{
			facility:       4,
			severity:       2,
			Priority:       34,
			Version:        1,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:       "mymachine.example.com",
			Appname:        "su",
			ProcID:         ``,
			MsgID:          "ID47",
			StructuredData: nil,
			Message:        "BOM'su root' failed for lonvick on /dev/pts/8",
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/o structure data, w/o timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"),
		true,
		SyslogMessage{
			facility:       23,
			severity:       3,
			Priority:       187,
			Version:        222,
			Timestamp:      time.Time{},
			Hostname:       "mymachine.example.com",
			Appname:        "su",
			ProcID:         ``,
			MsgID:          "ID47",
			StructuredData: nil,
			Message:        "'su root' failed for lonvick on /dev/pts/8",
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/o structure data, w/o msgid
	{
		[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% Time to make the do-nuts."),
		true,
		SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        1,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       "192.0.2.1",
			Appname:        "myproc",
			ProcID:         "8710",
			MsgID:          ``,
			StructuredData: nil,
			Message:        "%% Time to make the do-nuts.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, w/o msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - -"),
		true,
		SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        2,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       ``,
			Appname:        ``,
			ProcID:         ``,
			MsgID:          ``,
			StructuredData: nil,
			Message:        ``,
		},
		"",
		SyslogMessage{},
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, empty msg
	{
		[]byte("<165>222 2003-08-24T05:14:15.000003-07:00 - - - - - "),
		true,
		SyslogMessage{
			facility:       20,
			severity:       5,
			Priority:       165,
			Version:        222,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       ``,
			Appname:        ``,
			ProcID:         ``,
			MsgID:          ``,
			StructuredData: nil,
			Message:        ``,
		},
		"",
		SyslogMessage{},
	},
	// Valid, with structured data is, w/o structured data params
	{
		[]byte("<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid] some_message"),
		true,
		SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  "host1",
			Appname:   "CROND",
			ProcID:    "10391",
			MsgID:     ``,
			StructuredData: map[string]map[string]string{
				"sdid": map[string]string{},
			},
			Message: "some_message",
		},
		"",
		SyslogMessage{},
	},
	// Valid, with structured data id, with structured data params
	{
		[]byte(`<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="⌘"] some_message`),
		true,
		SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  "host1",
			Appname:   "CROND",
			ProcID:    "10391",
			MsgID:     ``,
			StructuredData: map[string]map[string]string{
				"sdid": map[string]string{
					"x": "⌘",
				},
			},
			Message: "some_message",
		},
		"",
		SyslogMessage{},
	},
	// Valid, with structured data is, with structured data params
	{
		[]byte(`<78>2 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="hey \\u2318 hey"] some_message`),
		true,
		SyslogMessage{
			facility:  9,
			severity:  6,
			Priority:  78,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T00:04:01+00:00"),
			Hostname:  "host1",
			Appname:   "CROND",
			ProcID:    "10391",
			MsgID:     ``,
			StructuredData: map[string]map[string]string{
				"sdid": map[string]string{
					"x": `hey \u2318 hey`,
				},
			},
			Message: "some_message",
		},
		"",
		SyslogMessage{},
	},
	// Valid, with (escaped) backslash within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="\\valid"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"es": `\valid`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>52 2016-01-15T01:00:43Z hn S - - [meta one="\\one" two="\\two"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   52,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"one": `\one`,
					"two": `\two`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>53 2016-01-15T01:00:43Z hn S - - [meta one="\\one"][other two="\\two" double="\\a\\b"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   53,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"one": `\one`,
				},
				"other": map[string]string{
					"two":    `\two`,
					"double": `\a\b`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>51 2016-01-15T01:00:43Z hn S - - [meta es="\\double\\slash"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   51,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"es": `\double\slash`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>54 2016-01-15T01:00:43Z hn S - - [meta es="in \\middle of the string"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   54,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"es": `in \middle of the string`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>55 2016-01-15T01:00:43Z hn S - - [meta es="at the \\end"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   55,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"es": `at the \end`,
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, with control characters within structured data param value
	{
		[]byte("<29>50 2016-01-15T01:00:43Z hn S - - [meta es=\"\t5Ὂg̀9!℃ᾭGa b\"] 127.0.0.1 - - 1452819643 \"GET\""),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"es": "\t5Ὂg̀9!℃ᾭGa b",
				},
			},
			Message: `127.0.0.1 - - 1452819643 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, with utf8 within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta gr="κόσμε" es="ñ"][beta pr="₡"] 𐌼 "GET"`),
		true,
		SyslogMessage{
			Priority:  29,
			facility:  3,
			severity:  5,
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"gr": "κόσμε",
					"es": "ñ",
				},
				"beta": map[string]string{
					"pr": "₡",
				},
			},
			Message: `𐌼 "GET"`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, with structured data, w/o msg
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"),
		true,
		SyslogMessage{
			facility:  20,
			severity:  5,
			Priority:  165,
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  "example.com",
			Appname:   "evnts",
			MsgID:     "ID27",
			StructuredData: map[string]map[string]string{
				"exampleSDID@32473": map[string]string{
					"iut":         "3",
					"eventSource": "Application",
					"eventID":     "1011",
				},
				"examplePriority@32473": map[string]string{
					"class": "high",
				},
			},
		},
		"",
		SyslogMessage{},
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [id1][id1]"),
		false,
		SyslogMessage{},
		"duplicate structured data element id [col 66]",
		SyslogMessage{
			Priority:  165,
			facility:  20,
			severity:  5,
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  "example.com",
			Appname:   "evnts",
			MsgID:     "ID27",
			StructuredData: map[string]map[string]string{
				"id1": map[string]string{},
			},
		},
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [dupe e=\"1\"][id1][dupe class=\"l\"]"),
		false,
		SyslogMessage{},
		"duplicate structured data element id [col 79]",
		SyslogMessage{
			Priority:  165,
			facility:  20,
			severity:  5,
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  "example.com",
			Appname:   "evnts",
			MsgID:     "ID27",
			StructuredData: map[string]map[string]string{
				"id1": map[string]string{},
				"dupe": map[string]string{
					"e": "1",
				},
			},
		},
	},
	// Valid, with structured data w/o msg
	{
		[]byte(`<165>4 2003-10-11T22:14:15.003Z mymachine.it e - 1 [ex@32473 iut="3" eventSource="A"] An application event log entry...`),
		true,
		SyslogMessage{
			facility:  20,
			severity:  5,
			Priority:  165,
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  "mymachine.it",
			Appname:   "e",
			MsgID:     "1",
			StructuredData: map[string]map[string]string{
				"ex@32473": map[string]string{
					"iut":         "3",
					"eventSource": "A",
				},
			},
			Message: "An application event log entry...",
		},
		"",
		SyslogMessage{},
	},
	// Valid, with double quotes in the message
	{
		[]byte(`<29>1 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [origin x-service="svcname"][meta sequenceId="1"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "some-host-name",
			Appname:   "SEKRETPROGRAM",
			ProcID:    "prg",
			StructuredData: map[string]map[string]string{
				"origin": map[string]string{
					"x-service": "svcname",
				},
				"meta": map[string]string{
					"sequenceId": "1",
				},
			},
			Message: "127.0.0.1 - - 1452819643 \"GET\"",
		},
		"",
		SyslogMessage{},
	},
	// Valid, with empty structured data param value
	{
		[]byte(`<1>1 - - - - - [id pk=""]`),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			StructuredData: map[string]map[string]string{
				"id": map[string]string{
					"pk": "",
				},
			},
		},
		"",
		SyslogMessage{},
	},
	// Valid, with double quotes in the message and escaped character within param
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]"] some "mex"`),
		true,
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "some-host-name",
			Appname:   "SEKRETPROGRAM",
			ProcID:    "prg",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"escape": "]",
				},
			},
			Message: `some "mex"`,
		},
		"",
		SyslogMessage{},
	},
	// Valid, with escaped character within param
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\\"]`),
		true,
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "some-host-name",
			Appname:   "SEKRETPROGRAM",
			ProcID:    "prg",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"escape": `\`,
				},
			},
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\""]`),
		true,
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "some-host-name",
			Appname:   "SEKRETPROGRAM",
			ProcID:    "prg",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{
					"escape": `"`,
				},
			},
		},
		"",
		SyslogMessage{},
	},
	// Invalid, param value can not contain closing square bracket - ie., ]
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="]q"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   5,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="p]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Invalid, param value can not contain doublequote char - ie., ""
	{
		[]byte(`<29>4 2017-01-15T01:00:43Z hn S - - [meta escape="""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2017-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>6 2016-01-15T01:00:43Z hn S - - [meta escape="a""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 52]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   6,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>4 2018-01-15T01:00:43Z hn S - - [meta escape=""b"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2018-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Invalid, param value can not contain backslash - ie., \
	{
		[]byte(`<29>5 2019-01-15T01:00:43Z hn S - - [meta escape="\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 52]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   5,
			Timestamp: timeParse(time.RFC3339Nano, "2019-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>7 2019-01-15T01:00:43Z hn S - - [meta escape="a\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 53]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   7,
			Timestamp: timeParse(time.RFC3339Nano, "2019-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>8 2016-01-15T01:00:43Z hn S - - [meta escape="\n"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		SyslogMessage{},
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
		SyslogMessage{
			facility:  3,
			severity:  5,
			Priority:  29,
			Version:   8,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  "hn",
			Appname:   "S",
			StructuredData: map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Valid, message starting with byte order mark (BOM, \uFEFF)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\ufeff",
		},
		"",
		SyslogMessage{},
	},
	// Valid, greek
	{
		[]byte("<1>1 - - - - - - κόσμε"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "κόσμε",
		},
		"",
		SyslogMessage{},
	},
	// Valid, 2 octet sequence
	{
		[]byte("<1>1 - - - - - - "),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "",
		},
		"",
		SyslogMessage{},
	},
	// Valid, spanish (2 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xc3\xb1"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "ñ",
		},
		"",
		SyslogMessage{},
	},
	// Valid, colon currency sign (3 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xe2\x82\xa1"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "₡",
		},
		"",
		SyslogMessage{},
	},
	// Valid, gothic letter (4 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF \xf0\x90\x8c\xbc"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\ufeff 𐌼",
		},
		"",
		SyslogMessage{},
	},
	// Valid, 5 octet sequence
	{
		[]byte("<1>1 - - - - - - \xC8\x80\x30\x30\x30"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "Ȁ000",
		},
		"",
		SyslogMessage{},
	},
	// Valid, 6 octet sequence
	{
		[]byte("<1>1 - - - - - - \xE4\x80\x80\x30\x30\x30"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "䀀000",
		},
		"",
		SyslogMessage{},
	},
	// Valid, UTF-8 boundary conditions
	{
		[]byte("<1>1 - - - - - - \xC4\x90\x30\x30\x30"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "Đ000",
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte("<1>1 - - - - - - \x0D\x37\x46\x46"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\r7FF",
		},
		"",
		SyslogMessage{},
	},
	// Valid, Tamil poetry of Subramaniya Bharathiyar
	{
		[]byte("<1>1 - - - - - - யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Milanese)
	{
		[]byte("<1>1 - - - - - - Sôn bôn de magnà el véder, el me fa minga mal."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "Sôn bôn de magnà el véder, el me fa minga mal.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Romano)
	{
		[]byte("<1>1 - - - - - - Me posso magna' er vetro, e nun me fa male."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "Me posso magna' er vetro, e nun me fa male.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Braille)
	{
		[]byte("<1>1 - - - - - - ⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Sanskrit)
	{
		[]byte("<1>1 - - - - - - काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Urdu)
	{
		[]byte("<1>1 - - - - - - میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Yiddish)
	{
		[]byte("<1>1 - - - - - - איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "איך קען עסן גלאָז און עס טוט מיר נישט װײ.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Polish)
	{
		[]byte("<1>1 - - - - - - Mogę jeść szkło, i mi nie szkodzi."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "Mogę jeść szkło, i mi nie szkodzi.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Japanese)
	{
		[]byte("<1>1 - - - - - - 私はガラスを食べられます。それは私を傷つけません。"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "私はガラスを食べられます。それは私を傷つけません。",
		},
		"",
		SyslogMessage{},
	},
	// Valid, I Can Eat Glass (Arabic)
	{
		[]byte("<1>1 - - - - - - أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "أنا قادر على أكل الزجاج و هذا لا يؤلمني.",
		},
		"",
		SyslogMessage{},
	},
	// Valid, russian alphabet
	{
		[]byte("<1>1 - - - - - - абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "абвгдеёжзийклмнопрстуфхцчшщъыьэюя",
		},
		"",
		SyslogMessage{},
	},
	// Valid, armenian letters
	{
		[]byte("<1>1 - - - - - - ԰ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ՗՘ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆևֈ։֊֋֌֍֎֏"),
		true,
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\u0530ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ\u0557\u0558ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆև\u0588։֊\u058b\u058c֍֎֏",
		},
		"",
		SyslogMessage{},
	},
	// Valid, new line within message
	{
		[]byte("<1>1 - - - - - - x\x0Ay"),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  1,
			Message:  "x\ny",
		},
		"",
		SyslogMessage{},
	},
	{
		[]byte(`<1>2 - - - - - - x
y`),
		true,
		SyslogMessage{
			facility: 0,
			severity: 1,
			Priority: 1,
			Version:  2,
			Message:  "x\ny",
		},
		"",
		SyslogMessage{},
	},
	// Invalid, out of range code within message
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xC1"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF",
		},
	},
	{
		[]byte("<1>2 - - - - - - \xC1"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  2,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xc3\x28"), // invalid 2 octet sequence
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xc3",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc3\x28"), // invalid 2 octet sequence
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xc3",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xa0\xa1"), // invalid sequence identifier
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xa0\xa1"), // invalid sequence identifier
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xe2",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xe2",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xe2\x82",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xe2\x82",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xf0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xf0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xf0\x90",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xf0\x90",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xEF\xBB\xBF\xf0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xf0",
		},
	},
	// Invalid, impossible bytes
	{
		[]byte("<1>1 - - - - - - \xfe\xfe\xff\xff"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xfe"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xff"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x80\x80\x80\x80\xaf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x80\x80\x80\xaf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x80\x80\xaf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xf0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x80\xaf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xe0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc0\xaf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, maximum overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x83\xbf\xbf\xbf\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x87\xbf\xbf\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x8f\xbf\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xf0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x9f\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xe0",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc1\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
		},
	},
	// Invalid, illegal code positions, single utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xad\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xae\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xaf\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xb0\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbe\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbf\xbf"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	// Invalid, illegal code positions, paired utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80\xed\xb0\x80"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "\xed",
		},
	},
	// Invalid, out of range code within message after valid string
	{
		[]byte("<1>1 - - - - - - valid\xEF\xBB\xBF\xC1"),
		false,
		SyslogMessage{},
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 25]",
		SyslogMessage{
			Priority: 1,
			facility: 0,
			severity: 1,
			Version:  1,
			Message:  "valid\ufeff",
		},
	},

	// Invalid, missing whitespace after nil timestamp
	// {
	// 	[]byte("<1>1 -- - - - -"),
	// 	false,
	// 	nil,
	// 	"parsing error [col 6]",
	// 	SyslogMessage{
	// 		Priority: 1,
	// 		facility: 0,
	// 		severity: 1,
	// 		Version: 1,
	// 	},
	// },

	// (fixme) > evaluate non characters for UTF-8 security concerns, eg. \xef\xbf\xbe
}

func TestMachineParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(rxpad(string(tc.input), 50), func(t *testing.T) {
			t.Parallel()

			fsm := NewMachine()
			message, merr := fsm.Parse(tc.input, false)
			partial, perr := fsm.Parse(tc.input, true)

			if !tc.valid {
				assert.Equal(t, SyslogMessage{}, message)
				assert.Error(t, merr)
				assert.EqualError(t, merr, tc.errorString)

				assert.Equal(t, tc.partialValue, partial)
				assert.EqualError(t, perr, tc.errorString)
			} else {
				assert.Nil(t, merr)
				assert.NotEmpty(t, message)
				assert.Equal(t, message, partial)
				assert.Equal(t, merr, perr)
			}

			assert.Equal(t, tc.value, message)
		})
	}
}
