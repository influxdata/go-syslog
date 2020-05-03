package rfc5424

import (
	"fmt"
	"math/rand"
	"testing"
	"unicode/utf8"

	"github.com/influxdata/go-syslog/v3"
	syslogtesting "github.com/influxdata/go-syslog/v3/testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
)

const BOM = "\xEF\xBB\xBF"

type testCase struct {
	input        []byte
	valid        bool
	value        syslog.Message
	errorString  string
	partialValue syslog.Message
}

var testCases = []testCase{
	// Invalid, empty input
	{
		[]byte(""),
		false,
		nil,
		fmt.Sprintf(ErrPri+ColumnPositionTemplate, 0),
		nil,
	},
	// Invalid, multiple syslog messages on multiple lines
	{
		[]byte(`<1>1 - - - - - -
		<2>1 - - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, new lines allowed only within message part
	{
		[]byte("<1>1 - \nhostname - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 7),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - host\x0Aname - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 11),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - \nan - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 9),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - a\x0An - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 10),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - - \npid - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 11),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - - p\x0Aid - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 12),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - - - \nmid -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 13),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - - - m\x0Aid -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 14),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, malformed pri
	{
		[]byte("(190>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPri+ColumnPositionTemplate, 0),
		nil,
	},
	// Malformed pri outputs wrong error
	{
		[]byte("<87]123 -"),
		false,
		nil,
		// (note) > machine can only understand that the ] char is not in the reachable states (just as any number would be in this situation), so it gives the error about the priority val submachine, not about the pri submachine (ie., <prival>)
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 3),
		nil, // nil since cannot reach version
	},
	// Invalid, missing pri
	{
		[]byte("122 - - - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrPri+ColumnPositionTemplate, 0),
		nil,
	},
	// Invalid, missing prival
	{
		[]byte("<>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 1),
		nil,
	},
	// Invalid, prival with too much digits
	{
		[]byte("<19000021>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 4),
		nil, // no valid partial message since was not able to reach and extract version (which is mandatory for a valid message)
	},
	// Invalid, prival too high
	{
		[]byte("<192>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 3),
		nil,
	},
	// Invalid, 0 starting prival
	{
		[]byte("<002>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 2),
		nil,
	},
	// Invalid, non numeric prival
	{
		[]byte("<aaa>122 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrPrival+ColumnPositionTemplate, 1),
		nil,
	},
	// Invalid, missing version
	{
		[]byte("<100> 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrVersion+ColumnPositionTemplate, 5),
		nil,
	},
	// Invalid, 0 version
	{
		[]byte("<103>0 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrVersion+ColumnPositionTemplate, 5),
		nil,
	},
	// Invalid, out of range version
	{
		[]byte("<101>1000 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrVersion+ColumnPositionTemplate, 8),
		(&SyslogMessage{}).SetVersion(100).SetPriority(101),
	},
	// Invalid, truncated after version whitespace
	{
		[]byte("<1>2 "),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 5),
		(&SyslogMessage{}).SetVersion(2).SetPriority(1),
	},
	// Invalid, truncated after version
	{
		[]byte("<1>1"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// fixme(leodido) > when space after multi-digit version is missing, the version error handler launches (should not)
	// {
	// 	[]byte("<3>22"),
	// 	false,
	// 	nil,
	// 	fmt.Sprintf(ErrParse+ColumnPositionTemplate, 6),
	// 	(&SyslogMessage{}).SetVersion(22).SetPriority(3),
	// },
	// Invalid, non numeric (also partially) version
	{
		[]byte("<1>3a"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		(&SyslogMessage{}).SetVersion(3).SetPriority(1),
	},
	{
		[]byte("<1>4a "),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		(&SyslogMessage{}).SetVersion(4).SetPriority(1),
	},
	{
		[]byte("<102>abc 2018-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrVersion+ColumnPositionTemplate, 5),
		nil,
	},
	// Invalid, letter rather than timestamp
	{
		[]byte("<1>5 A"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 5),
		(&SyslogMessage{}).SetVersion(5).SetPriority(1),
	},
	// Invalid, timestamp T and Z must be uppercase
	{
		[]byte(`<29>1 2006-01-02t15:04:05Z - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(29),
	},
	{
		[]byte(`<29>2 2006-01-02T15:04:05z - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 25),
		(&SyslogMessage{}).SetVersion(2).SetPriority(29),
	},
	// Invalid, wrong year
	{
		[]byte("<101>123 2"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 10),
		(&SyslogMessage{}).SetVersion(123).SetPriority(101),
	},
	{
		[]byte("<101>124 20"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 11),
		(&SyslogMessage{}).SetVersion(124).SetPriority(101),
	},
	{
		[]byte("<101>125 201"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 12),
		(&SyslogMessage{}).SetVersion(125).SetPriority(101),
	},
	{
		[]byte("<101>125 2013"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 13),
		(&SyslogMessage{}).SetVersion(125).SetPriority(101),
	},
	{
		[]byte("<101>126 2013-"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 14),
		(&SyslogMessage{}).SetVersion(126).SetPriority(101),
	},
	{
		[]byte("<101>122 201-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 12),
		(&SyslogMessage{}).SetVersion(122).SetPriority(101),
	},
	{
		[]byte("<101>189 0-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 10),
		(&SyslogMessage{}).SetVersion(189).SetPriority(101),
	},
	// Invalid, wrong month
	{
		[]byte("<101>121 2018-112-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(121).SetPriority(101),
	},
	// Invalid, wrong day
	{
		[]byte("<101>123 2018-02-32"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 18),
		(&SyslogMessage{}).SetVersion(123).SetPriority(101),
	},
	// Invalid, wrong hour
	{
		[]byte("<101>124 2018-02-01:25:15Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 19),
		(&SyslogMessage{}).SetVersion(124).SetPriority(101),
	},
	// Invalid, wrong minutes
	{
		[]byte("<101>125 2003-09-29T22:99:16Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 23),
		(&SyslogMessage{}).SetVersion(125).SetPriority(101),
	},
	// Invalid, wrong seconds
	{
		[]byte("<101>126 2003-09-29T22:09:99Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 26),
		(&SyslogMessage{}).SetVersion(126).SetPriority(101),
	},
	// Invalid, wrong sec fraction
	{
		[]byte("<101>127 2003-09-29T22:09:01.000000000009Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 35),
		(&SyslogMessage{}).SetVersion(127).SetPriority(101),
	},
	{
		[]byte("<101>128 2003-09-29T22:09:01.Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 29),
		(&SyslogMessage{}).SetVersion(128).SetPriority(101),
	},
	{
		[]byte("<101>28 2003-09-29T22:09:01."),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 28),
		(&SyslogMessage{}).SetVersion(28).SetPriority(101),
	},
	// Invalid, wrong time offset
	{
		[]byte("<101>129 2003-09-29T22:09:01A"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 28),
		(&SyslogMessage{}).SetVersion(129).SetPriority(101),
	},
	{
		[]byte("<101>130 2003-08-24T05:14:15.000003-24:00"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 37),
		(&SyslogMessage{}).SetVersion(130).SetPriority(101),
	},
	{
		[]byte("<101>131 2003-08-24T05:14:15.000003-60:00"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 36),
		(&SyslogMessage{}).SetVersion(131).SetPriority(101),
	},
	{
		[]byte("<101>132 2003-08-24T05:14:15.000003-07:61"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 39),
		(&SyslogMessage{}).SetVersion(132).SetPriority(101),
	},
	{
		[]byte(`<29>1 2006-01-02T15:04:05Z+07:00 - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 26), // after the Z (valid and complete timestamp) it searches for a whitespace
		(&SyslogMessage{}).SetVersion(1).SetTimestamp("2006-01-02T15:04:05Z").SetPriority(29),
	},
	// Invalid, non existing dates
	{
		[]byte("<101>11 2003-09-31T22:14:15.003Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:15.003Z\": day out of range [col 32]",
		(&SyslogMessage{}).SetVersion(11).SetPriority(101),
	},
	{
		[]byte("<101>12 2003-09-31T22:14:16Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:16Z\": day out of range [col 28]",
		(&SyslogMessage{}).SetVersion(12).SetPriority(101),
	},
	{
		[]byte("<101>12 2018-02-29T22:14:16+01:00"),
		false,
		nil,
		"parsing time \"2018-02-29T22:14:16+01:00\": day out of range [col 33]",
		(&SyslogMessage{}).SetVersion(12).SetPriority(101),
	},
	// Invalid, hostname too long
	{
		[]byte("<1>1 - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 262),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 281),
		(&SyslogMessage{}).SetVersion(1).SetTimestamp("2003-09-29T22:14:16Z").SetPriority(1),
	},
	// Invalid, appname too long
	{
		[]byte("<1>1 - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 57),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 76),
		(&SyslogMessage{}).SetVersion(1).SetTimestamp("2003-09-29T22:14:16Z").SetPriority(1),
	},
	{
		[]byte("<1>1 - host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 60),
		(&SyslogMessage{}).SetVersion(1).SetHostname("host").SetPriority(1),
	},
	{
		[]byte("<1>1 2004-09-29T22:14:16Z host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 79),
		(&SyslogMessage{}).SetVersion(1).SetTimestamp("2004-09-29T22:14:16Z").SetHostname("host").SetPriority(1),
	},
	// Invalid, procid too long
	{
		[]byte("<1>1 - - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabX - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 139),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, msgid too long
	{
		[]byte("<1>1 - - - - abcdefghilmnopqrstuvzabcdefghilmX -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 45),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Not print US-ASCII chars for hostname, appname, procid, and msgid
	{
		[]byte("<1>1 -   - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 7),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - -   - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 9),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - -   - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 11),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	{
		[]byte("<1>1 - - - -   -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 13),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with malformed structured data
	{
		[]byte("<1>1 - - - - - X"),
		false,
		nil,
		fmt.Sprintf(ErrStructuredData+ColumnPositionTemplate, 15),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with empty structured data
	{
		[]byte("<1>1 - - - - - []"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with structured data id containing space
	{
		[]byte("<1>1 - - - - - [ ]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with structured data id containing =
	{
		[]byte("<1>1 - - - - - [=]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with structured data id containing ]
	{
		[]byte("<1>1 - - - - - []]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, with structured data id containing "
	{
		[]byte(`<6>1 - - - - - ["]`),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		(&SyslogMessage{}).SetVersion(1).SetPriority(6),
	},
	// Invalid, too long structured data id
	{
		[]byte(`<1>1 - - - - - [abcdefghilmnopqrstuvzabcdefghilmX]`),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 48),
		(&SyslogMessage{}).SetVersion(1).SetPriority(1),
	},
	// Invalid, too long structured data param key
	{
		[]byte(`<1>1 - - - - - [id abcdefghilmnopqrstuvzabcdefghilmX="val"]`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		(&SyslogMessage{}).SetVersion(1).SetElementID("id").SetPriority(1),
	},
	// Valid, minimal
	{
		[]byte("<10>1 - - - - - -"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetPriority(10),
		"",
		nil,
	},
	{
		[]byte("<0>1 - - - - - -"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetPriority(0),
		"",
		nil,
	},
	// Valid, average message
	{
		[]byte(`<29>1 2016-02-21T04:32:57+00:00 web1 someservice - - [origin x-service="someservice"][meta sequenceId="14125553"] 127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`),
		true,
		(&SyslogMessage{}).
			SetVersion(1).
			SetHostname("web1").
			SetAppname("someservice").
			SetTimestamp("2016-02-21T04:32:57+00:00").
			SetParameter("origin", "x-service", "someservice").
			SetParameter("meta", "sequenceId", "14125553").
			SetMessage(`127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`).
			SetPriority(29),
		"",
		nil,
	},
	// Valid, hostname, appname, procid, msgid can contain dashes
	{
		[]byte("<1>100 - host-name - - - -"),
		true,
		(&SyslogMessage{}).SetVersion(100).SetHostname("host-name").SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>101 - host-name app-name - - -"),
		true,
		(&SyslogMessage{}).SetVersion(101).SetHostname("host-name").SetAppname("app-name").SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>102 - host-name app-name proc-id - -"),
		true,
		(&SyslogMessage{}).
			SetVersion(102).
			SetHostname("host-name").
			SetAppname("app-name").
			SetProcID("proc-id").
			SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>103 - host-name app-name proc-id msg-id -"),
		true,
		(&SyslogMessage{}).
			SetVersion(103).
			SetHostname("host-name").
			SetAppname("app-name").
			SetProcID("proc-id").
			SetMsgID("msg-id").
			SetPriority(1),
		"",
		nil,
	},
	// Valid, w/0 structured data and w/o message, with other fields all max length
	{
		[]byte("<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm -"),
		true,
		(&SyslogMessage{}).
			SetVersion(999).
			SetHostname("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc").
			SetAppname("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef").
			SetTimestamp("2018-12-31T23:59:59.999999-23:59").
			SetProcID("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab").
			SetMsgID("abcdefghilmnopqrstuvzabcdefghilm").
			SetPriority(191),
		"",
		nil,
	},
	// Valid, all fields max length, with structured data and message
	{
		[]byte(`<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm [an@id key1="val1" key2="val2"][another@id key1="val1"] Some message "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(999).
			SetHostname("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc").
			SetAppname("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef").
			SetTimestamp("2018-12-31T23:59:59.999999-23:59").
			SetProcID("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab").
			SetMsgID("abcdefghilmnopqrstuvzabcdefghilm").
			SetParameter("an@id", "key1", "val1").
			SetParameter("an@id", "key2", "val2").
			SetParameter("another@id", "key1", "val1").
			SetMessage(`Some message "GET"`).
			SetPriority(191),
		"",
		nil,
	},
	// Valid, w/o structure data, w/0 procid
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		(&SyslogMessage{}).
			SetVersion(1).
			SetHostname("mymachine.example.com").
			SetAppname("su").
			SetTimestamp("2003-10-11T22:14:15.003Z").
			SetMsgID("ID47").
			SetMessage("BOM'su root' failed for lonvick on /dev/pts/8").
			SetPriority(34),
		"",
		nil,
	},
	// Valid, w/o structure data, w/o timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"),
		true,
		(&SyslogMessage{}).
			SetVersion(222).
			SetHostname("mymachine.example.com").
			SetAppname("su").
			SetMsgID("ID47").
			SetMessage("'su root' failed for lonvick on /dev/pts/8").
			SetPriority(187),
		"",
		nil,
	},
	// Valid, w/o structure data, w/o msgid
	{
		[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% Time to make the do-nuts."),
		true,
		(&SyslogMessage{}).
			SetVersion(1).
			SetHostname("192.0.2.1").
			SetAppname("myproc").
			SetTimestamp("2003-08-24T05:14:15.000003-07:00").
			SetProcID("8710").
			SetMessage("%% Time to make the do-nuts.").
			SetPriority(165),
		"",
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, w/o msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - -"),
		true,
		(&SyslogMessage{}).
			SetVersion(2).
			SetTimestamp("2003-08-24T05:14:15.000003-07:00").
			SetPriority(165),
		"",
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, empty msg
	{
		[]byte("<165>222 2003-08-24T05:14:15.000002-01:00 - - - - - "),
		true,
		(&SyslogMessage{}).
			SetVersion(222).
			SetTimestamp("2003-08-24T05:14:15.000002-01:00").
			SetPriority(165),
		"",
		nil,
	},
	// Valid, with structured data is, w/o structured data params
	{
		[]byte("<78>5 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid] some_message"),
		true,
		(&SyslogMessage{}).
			SetVersion(5).
			SetHostname("host1").
			SetAppname("CROND").
			SetTimestamp("2016-01-15T00:04:01+00:00").
			SetProcID("10391").
			SetMessage("some_message").
			SetElementID("sdid").
			SetPriority(78),
		"",
		nil,
	},
	// Valid, with structured data id, with structured data params
	{
		[]byte(`<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="⌘"] some_message`),
		true,
		(&SyslogMessage{}).
			SetVersion(1).
			SetHostname("host1").
			SetAppname("CROND").
			SetTimestamp("2016-01-15T00:04:01+00:00").
			SetProcID("10391").
			SetMessage("some_message").
			SetParameter("sdid", "x", "⌘").
			SetPriority(78),
		"",
		nil,
	},
	// Valid, with structured data is, with structured data params
	{
		[]byte(`<78>2 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="hey \\u2318 hey"] some_message`),
		true,
		(&SyslogMessage{}).
			SetVersion(2).
			SetHostname("host1").
			SetAppname("CROND").
			SetTimestamp("2016-01-15T00:04:01+00:00").
			SetProcID("10391").
			SetMessage("some_message").
			SetParameter("sdid", "x", `hey \\u2318 hey`).
			SetPriority(78),
		"",
		nil,
	},
	// Valid, with (escaped) backslash within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="\\valid"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(50).
			SetHostname("hn").
			SetAppname("S").
			SetTimestamp("2016-01-15T01:00:43Z").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "es", `\\valid`).
			SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>52 2016-01-15T01:00:43Z hn S - - [meta one="\\one" two="\\two"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(52).
			SetHostname("hn").
			SetAppname("S").
			SetTimestamp("2016-01-15T01:00:43Z").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "one", `\\one`).
			SetParameter("meta", "two", `\\two`).
			SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>53 2016-01-15T01:00:43Z hn S - - [meta one="\\one"][other two="\\two" double="\\a\\b"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(53).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "one", `\\one`).
			SetParameter("other", "two", `\\two`).
			SetParameter("other", "double", `\\a\\b`).
			SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>51 2016-01-15T01:00:43Z hn S - - [meta es="\\double\\slash"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(51).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "es", `\\double\\slash`).
			SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>54 2016-01-15T01:00:43Z hn S - - [meta es="in \\middle of the string"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(54).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "es", `in \\middle of the string`).
			SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>55 2016-01-15T01:00:43Z hn S - - [meta es="at the \\end"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(55).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "es", `at the \\end`).
			SetPriority(29),
		"",
		nil,
	},
	// Valid, with control characters within structured data param value
	{
		[]byte("<29>50 2016-01-15T01:00:43Z hn S - - [meta es=\"\t5Ὂg̀9!℃ᾭGa b\"] 127.0.0.1 - - 1452819643 \"GET\""),
		true,
		(&SyslogMessage{}).
			SetVersion(50).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`127.0.0.1 - - 1452819643 "GET"`).
			SetParameter("meta", "es", "\t5Ὂg̀9!℃ᾭGa b").
			SetPriority(29),
		"",
		nil,
	},
	// Valid, with utf8 within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta gr="κόσμε" es="ñ"][beta pr="₡"] 𐌼 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(50).
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("hn").
			SetAppname("S").
			SetMessage(`𐌼 "GET"`).
			SetParameter("meta", "gr", "κόσμε").
			SetParameter("meta", "es", "ñ").
			SetParameter("beta", "pr", "₡").
			SetPriority(29),
		"",
		nil,
	},
	// Valid, with structured data, w/o msg
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"),
		true,
		(&SyslogMessage{}).
			SetVersion(3).
			SetTimestamp("2003-10-11T22:14:15.003Z").
			SetHostname("example.com").
			SetAppname("evnts").
			SetMsgID("ID27").
			SetParameter("exampleSDID@32473", "iut", "3").
			SetParameter("exampleSDID@32473", "eventSource", "Application").
			SetParameter("exampleSDID@32473", "eventID", "1011").
			SetParameter("examplePriority@32473", "class", "high").
			SetPriority(165),
		"",
		nil,
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [id1][id1]"),
		false,
		nil,
		"duplicate structured data element id [col 66]",
		(&SyslogMessage{}).
			SetVersion(3).
			SetTimestamp("2003-10-11T22:14:15.003Z").
			SetHostname("example.com").
			SetAppname("evnts").
			SetMsgID("ID27").
			SetElementID("id1").
			SetPriority(165),
	},
	// Invalid, with duplicated structured data id
	{
		[]byte("<165>3 2003-10-12T22:14:15.003Z example.com evnts - ID27 [dupe e=\"1\"][id1][dupe class=\"l\"]"),
		false,
		nil,
		"duplicate structured data element id [col 79]",
		(&SyslogMessage{}).
			SetVersion(3).
			SetTimestamp("2003-10-12T22:14:15.003Z").
			SetHostname("example.com").
			SetAppname("evnts").
			SetMsgID("ID27").
			SetElementID("id1").
			SetParameter("dupe", "e", "1").
			SetPriority(165),
	},
	// Valid, with structured data w/o msg
	{
		[]byte(`<165>4 2003-10-11T22:14:15.003Z mymachine.it e - 1 [ex@32473 iut="3" eventSource="A"] An application event log entry...`),
		true,
		(&SyslogMessage{}).
			SetVersion(4).
			SetMessage("An application event log entry...").
			SetTimestamp("2003-10-11T22:14:15.003Z").
			SetHostname("mymachine.it").
			SetAppname("e").
			SetMsgID("1").
			SetParameter("ex@32473", "iut", "3").
			SetParameter("ex@32473", "eventSource", "A").
			SetPriority(165),
		"",
		nil,
	},
	// Valid, with double quotes in the message
	{
		[]byte(`<29>1 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [origin x-service="svcname"][meta sequenceId="1"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		(&SyslogMessage{}).
			SetVersion(1).
			SetMessage("127.0.0.1 - - 1452819643 \"GET\"").
			SetTimestamp("2016-01-15T01:00:43Z").
			SetHostname("some-host-name").
			SetAppname("SEKRETPROGRAM").
			SetProcID("prg").
			SetParameter("origin", "x-service", "svcname").
			SetParameter("meta", "sequenceId", "1").
			SetPriority(29),
		"",
		nil,
	},
	// Valid, with empty structured data param value
	{
		[]byte(`<1>1 - - - - - [id pk=""]`),
		true,
		(&SyslogMessage{}).SetVersion(1).SetParameter("id", "pk", "").SetPriority(1),
		"",
		nil,
	},
	// Valid, with escaped character within param value
	{
		[]byte(`<29>2 2016-01-15T01:00:44Z some-host-name SEKRETPROGRAM prg - [meta escape="\]"] some "mex"`),
		true,
		(&SyslogMessage{}).SetVersion(2).SetMessage(`some "mex"`).SetTimestamp("2016-01-15T01:00:44Z").SetHostname("some-host-name").SetAppname("SEKRETPROGRAM").SetProcID("prg").SetParameter("meta", "escape", `\]`).SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\\"]`),
		true,
		(&SyslogMessage{}).SetVersion(2).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("some-host-name").SetAppname("SEKRETPROGRAM").SetProcID("prg").SetParameter("meta", "escape", `\\`).SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\""]`),
		true,
		(&SyslogMessage{}).SetVersion(2).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("some-host-name").SetAppname("SEKRETPROGRAM").SetProcID("prg").SetParameter("meta", "escape", `\"`).SetPriority(29),
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]\"\\\\\]\""]`),
		true,
		(&SyslogMessage{}).SetVersion(2).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("some-host-name").SetAppname("SEKRETPROGRAM").SetProcID("prg").SetParameter("meta", "escape", `\]\"\\\\\]\"`).SetPriority(29),
		"",
		nil,
	},
	// Invalid, param value can not contain closing square bracket - ie., ]
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 50),
		(&SyslogMessage{}).SetVersion(3).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="]q"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 50),
		(&SyslogMessage{}).SetVersion(5).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="p]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 51),
		(&SyslogMessage{}).SetVersion(4).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	// Invalid, param value can not contain doublequote char - ie., ""
	{
		[]byte(`<29>4 2017-01-15T01:00:43Z hn S - - [meta escape="""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		(&SyslogMessage{}).SetVersion(4).SetTimestamp("2017-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>6 2016-01-15T01:00:43Z hn S - - [meta escape="a""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 52),
		(&SyslogMessage{}).SetVersion(6).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>4 2018-01-15T01:00:43Z hn S - - [meta escape=""b"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		(&SyslogMessage{}).SetVersion(4).SetTimestamp("2018-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	// Invalid, param value can not contain backslash - ie., \
	{
		[]byte(`<29>5 2019-01-15T01:00:43Z hn S - - [meta escape="\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 52),
		(&SyslogMessage{}).SetVersion(5).SetTimestamp("2019-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>7 2019-01-15T01:00:43Z hn S - - [meta escape="a\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 53),
		(&SyslogMessage{}).SetVersion(7).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetTimestamp("2019-01-15T01:00:43Z").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	{
		[]byte(`<29>8 2016-01-15T01:00:43Z hn S - - [meta escape="\n"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 51),
		(&SyslogMessage{}).SetVersion(8).SetTimestamp("2016-01-15T01:00:43Z").SetHostname("hn").SetAppname("S").SetElementID("meta").SetPriority(29),
	},
	// Valid, message starting with byte order mark (BOM, \uFEFF)
	{
		[]byte("<1>8 - - - - - - " + BOM),
		true,
		(&SyslogMessage{}).SetVersion(8).SetMessage("\ufeff").SetPriority(1),
		"",
		nil,
	},
	// Valid, greek
	{
		[]byte("<1>1 - - - - - - " + BOM + "κόσμε"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "κόσμε").SetPriority(1),
		"",
		nil,
	},
	// Valid, 2 octet sequence
	{
		[]byte("<1>1 - - - - - - " + BOM + ""),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "").SetPriority(1),
		"",
		nil,
	},
	// Valid, spanish (2 octet sequence)
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xc3\xb1"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "ñ").SetPriority(1),
		"",
		nil,
	},
	// Valid, colon currency sign (3 octet sequence)
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xe2\x82\xa1"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "₡").SetPriority(1),
		"",
		nil,
	},
	// Valid, gothic letter (4 octet sequence)
	{
		[]byte("<3>1 - - - - - - " + BOM + " \xf0\x90\x8c\xbc"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + " 𐌼").SetPriority(3),
		"",
		nil,
	},
	// Valid, 5 octet sequence
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xC8\x80\x30\x30\x30"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "Ȁ000").SetPriority(1),
		"",
		nil,
	},
	// Valid, 6 octet sequence
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xE4\x80\x80\x30\x30\x30"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "䀀000").SetPriority(1),
		"",
		nil,
	},
	// Valid, UTF-8 boundary conditions
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xC4\x90\x30\x30\x30"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "Đ000").SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\x0D\x37\x46\x46"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "\r7FF").SetPriority(1),
		"",
		nil,
	},
	// Valid, Tamil poetry of Subramaniya Bharathiyar
	{
		[]byte("<1>1 - - - - - - " + BOM + "யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்.").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Milanese)
	{
		[]byte("<1>1 - - - - - - " + BOM + "Sôn bôn de magnà el véder, el me fa minga mal."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "Sôn bôn de magnà el véder, el me fa minga mal.").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Romano)
	{
		[]byte("<1>1 - - - - - - " + BOM + "Me posso magna' er vetro, e nun me fa male."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "Me posso magna' er vetro, e nun me fa male.").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Braille)
	{
		[]byte("<1>1 - - - - - - " + BOM + "⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Sanskrit)
	{
		[]byte("<1>1 - - - - - - " + BOM + "काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Urdu)
	{
		[]byte("<1>1 - - - - - - " + BOM + "میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Yiddish)
	{
		[]byte("<1>1 - - - - - - " + BOM + "איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "איך קען עסן גלאָז און עס טוט מיר נישט װײ.").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Polish)
	{
		[]byte("<1>1 - - - - - - " + BOM + "Mogę jeść szkło, i mi nie szkodzi."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "Mogę jeść szkło, i mi nie szkodzi.").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Japanese)
	{
		[]byte("<1>1 - - - - - - " + BOM + "私はガラスを食べられます。それは私を傷つけません。"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "私はガラスを食べられます。それは私を傷つけません。").SetPriority(1),
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Arabic)
	{
		[]byte("<1>1 - - - - - - " + BOM + "أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "أنا قادر على أكل الزجاج و هذا لا يؤلمني.").SetPriority(1),
		"",
		nil,
	},
	// Valid, russian alphabet
	{
		[]byte("<1>1 - - - - - - " + BOM + "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "абвгдеёжзийклмнопрстуфхцчшщъыьэюя").SetPriority(1),
		"",
		nil,
	},
	// Valid, armenian letters
	{
		[]byte("<1>1 - - - - - - " + BOM + "԰ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ՗՘ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆևֈ։֊֋֌֍֎֏"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM + "\u0530ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ\u0557\u0558ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆև\u0588։֊\u058b\u058c֍֎֏").SetPriority(1),
		"",
		nil,
	},
	// Valid and compliant, characters in any encoding are allowed as long as MSG does not start with BOM
	// fondue requires a lot of cheese (swiss german)
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("fondü bruucht vell chääs")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("fondü bruucht vell chääs")).SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - not starting with" + BOM + isoLatin1String("fondü bruucht vell chääs")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage("not starting with" + BOM + isoLatin1String("fondü bruucht vell chääs")).SetPriority(1),
		"",
		nil,
	},
	// Valid and compliant iso-latin-1
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("ö")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("ö")).SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("öö")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("öö")).SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("ööö")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("ööö")).SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("öööö")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("öööö")).SetPriority(1),
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - " + isoLatin1String("ööööö")),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage(isoLatin1String("ööööö")).SetPriority(1),
		"",
		nil,
	},
	// Valid, new line within message
	{
		[]byte("<1>1 - - - - - - x\x0Ay"),
		true,
		(&SyslogMessage{}).SetVersion(1).SetMessage("x\ny").SetPriority(1),
		"",
		nil,
	},
	{
		[]byte(`<1>3 - - - - - - x
y`),
		true,
		(&SyslogMessage{}).SetVersion(3).SetMessage("x\ny").SetPriority(1),
		"",
		nil,
	},
	// Invalid, missing whitespace after nil timestamp
	{
		[]byte("<1>10 -- - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 7),
		(&SyslogMessage{}).SetVersion(10).SetPriority(1),
	},
}

var nonCompliantMsgTestCases = []testCase{
	// Invalid, out of range code within message
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xC1"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		(&SyslogMessage{}).SetVersion(1).SetMessage(BOM).SetPriority(1),
	},
	{
		[]byte("<1>4 - - - - - - " + BOM + "\xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 4, syslogtesting.StringAddress(BOM+"\xc3")),
	},
	{
		[]byte("<7>1 - - - - - - " + BOM + "\xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(7, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xe2")),
	},
	{
		[]byte("<6>1 - - - - - - " + BOM + "\xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 22),
		genMessageWithPartialMessage(6, 1, syslogtesting.StringAddress(BOM+"\xe2\x82")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xf0")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 22),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xf0\x90")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xf0")),
	},
	// Invalid, impossible bytes
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xfe\xfe\xff\xff"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xfe"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xff"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	// Invalid, overlong sequences
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xfc\x80\x80\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf8\x80\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>3 - - - - - - " + BOM + "\xf0\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 3, syslogtesting.StringAddress(BOM+"\xf0")),
	},
	{
		[]byte("<1>3 - - - - - - " + BOM + "\xe0\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 3, syslogtesting.StringAddress(BOM+"\xe0")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xc0\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	// Invalid, maximum overlong sequences
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xfc\x83\xbf\xbf\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf8\x87\xbf\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xf0\x8f\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xf0")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xe0\x9f\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xe0")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xc1\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 20),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM)),
	},
	// Invalid, illegal code positions, single utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xa0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xa0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xad\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xae\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<22>23 - - - - - - " + BOM + "\xed\xaf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 23),
		genMessageWithPartialMessage(22, 23, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xb0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xbe\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	// Invalid, illegal code positions, paired utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - " + BOM + "\xed\xa0\x80\xed\xb0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 21),
		genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"\xed")),
	},
	// (fixme) > evaluate non characters for UTF-8 security concerns, eg. \xef\xbf\xbe
}

func genMessageWithPartialMessage(p uint8, v uint16, m *string) *SyslogMessage {
	mex := (&SyslogMessage{}).SetVersion(v).SetPriority(p).(*SyslogMessage)
	mex.Message = m
	return mex
}

// genIncompleteTimestampTestCases generates test cases with incomplete timestamp part.
func genIncompleteTimestampTestCases() []testCase {
	incompleteTimestamp := []byte("2003-11-02T23:12:46.012345")
	prefix := []byte("<1>1 ")
	mex := (&SyslogMessage{}).SetVersion(1).SetPriority(1)
	tCases := make([]testCase, 0, len(incompleteTimestamp))
	prev := make([]byte, 0, len(incompleteTimestamp))
	for i, d := range incompleteTimestamp {
		prev = append(prev, d)
		tc := testCase{
			input:        append(prefix, prev...),
			valid:        false,
			value:        nil,
			errorString:  fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, len(prefix)+i+1),
			partialValue: mex,
		}
		tCases = append(tCases, tc)
	}
	return tCases
}

// genPartialMessagesTestCases generates valid test cases
// iterating on the given data that will be put into the given part.
// It supports 4 parts - ie. hostname (0), appname (1), proc id (2), msg id (3).
func genPartialMessagesTestCases(data []byte, part int) []testCase {
	if part < 0 || part > 3 {
		panic("part not available")
	}
	templ := "<%d>%d - - - - - -"
	where := 9 + (part * 2)
	templ = templ[:where] + "%s" + templ[where+1:]

	tCases := []testCase{}
	prev := ""
	for _, c := range data {
		prev += string(c)
		randp := rand.Intn(9)
		randv := rand.Intn(9-1) + 1

		input := []byte(fmt.Sprintf(templ, randp, randv, prev))

		mex := (&SyslogMessage{}).SetVersion(uint16(randv)).SetPriority(uint8(randp)).(*SyslogMessage)
		switch part {
		case 0:
			mex.Hostname = syslogtesting.StringAddress(string(prev))
		case 1:
			mex.Appname = syslogtesting.StringAddress(string(prev))
		case 2:
			mex.ProcID = syslogtesting.StringAddress(string(prev))
		case 3:
			mex.MsgID = syslogtesting.StringAddress(string(prev))
		}

		t := testCase{
			input,
			true,
			mex,
			"",
			nil,
		}

		tCases = append(tCases, t)
	}
	return tCases
}

func init() {
	testCases = append(testCases, genIncompleteTimestampTestCases()...)
	testCases = append(testCases, genPartialMessagesTestCases(syslogtesting.MaxHostname, 0)...)
	testCases = append(testCases, genPartialMessagesTestCases(syslogtesting.MaxAppname, 1)...)
	testCases = append(testCases, genPartialMessagesTestCases(syslogtesting.MaxProcID, 2)...)
	testCases = append(testCases, genPartialMessagesTestCases(syslogtesting.MaxMsgID, 3)...)
}

func TestMachineBestEffortOption(t *testing.T) {
	p1 := NewMachine().(syslog.BestEfforter)
	assert.False(t, p1.HasBestEffort())

	p2 := NewMachine(WithBestEffort()).(syslog.BestEfforter)
	assert.True(t, p2.HasBestEffort())
}

func TestMachineParse(t *testing.T) {
	runTestCases(t, append(testCases))
}

func TestMachineParseWithCompliantMsgOn(t *testing.T) {
	runTestCases(t, append(testCases, nonCompliantMsgTestCases...), WithCompliantMsg())
}

func TestMachineParseWithCompliantMsgOff(t *testing.T) {
	// non compliant messages become valid when option is off
	for _, tc := range nonCompliantMsgTestCases {
		message, merr := NewMachine().Parse(tc.input)
		partial, perr := NewMachine(WithBestEffort()).Parse(tc.input)

		assert.NotEmpty(t, message)
		assert.Nil(t, merr)
		assert.NotEmpty(t, partial)
		assert.Nil(t, perr)
	}
}

func runTestCases(t *testing.T, tcs []testCase, machineOpts ...syslog.MachineOption) {
	t.Helper()

	for _, tc := range tcs {
		tc := tc

		t.Run(syslogtesting.RightPad(string(tc.input), 50), func(t *testing.T) {
			t.Parallel()

			message, merr := NewMachine(machineOpts...).Parse(tc.input)
			partial, perr := NewMachine(append(machineOpts, WithBestEffort())...).Parse(tc.input)

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

func TestMachineParseWithCompliantMsgWhenMessageStartsWithBOM(t *testing.T) {
	latin1 := isoLatin1String("chääs")
	subject := []byte("<1>1 - - - - - - " + BOM + latin1)

	message, merr := NewMachine(WithCompliantMsg()).Parse(subject)
	partial, perr := NewMachine(WithBestEffort(), WithCompliantMsg()).Parse(subject)

	assert.Nil(t, message)
	assert.EqualErrorf(t, merr, fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 23), "message must be valid UTF8 if starting with BOM")
	assert.Equal(t, partial, genMessageWithPartialMessage(1, 1, syslogtesting.StringAddress(BOM+"ch\xE4")))
	assert.EqualErrorf(t, perr, fmt.Sprintf(ErrMsgNotCompliant+ColumnPositionTemplate, 23), "message must be valid UTF8 if starting with BOM")
}

func isoLatin1String(s string) string {
	latin1, err := charmap.ISO8859_1.NewEncoder().String(s)
	if err != nil {
		panic(fmt.Errorf("error while encoding string as latin-1: %w", err))
	}
	if utf8.ValidString(latin1) {
		panic(fmt.Errorf("latin-1 string must not be valid utf8: %s", latin1))
	}

	return latin1
}
