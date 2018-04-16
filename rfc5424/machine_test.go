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

func getUint8Address(x uint8) *uint8 {
	return &x
}

func createName(x []byte) string {
	str := string(x)
	lim := 60
	if len(str) > lim {
		return string(str[0:lim])
	}
	return str
}

type testCase struct {
	input        []byte
	valid        bool
	value        *SyslogMessage
	errorString  string
	partialValue *SyslogMessage
}

var testCases = []testCase{
	// Invalid, empty input
	{
		[]byte(""),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
		nil,
	},
	// Invalid, multiple syslog messages on multiple lines
	{
		[]byte(`<1>1 - - - - - -
	<2>1 - - - - - -`),
		false,
		nil,
		"parsing error [col 16]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			severity: getUint8Address(1),
			facility: getUint8Address(0),
			Version:  1,
		},
	},
	// Invalid, malformed pri
	{
		[]byte("(190>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
		nil,
	},
	// (fixme) > Malformed pri outputs wrong error (inner error regarding prival, not the outer one regarding pri)
	// {
	// 	[]byte("<87]123 -"),
	// 	false,
	// 	nil,
	// 	"expecting a priority value within angle brackets [col 3]",
	// 	nil, // nil since cannot reach version
	// },
	// Invalid, missing pri
	{
		[]byte("122 - - - - - -"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
		nil,
	},
	// Invalid, missing prival
	{
		[]byte("<>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
		nil,
	},
	// Invalid, prival with too much digits
	{
		[]byte("<19000021>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 4]",
		nil, // no valid partial message since was not able to reach and extract version (which is mandatory for a valid message)
	},
	// Invalid, prival too high
	{
		[]byte("<192>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 3]",
		nil,
	},
	// Invalid, 0 starting prival
	{
		[]byte("<002>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 2]",
		nil,
	},
	// Invalid, non numeric prival
	{
		[]byte("<aaa>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
		nil,
	},
	// Invalid, missing version
	{
		[]byte("<100> 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
		nil,
	},
	// Invalid, 0 version
	{
		[]byte("<103>0 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
		nil,
	},
	// Invalid, out of range version
	{
		[]byte("<101>1000 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 8]",
		nil,
	},
	// Invalid, non numeric version
	{
		[]byte("<102>abc 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
		nil,
	},
	// Invalid, wrong year
	{
		[]byte("<101>122 201-11-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 12]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  122,
		},
	},
	{
		[]byte("<101>189 0-11-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 10]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  189,
		},
	},
	// Invalid, wrong month
	{
		[]byte("<101>122 2018-112-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 16]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  122,
		},
	},
	// Invalid, wrong day
	{
		[]byte("<101>123 2018-02-32"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  123,
		},
	},
	// Invalid, wrong hour
	{
		[]byte("<101>124 2018-02-01:25:15Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 19]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  124,
		},
	},
	// Invalid, wrong minutes
	{
		[]byte("<101>125 2003-09-29T22:99:16Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 23]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  125,
		},
	},
	// Invalid, wrong seconds
	{
		[]byte("<101>126 2003-09-29T22:09:99Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 26]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  126,
		},
	},
	// Invalid, wrong sec fraction
	{
		[]byte("<101>127 2003-09-29T22:09:01.000000000009Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 35]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  127,
		},
	},
	{
		[]byte("<101>128 2003-09-29T22:09:01.Z"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 29]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  128,
		},
	},
	{
		[]byte("<101>28 2003-09-29T22:09:01."),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 28]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  28,
		},
	},
	// Invalid, wrong time offset
	{
		[]byte("<101>129 2003-09-29T22:09:01A"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 28]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  129,
		},
	},
	{
		[]byte("<101>130 2003-08-24T05:14:15.000003-24:00"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 37]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  130,
		},
	},
	{
		[]byte("<101>131 2003-08-24T05:14:15.000003-60:00"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 36]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  131,
		},
	},
	{
		[]byte("<101>132 2003-08-24T05:14:15.000003-07:61"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 39]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  132,
		},
	},
	// Invalid, non existing dates
	{
		[]byte("<101>11 2003-09-31T22:14:15.003Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:15.003Z\": day out of range [col 32]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  11,
		},
	},
	{
		[]byte("<101>12 2003-09-31T22:14:16Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:16Z\": day out of range [col 28]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  12,
		},
	},
	{
		[]byte("<101>12 2018-02-29T22:14:16+01:00"),
		false,
		nil,
		"parsing time \"2018-02-29T22:14:16+01:00\": day out of range [col 33]",
		&SyslogMessage{
			Priority: getUint8Address(101),
			facility: getUint8Address(12),
			severity: getUint8Address(5),
			Version:  12,
		},
	},
	// Invalid, hostname too long
	{
		[]byte("<1>1 - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 262]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 281]",
		&SyslogMessage{
			Priority:  getUint8Address(1),
			facility:  getUint8Address(0),
			severity:  getUint8Address(1),
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
		},
	},
	// Invalid, appname too long
	{
		[]byte("<1>1 - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 57]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 76]",
		&SyslogMessage{
			Priority:  getUint8Address(1),
			facility:  getUint8Address(0),
			severity:  getUint8Address(1),
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
		},
	},
	{
		[]byte("<1>1 - host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 60]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Hostname: getStringAddress("host"),
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 79]",
		&SyslogMessage{
			Priority:  getUint8Address(1),
			facility:  getUint8Address(0),
			severity:  getUint8Address(1),
			Version:   1,
			Timestamp: timeParse(time.RFC3339Nano, "2003-09-29T22:14:16Z"),
			Hostname:  getStringAddress("host"),
		},
	},
	// Invalid, procid too long
	{
		[]byte("<1>1 - - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabX - -"),
		false,
		nil,
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 139]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Invalid, msgid too long
	{
		[]byte("<1>1 - - - - abcdefghilmnopqrstuvzabcdefghilmX -"),
		false,
		nil,
		"expecting a msgid (from 1 to max 32 US-ASCII characters) [col 45]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Not print US-ASCII chars for hostname, appname, procid, and msgid
	{
		[]byte("<1>1 -   - - - -"),
		false,
		nil,
		"expecting an hostname (from 1 to max 255 US-ASCII characters) or a nil value [col 7]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - -   - - -"),
		false,
		nil,
		"expecting an app-name (from 1 to max 48 US-ASCII characters) or a nil value [col 9]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - -   - -"),
		false,
		nil,
		"expecting a procid (from 1 to max 128 US-ASCII characters) or a nil value [col 11]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - -   -"),
		false,
		nil,
		"expecting a msgid (from 1 to max 32 US-ASCII characters) [col 13]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Invalid, with malformed structured data
	{
		[]byte("<1>1 - - - - - X"),
		false,
		nil,
		"expecting a structured data section containing one or more elements (`[id ( key=\"value\")*]+`) or a nil value [col 15]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Invalid, with empty structured data
	{
		[]byte("<1>1 - - - - - []"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing space
	{
		[]byte("<1>1 - - - - - [ ]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing =
	{
		[]byte("<1>1 - - - - - [=]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing ]
	{
		[]byte("<1>1 - - - - - []]"),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, with structured data id containing "
	{
		[]byte(`<1>1 - - - - - ["]`),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 16]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, too long structured data id
	{
		[]byte(`<1>1 - - - - - [abcdefghilmnopqrstuvzabcdefghilmX]`),
		false,
		nil,
		"expecting a structured data element id (from 1 to max 32 US-ASCII characters; except `=`, ` `, `]`, and `\"` [col 48]",
		&SyslogMessage{
			Priority:       getUint8Address(1),
			facility:       getUint8Address(0),
			severity:       getUint8Address(1),
			Version:        1,
			StructuredData: nil,
		},
	},
	// Invalid, too long structured data param key
	{
		[]byte(`<1>1 - - - - - [id abcdefghilmnopqrstuvzabcdefghilmX=val]`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			StructuredData: &map[string]map[string]string{
				"id": map[string]string{},
			},
		},
	},
	// Valid, minimal
	{
		[]byte("<1>1 - - - - - -"),
		true,
		&SyslogMessage{
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Priority: getUint8Address(1),
			Version:  1,
		},
		"",
		nil,
	},
	{
		[]byte("<0>1 - - - - - -"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(0),
			facility: getUint8Address(0),
			severity: getUint8Address(0),
			Version:  1,
		},
		"",
		nil,
	},
	// Valid, with nil structured data and message, with other fields all max length
	{
		[]byte("<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm -"),
		true,
		&SyslogMessage{
			Priority:  getUint8Address(191),
			facility:  getUint8Address(23),
			severity:  getUint8Address(7),
			Version:   999,
			Timestamp: timeParse(time.RFC3339Nano, "2018-12-31T23:59:59.999999-23:59"),
			Hostname:  getStringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc"),
			Appname:   getStringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef"),
			ProcID:    getStringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab"),
			MsgID:     getStringAddress("abcdefghilmnopqrstuvzabcdefghilm"),
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/0 procid
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       getUint8Address(4),
			severity:       getUint8Address(2),
			Priority:       getUint8Address(34),
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
		nil,
	},
	// Valid, w/o structure data, w/o timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       getUint8Address(23),
			severity:       getUint8Address(3),
			Priority:       getUint8Address(187),
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
		nil,
	},
	// Valid, w/o structure data, w/o msgid
	{
		[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% Time to make the do-nuts."),
		true,
		&SyslogMessage{
			facility:       getUint8Address(20),
			severity:       getUint8Address(5),
			Priority:       getUint8Address(165),
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
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, w/o msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - -"),
		true,
		&SyslogMessage{
			facility:       getUint8Address(20),
			severity:       getUint8Address(5),
			Priority:       getUint8Address(165),
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
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, empty msg
	{
		[]byte("<165>222 2003-08-24T05:14:15.000003-07:00 - - - - - "),
		true,
		&SyslogMessage{
			facility:       getUint8Address(20),
			severity:       getUint8Address(5),
			Priority:       getUint8Address(165),
			Version:        222,
			Timestamp:      timeParse(time.RFC3339Nano, "2003-08-24T05:14:15.000003-07:00"),
			Hostname:       nil,
			Appname:        nil,
			ProcID:         nil,
			MsgID:          nil,
			StructuredData: nil,
			Message:        nil,
		},
		"",
		nil,
	},
	// Valid, with structured data is, w/o structured data params
	{
		[]byte("<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid] some_message"),
		true,
		&SyslogMessage{
			facility:  getUint8Address(9),
			severity:  getUint8Address(6),
			Priority:  getUint8Address(78),
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
		nil,
	},
	// Valid, with structured data is, with structured data params
	{
		[]byte(`<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="⌘"] some_message`),
		true,
		&SyslogMessage{
			facility:  getUint8Address(9),
			severity:  getUint8Address(6),
			Priority:  getUint8Address(78),
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
		nil,
	},
	// Valid, with structured data is, wit structured data params
	{
		[]byte(`<78>2 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="hey \\u2318 hey"] some_message`),
		true,
		&SyslogMessage{
			facility:  getUint8Address(9),
			severity:  getUint8Address(6),
			Priority:  getUint8Address(78),
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
		nil,
	},
	// Valid, with (escaped) backslash within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="\\valid"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			Priority:  getUint8Address(29),
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
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
		nil,
	},
	// Valid, with control characters within structured data param value
	{
		[]byte("<29>50 2016-01-15T01:00:43Z hn S - - [meta es=\"\b5Ὂg̀9! ℃ᾭG\"] 127.0.0.1 - - 1452819643 \"GET\""),
		true,
		&SyslogMessage{
			Priority:  getUint8Address(29),
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
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
		nil,
	},
	// Valid, with utf8 within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta gr="κόσμε" es="ñ"][beta pr="₡"] 𐌼 "GET"`),
		true,
		&SyslogMessage{
			Priority:  getUint8Address(29),
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Version:   50,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"gr": "κόσμε",
					"es": "ñ",
				},
				"beta": map[string]string{
					"pr": "₡",
				},
			},
			Message: getStringAddress(`𐌼 "GET"`),
		},
		"",
		nil,
	},
	// Valid, with structured data, w/o msg
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"),
		true,
		&SyslogMessage{
			facility:  getUint8Address(20),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(165),
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
		nil,
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [id1][id1]"),
		false,
		nil,
		"duplicate structured data element id [col 66]",
		&SyslogMessage{
			Priority:  getUint8Address(165),
			facility:  getUint8Address(20),
			severity:  getUint8Address(5),
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  getStringAddress("example.com"),
			Appname:   getStringAddress("evnts"),
			MsgID:     getStringAddress("ID27"),
			StructuredData: &map[string]map[string]string{
				"id1": map[string]string{},
			},
		},
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [dupe e=\"1\"][id1][dupe class=\"l\"]"),
		false,
		nil,
		"duplicate structured data element id [col 79]",
		&SyslogMessage{
			Priority:  getUint8Address(165),
			facility:  getUint8Address(20),
			severity:  getUint8Address(5),
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
			Hostname:  getStringAddress("example.com"),
			Appname:   getStringAddress("evnts"),
			MsgID:     getStringAddress("ID27"),
			StructuredData: &map[string]map[string]string{
				"id1": map[string]string{},
				"dupe": map[string]string{
					"e": "1",
				},
			},
		},
	},
	// Valid, with structured data w/o msg
	{
		[]byte("<165>4 2003-10-11T22:14:15.003Z mymachine.it e - 1 [ex@32473 iut=\"3\" eventSource=\"A\"] An application event log entry..."),
		true,
		&SyslogMessage{
			facility:  getUint8Address(20),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(165),
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
		nil,
	},
	// Valid, with double quotes in the message
	{
		[]byte(`<29>1 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [origin x-service="svcname"][meta sequenceId="1"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
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
		nil,
	},
	// Valid, with double quotes in the message and escaped character within param
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   2,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("some-host-name"),
			Appname:   getStringAddress("SEKRETPROGRAM"),
			ProcID:    getStringAddress("prg"),
			MsgID:     nil,
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{
					"escape": `\]`, // (todo) > verify output is correct wrt the input param value
				},
			},
			Message: getStringAddress("127.0.0.1 - - 1452819643 \"GET\""),
		},
		"",
		nil,
	},
	// Invalid, param value can not contain closing square bracket - ie., ]
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   3,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="]q"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 50]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   5,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="p]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Invalid, param value can not contain doublequote char - ie., ""
	{
		[]byte(`<29>4 2017-01-15T01:00:43Z hn S - - [meta escape="""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2017-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>6 2016-01-15T01:00:43Z hn S - - [meta escape="a""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 52]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   6,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>4 2018-01-15T01:00:43Z hn S - - [meta escape=""b"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting a structured data parameter (`key=\"value\"`, both part from 1 to max 32 US-ASCII characters; key cannot contain `=`, ` `, `]`, and `\"`, while value cannot contain `]`, backslash, and `\"` unless escaped) [col 51]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   4,
			Timestamp: timeParse(time.RFC3339Nano, "2018-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Invalid, param value can not contain backslash - ie., \
	{
		[]byte(`<29>5 2019-01-15T01:00:43Z hn S - - [meta escape="\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 52]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   5,
			Timestamp: timeParse(time.RFC3339Nano, "2019-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>7 2019-01-15T01:00:43Z hn S - - [meta escape="a\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 53]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   7,
			Timestamp: timeParse(time.RFC3339Nano, "2019-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	{
		[]byte(`<29>8 2016-01-15T01:00:43Z hn S - - [meta escape="\n"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		"expecting chars `]`, `\"`, and `\\` to be escaped within param value [col 51]",
		&SyslogMessage{
			facility:  getUint8Address(3),
			severity:  getUint8Address(5),
			Priority:  getUint8Address(29),
			Version:   8,
			Timestamp: timeParse(time.RFC3339Nano, "2016-01-15T01:00:43Z"),
			Hostname:  getStringAddress("hn"),
			Appname:   getStringAddress("S"),
			StructuredData: &map[string]map[string]string{
				"meta": map[string]string{},
			},
		},
	},
	// Valid, message starting with byte order mark (BOM, \uFEFF)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\ufeff"),
		},
		"",
		nil,
	},
	// Valid, greek
	{
		[]byte("<1>1 - - - - - - κόσμε"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("κόσμε"),
		},
		"",
		nil,
	},
	// Valid, 2 octet sequence
	{
		[]byte("<1>1 - - - - - - "),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress(""),
		},
		"",
		nil,
	},
	// Valid, spanish (2 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xc3\xb1"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("ñ"),
		},
		"",
		nil,
	},
	// Valid, colon currency sign (3 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xe2\x82\xa1"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("₡"),
		},
		"",
		nil,
	},
	// Valid, gothic letter (4 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF \xf0\x90\x8c\xbc"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\ufeff 𐌼"),
		},
		"",
		nil,
	},
	// Valid, 5 octet sequence
	{
		[]byte("<1>1 - - - - - - \xC8\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("Ȁ000"),
		},
		"",
		nil,
	},
	// Valid, 6 octet sequence
	{
		[]byte("<1>1 - - - - - - \xE4\x80\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("䀀000"),
		},
		"",
		nil,
	},
	// Valid, UTF-8 boundary conditions
	{
		[]byte("<1>1 - - - - - - \xC4\x90\x30\x30\x30"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("Đ000"),
		},
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - \x0D\x37\x46\x46"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\r7FF"),
		},
		"",
		nil,
	},
	// Valid, Tamil poetry of Subramaniya Bharathiyar
	{
		[]byte("<1>1 - - - - - - யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Milanese)
	{
		[]byte("<1>1 - - - - - - Sôn bôn de magnà el véder, el me fa minga mal."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("Sôn bôn de magnà el véder, el me fa minga mal."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Romano)
	{
		[]byte("<1>1 - - - - - - Me posso magna' er vetro, e nun me fa male."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("Me posso magna' er vetro, e nun me fa male."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Braille)
	{
		[]byte("<1>1 - - - - - - ⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Sanskrit)
	{
		[]byte("<1>1 - - - - - - काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Urdu)
	{
		[]byte("<1>1 - - - - - - میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Yiddish)
	{
		[]byte("<1>1 - - - - - - איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Polish)
	{
		[]byte("<1>1 - - - - - - Mogę jeść szkło, i mi nie szkodzi."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("Mogę jeść szkło, i mi nie szkodzi."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Japanese)
	{
		[]byte("<1>1 - - - - - - 私はガラスを食べられます。それは私を傷つけません。"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("私はガラスを食べられます。それは私を傷つけません。"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Arabic)
	{
		[]byte("<1>1 - - - - - - أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		},
		"",
		nil,
	},
	// Valid, russian alphabet
	{
		[]byte("<1>1 - - - - - - абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		},
		"",
		nil,
	},
	// Valid, armenian letters
	{
		[]byte("<1>1 - - - - - - ԰ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ՗՘ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆևֈ։֊֋֌֍֎֏"),
		true,
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\u0530ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ\u0557\u0558ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆև\u0588։֊\u058b\u058c֍֎֏"),
		},
		"",
		nil,
	},
	// Valid, new line within message
	{
		[]byte("<1>1 - - - - - - x\x0Ay"),
		true,
		&SyslogMessage{
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Priority: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("x\ny"),
		},
		"",
		nil,
	},
	{
		[]byte(`<1>2 - - - - - - x
y`),
		true,
		&SyslogMessage{
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Priority: getUint8Address(1),
			Version:  2,
			Message:  getStringAddress("x\ny"),
		},
		"",
		nil,
	},
	// Invalid, out of range code within message
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xC1"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF"),
		},
	},
	{
		[]byte("<1>2 - - - - - - \xC1"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  2,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xc3"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xc3"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 20]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xe2"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xe2"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xe2\x82"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xe2\x82"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 22]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xf0\x90"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 19]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xf0\x90"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 21]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xEF\xBB\xBF\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xf0"),
		},
	},
	// Invalid, impossible bytes
	{
		[]byte("<1>1 - - - - - - \xfe\xfe\xff\xff"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xfe"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xff"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Invalid, overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x80\x80\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x80\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x80\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x80\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xe0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc0\xaf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	// Invalid, maximum overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x83\xbf\xbf\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x87\xbf\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x8f\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x9f\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xe0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc1\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 17]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  nil,
		},
	},
	// Invalid, illegal code positions, single utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xad\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xae\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xaf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xb0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbe\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbf\xbf"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	// Invalid, illegal code positions, paired utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80\xed\xb0\x80"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 18]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("\xed"),
		},
	},
	// Invalid, out of range code within message after valid string
	{
		[]byte("<1>1 - - - - - - valid\xEF\xBB\xBF\xC1"),
		false,
		nil,
		"expecting a free-form optional message in UTF-8 (starting with or without BOM) [col 25]",
		&SyslogMessage{
			Priority: getUint8Address(1),
			facility: getUint8Address(0),
			severity: getUint8Address(1),
			Version:  1,
			Message:  getStringAddress("valid\ufeff"),
		},
	},
	// Invalid, missing whitespace after nil timestamp
	// {
	// 	[]byte("<1>1 -- - - - -"),
	// 	false,
	// 	nil,
	// 	"parsing error [col 6]",
	// 	&SyslogMessage{
	// 		Priority: getUint8Address(1),
	// 		facility: getUint8Address(0),
	// 		severity: getUint8Address(1),
	// 		Version: 1,
	// 	},
	// },

	// (fixme) > evaluate non characters for UTF-8 security concerns, eg. \xef\xbf\xbe
	// (fixme) > hostname & c. can contain dashes (eg. nil value), and spaces? it is correct?
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(createName(tc.input), func(t *testing.T) {
			t.Parallel()

			bestEffort := true
			fsm := NewMachine()
			message, merr := fsm.Parse(tc.input, nil)
			partial, perr := fsm.Parse(tc.input, &bestEffort)

			if !tc.valid {
				assert.Nil(t, message)
				assert.Error(t, merr)
				assert.EqualError(t, merr, tc.errorString)

				assert.Equal(t, tc.partialValue, partial)
				assert.EqualError(t, perr, tc.errorString)
			}
			if tc.valid {
				assert.Nil(t, merr)
				assert.NotEmpty(t, message)
				assert.Equal(t, message, partial)
				assert.Equal(t, merr, perr)
			}

			assert.Equal(t, tc.value, message)
		})
	}
}

// // This is here to avoid compiler optimizations that
// // could remove the actual call we are benchmarking
// // during benchmarks
// var benchParseResult *SyslogMessage

// func BenchmarkParse(b *testing.B) {
// 	for _, tc := range testCases {
// 		tc := tc
// 		b.Run(createName(tc.input), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				benchParseResult, _ = NewMachine().Parse(tc.input)
// 			}
// 		})
// 	}
// }
