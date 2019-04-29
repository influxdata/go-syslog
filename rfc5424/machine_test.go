package rfc5424

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/influxdata/go-syslog/v2"
	syslogtesting "github.com/influxdata/go-syslog/v2/common/testing"
	"github.com/stretchr/testify/assert"
)

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
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			severity: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			version:  1,
		},
	},
	// Invalid, new lines allowed only within message part
	{
		[]byte("<1>1 - \nhostname - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 7),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - host\x0Aname - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 11),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - \nan - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 9),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - a\x0An - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 10),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - \npid - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 11),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - p\x0Aid - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 12),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - \nmid -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 13),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - m\x0Aid -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 14),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
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
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  100,
		},
	},
	// Invalid, truncated after version whitespace
	{
		[]byte("<1>2 "),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 5),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  2,
		},
	},
	// Invalid, truncated after version
	{
		[]byte("<1>1"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// fixme(leodido) > when space after multi-digit version is missing, the version error handler launches (should not)
	// {
	// 	[]byte("<3>22"),
	// 	false,
	// 	nil,
	// 	fmt.Sprintf(ErrParse+ColumnPositionTemplate, 6),
	// 	(&SyslogMessage{}).SetPriority(3).SetVersion(22),
	// },
	// Invalid, non numeric (also partially) version
	{
		[]byte("<1>3a"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  3,
		},
	},
	{
		[]byte("<1>4a "),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 4),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  4,
		},
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
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  5,
		},
	},
	// Invalid, timestamp T and Z must be uppercase
	{
		[]byte(`<29>1 2006-01-02t15:04:05Z - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 16),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(3),
			severity: syslogtesting.Uint8Address(5),
			priority: syslogtesting.Uint8Address(29),
			version:  1,
		},
	},
	{
		[]byte(`<29>2 2006-01-02T15:04:05z - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 25),
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(3),
			severity: syslogtesting.Uint8Address(5),
			priority: syslogtesting.Uint8Address(29),
			version:  2,
		},
	},
	// Invalid, wrong year
	{
		[]byte("<101>123 2"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 10),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  123,
		},
	},
	{
		[]byte("<101>124 20"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 11),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  124,
		},
	},
	{
		[]byte("<101>125 201"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 12),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  125,
		},
	},
	{
		[]byte("<101>125 2013"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 13),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  125,
		},
	},
	{
		[]byte("<101>126 2013-"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 14),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  126,
		},
	},
	{
		[]byte("<101>122 201-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 12),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  122,
		},
	},
	{
		[]byte("<101>189 0-11-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 10),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  189,
		},
	},
	// Invalid, wrong month
	{
		[]byte("<101>122 2018-112-22"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  122,
		},
	},
	// Invalid, wrong day
	{
		[]byte("<101>123 2018-02-32"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  123,
		},
	},
	// Invalid, wrong hour
	{
		[]byte("<101>124 2018-02-01:25:15Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 19),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  124,
		},
	},
	// Invalid, wrong minutes
	{
		[]byte("<101>125 2003-09-29T22:99:16Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 23),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  125,
		},
	},
	// Invalid, wrong seconds
	{
		[]byte("<101>126 2003-09-29T22:09:99Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 26),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  126,
		},
	},
	// Invalid, wrong sec fraction
	{
		[]byte("<101>127 2003-09-29T22:09:01.000000000009Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 35),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  127,
		},
	},
	{
		[]byte("<101>128 2003-09-29T22:09:01.Z"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 29),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  128,
		},
	},
	{
		[]byte("<101>28 2003-09-29T22:09:01."),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 28),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  28,
		},
	},
	// Invalid, wrong time offset
	{
		[]byte("<101>129 2003-09-29T22:09:01A"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 28),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  129,
		},
	},
	{
		[]byte("<101>130 2003-08-24T05:14:15.000003-24:00"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 37),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  130,
		},
	},
	{
		[]byte("<101>131 2003-08-24T05:14:15.000003-60:00"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 36),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  131,
		},
	},
	{
		[]byte("<101>132 2003-08-24T05:14:15.000003-07:61"),
		false,
		nil,
		fmt.Sprintf(ErrTimestamp+ColumnPositionTemplate, 39),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  132,
		},
	},
	{
		[]byte(`<29>1 2006-01-02T15:04:05Z+07:00 - - - - -`),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 26), // after the Z (valid and complete timestamp) it searches for a whitespace
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2006-01-02T15:04:05Z"),
		},
	},
	// Invalid, non existing dates
	{
		[]byte("<101>11 2003-09-31T22:14:15.003Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:15.003Z\": day out of range [col 32]",
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  11,
		},
	},
	{
		[]byte("<101>12 2003-09-31T22:14:16Z"),
		false,
		nil,
		"parsing time \"2003-09-31T22:14:16Z\": day out of range [col 28]",
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  12,
		},
	},
	{
		[]byte("<101>12 2018-02-29T22:14:16+01:00"),
		false,
		nil,
		"parsing time \"2018-02-29T22:14:16+01:00\": day out of range [col 33]",
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(101),
			facility: syslogtesting.Uint8Address(12),
			severity: syslogtesting.Uint8Address(5),
			version:  12,
		},
	},
	// Invalid, hostname too long
	{
		[]byte("<1>1 - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 262),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 281),
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(1),
			facility:  syslogtesting.Uint8Address(0),
			severity:  syslogtesting.Uint8Address(1),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-09-29T22:14:16Z"),
		},
	},
	// Invalid, appname too long
	{
		[]byte("<1>1 - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 57),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 76),
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(1),
			facility:  syslogtesting.Uint8Address(0),
			severity:  syslogtesting.Uint8Address(1),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-09-29T22:14:16Z"),
		},
	},
	{
		[]byte("<1>1 - host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 60),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			hostname: syslogtesting.StringAddress("host"),
		},
	},
	{
		[]byte("<1>1 2003-09-29T22:14:16Z host abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefX - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 79),
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(1),
			facility:  syslogtesting.Uint8Address(0),
			severity:  syslogtesting.Uint8Address(1),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-09-29T22:14:16Z"),
			hostname:  syslogtesting.StringAddress("host"),
		},
	},
	// Invalid, procid too long
	{
		[]byte("<1>1 - - - abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabX - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 139),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Invalid, msgid too long
	{
		[]byte("<1>1 - - - - abcdefghilmnopqrstuvzabcdefghilmX -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 45),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Not print US-ASCII chars for hostname, appname, procid, and msgid
	{
		[]byte("<1>1 -   - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrHostname+ColumnPositionTemplate, 7),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - -   - - -"),
		false,
		nil,
		fmt.Sprintf(ErrAppname+ColumnPositionTemplate, 9),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - -   - -"),
		false,
		nil,
		fmt.Sprintf(ErrProcID+ColumnPositionTemplate, 11),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - -   -"),
		false,
		nil,
		fmt.Sprintf(ErrMsgID+ColumnPositionTemplate, 13),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Invalid, with malformed structured data
	{
		[]byte("<1>1 - - - - - X"),
		false,
		nil,
		fmt.Sprintf(ErrStructuredData+ColumnPositionTemplate, 15),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Invalid, with empty structured data
	{
		[]byte("<1>1 - - - - - []"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, with structured data id containing space
	{
		[]byte("<1>1 - - - - - [ ]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, with structured data id containing =
	{
		[]byte("<1>1 - - - - - [=]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, with structured data id containing ]
	{
		[]byte("<1>1 - - - - - []]"),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, with structured data id containing "
	{
		[]byte(`<1>1 - - - - - ["]`),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 16),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, too long structured data id
	{
		[]byte(`<1>1 - - - - - [abcdefghilmnopqrstuvzabcdefghilmX]`),
		false,
		nil,
		fmt.Sprintf(ErrSdID+ColumnPositionTemplate, 48),
		&SyslogMessage{
			priority:       syslogtesting.Uint8Address(1),
			facility:       syslogtesting.Uint8Address(0),
			severity:       syslogtesting.Uint8Address(1),
			version:        1,
			structuredData: nil,
		},
	},
	// Invalid, too long structured data param key
	{
		[]byte(`<1>1 - - - - - [id abcdefghilmnopqrstuvzabcdefghilmX="val"]`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			structuredData: &map[string]map[string]string{
				"id": {},
			},
		},
	},
	// Valid, minimal
	{
		[]byte("<1>1 - - - - - -"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
		},
		"",
		nil,
	},
	{
		[]byte("<0>1 - - - - - -"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(0),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(0),
			version:  1,
		},
		"",
		nil,
	},
	// Valid, average message
	{
		[]byte(`<29>1 2016-02-21T04:32:57+00:00 web1 someservice - - [origin x-service="someservice"][meta sequenceId="14125553"] 127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-02-21T04:32:57+00:00"),
			hostname:  syslogtesting.StringAddress("web1"),
			appname:   syslogtesting.StringAddress("someservice"),
			structuredData: &map[string]map[string]string{
				"origin": {
					"x-service": "someservice",
				},
				"meta": {
					"sequenceId": "14125553",
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1456029177 "GET /v1/ok HTTP/1.1" 200 145 "-" "hacheck 0.9.0" 24306 127.0.0.1:40124 575`),
		},
		"",
		nil,
	},
	// Valid, hostname, appname, procid, msgid can contain dashes
	{
		[]byte("<1>100 - host-name - - - -"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  100,
			hostname: syslogtesting.StringAddress("host-name"),
		},
		"",
		nil,
	},
	{
		[]byte("<1>101 - host-name app-name - - -"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  101,
			hostname: syslogtesting.StringAddress("host-name"),
			appname:  syslogtesting.StringAddress("app-name"),
		},
		"",
		nil,
	},
	{
		[]byte("<1>102 - host-name app-name proc-id - -"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  102,
			hostname: syslogtesting.StringAddress("host-name"),
			appname:  syslogtesting.StringAddress("app-name"),
			procID:   syslogtesting.StringAddress("proc-id"),
		},
		"",
		nil,
	},
	{
		[]byte("<1>103 - host-name app-name proc-id msg-id -"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  103,
			hostname: syslogtesting.StringAddress("host-name"),
			appname:  syslogtesting.StringAddress("app-name"),
			procID:   syslogtesting.StringAddress("proc-id"),
			msgID:    syslogtesting.StringAddress("msg-id"),
		},
		"",
		nil,
	},
	// Valid, w/0 structured data and w/o message, with other fields all max length
	{
		[]byte("<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm -"),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(191),
			facility:  syslogtesting.Uint8Address(23),
			severity:  syslogtesting.Uint8Address(7),
			version:   999,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2018-12-31T23:59:59.999999-23:59"),
			hostname:  syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc"),
			appname:   syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef"),
			procID:    syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab"),
			msgID:     syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilm"),
		},
		"",
		nil,
	},
	// Valid, all fields max length, with structured data and message
	{
		[]byte(`<191>999 2018-12-31T23:59:59.999999-23:59 abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab abcdefghilmnopqrstuvzabcdefghilm [an@id key1="val1" key2="val2"][another@id key1="val1"] Some message "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(191),
			facility:  syslogtesting.Uint8Address(23),
			severity:  syslogtesting.Uint8Address(7),
			version:   999,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2018-12-31T23:59:59.999999-23:59"),
			hostname:  syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc"),
			appname:   syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef"),
			procID:    syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab"),
			msgID:     syslogtesting.StringAddress("abcdefghilmnopqrstuvzabcdefghilm"),
			structuredData: &map[string]map[string]string{
				"an@id": {
					"key1": "val1",
					"key2": "val2",
				},
				"another@id": {
					"key1": "val1",
				},
			},
			message: syslogtesting.StringAddress(`Some message "GET"`),
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/0 procid
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       syslogtesting.Uint8Address(4),
			severity:       syslogtesting.Uint8Address(2),
			priority:       syslogtesting.Uint8Address(34),
			version:        1,
			timestamp:      syslogtesting.TimeParse(RFC3339MICRO, "2003-10-11T22:14:15.003Z"),
			hostname:       syslogtesting.StringAddress("mymachine.example.com"),
			appname:        syslogtesting.StringAddress("su"),
			procID:         nil,
			msgID:          syslogtesting.StringAddress("ID47"),
			structuredData: nil,
			message:        syslogtesting.StringAddress("BOM'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/o timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			facility:       syslogtesting.Uint8Address(23),
			severity:       syslogtesting.Uint8Address(3),
			priority:       syslogtesting.Uint8Address(187),
			version:        222,
			timestamp:      nil,
			hostname:       syslogtesting.StringAddress("mymachine.example.com"),
			appname:        syslogtesting.StringAddress("su"),
			procID:         nil,
			msgID:          syslogtesting.StringAddress("ID47"),
			structuredData: nil,
			message:        syslogtesting.StringAddress("'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/o msgid
	{
		[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% Time to make the do-nuts."),
		true,
		&SyslogMessage{
			facility:       syslogtesting.Uint8Address(20),
			severity:       syslogtesting.Uint8Address(5),
			priority:       syslogtesting.Uint8Address(165),
			version:        1,
			timestamp:      syslogtesting.TimeParse(RFC3339MICRO, "2003-08-24T05:14:15.000003-07:00"),
			hostname:       syslogtesting.StringAddress("192.0.2.1"),
			appname:        syslogtesting.StringAddress("myproc"),
			procID:         syslogtesting.StringAddress("8710"),
			msgID:          nil,
			structuredData: nil,
			message:        syslogtesting.StringAddress("%% Time to make the do-nuts."),
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, w/o msg
	{
		[]byte("<165>2 2003-08-24T05:14:15.000003-07:00 - - - - -"),
		true,
		&SyslogMessage{
			facility:       syslogtesting.Uint8Address(20),
			severity:       syslogtesting.Uint8Address(5),
			priority:       syslogtesting.Uint8Address(165),
			version:        2,
			timestamp:      syslogtesting.TimeParse(RFC3339MICRO, "2003-08-24T05:14:15.000003-07:00"),
			hostname:       nil,
			appname:        nil,
			procID:         nil,
			msgID:          nil,
			structuredData: nil,
			message:        nil,
		},
		"",
		nil,
	},
	// Valid, w/o structure data, w/o hostname, w/o appname, w/o procid, w/o msgid, empty msg
	{
		[]byte("<165>222 2003-08-24T05:14:15.000003-07:00 - - - - - "),
		true,
		&SyslogMessage{
			facility:       syslogtesting.Uint8Address(20),
			severity:       syslogtesting.Uint8Address(5),
			priority:       syslogtesting.Uint8Address(165),
			version:        222,
			timestamp:      syslogtesting.TimeParse(RFC3339MICRO, "2003-08-24T05:14:15.000003-07:00"),
			hostname:       nil,
			appname:        nil,
			procID:         nil,
			msgID:          nil,
			structuredData: nil,
			message:        nil,
		},
		"",
		nil,
	},
	// Valid, with structured data is, w/o structured data params
	{
		[]byte("<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid] some_message"),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(9),
			severity:  syslogtesting.Uint8Address(6),
			priority:  syslogtesting.Uint8Address(78),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T00:04:01+00:00"),
			hostname:  syslogtesting.StringAddress("host1"),
			appname:   syslogtesting.StringAddress("CROND"),
			procID:    syslogtesting.StringAddress("10391"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"sdid": {},
			},
			message: syslogtesting.StringAddress("some_message"),
		},
		"",
		nil,
	},
	// Valid, with structured data id, with structured data params
	{
		[]byte(`<78>1 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="⌘"] some_message`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(9),
			severity:  syslogtesting.Uint8Address(6),
			priority:  syslogtesting.Uint8Address(78),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T00:04:01+00:00"),
			hostname:  syslogtesting.StringAddress("host1"),
			appname:   syslogtesting.StringAddress("CROND"),
			procID:    syslogtesting.StringAddress("10391"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"sdid": {
					"x": "⌘",
				},
			},
			message: syslogtesting.StringAddress("some_message"),
		},
		"",
		nil,
	},
	// Valid, with structured data is, with structured data params
	{
		[]byte(`<78>2 2016-01-15T00:04:01+00:00 host1 CROND 10391 - [sdid x="hey \\u2318 hey"] some_message`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(9),
			severity:  syslogtesting.Uint8Address(6),
			priority:  syslogtesting.Uint8Address(78),
			version:   2,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T00:04:01+00:00"),
			hostname:  syslogtesting.StringAddress("host1"),
			appname:   syslogtesting.StringAddress("CROND"),
			procID:    syslogtesting.StringAddress("10391"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"sdid": {
					"x": `hey \u2318 hey`,
				},
			},
			message: syslogtesting.StringAddress("some_message"),
		},
		"",
		nil,
	},
	// Valid, with (escaped) backslash within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta es="\\valid"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   50,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"es": `\valid`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>52 2016-01-15T01:00:43Z hn S - - [meta one="\\one" two="\\two"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   52,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"one": `\one`,
					"two": `\two`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>53 2016-01-15T01:00:43Z hn S - - [meta one="\\one"][other two="\\two" double="\\a\\b"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   53,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"one": `\one`,
				},
				"other": {
					"two":    `\two`,
					"double": `\a\b`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>51 2016-01-15T01:00:43Z hn S - - [meta es="\\double\\slash"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   51,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"es": `\double\slash`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>54 2016-01-15T01:00:43Z hn S - - [meta es="in \\middle of the string"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   54,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"es": `in \middle of the string`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>55 2016-01-15T01:00:43Z hn S - - [meta es="at the \\end"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   55,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"es": `at the \end`,
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	// Valid, with control characters within structured data param value
	{
		[]byte("<29>50 2016-01-15T01:00:43Z hn S - - [meta es=\"\t5Ὂg̀9!℃ᾭGa b\"] 127.0.0.1 - - 1452819643 \"GET\""),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   50,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"es": "\t5Ὂg̀9!℃ᾭGa b",
				},
			},
			message: syslogtesting.StringAddress(`127.0.0.1 - - 1452819643 "GET"`),
		},
		"",
		nil,
	},
	// Valid, with utf8 within structured data param value
	{
		[]byte(`<29>50 2016-01-15T01:00:43Z hn S - - [meta gr="κόσμε" es="ñ"][beta pr="₡"] 𐌼 "GET"`),
		true,
		&SyslogMessage{
			priority:  syslogtesting.Uint8Address(29),
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			version:   50,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {
					"gr": "κόσμε",
					"es": "ñ",
				},
				"beta": {
					"pr": "₡",
				},
			},
			message: syslogtesting.StringAddress(`𐌼 "GET"`),
		},
		"",
		nil,
	},
	// Valid, with structured data, w/o msg
	{
		[]byte("<165>3 2003-10-11T22:14:15.003Z example.com evnts - ID27 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(20),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(165),
			version:   3,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-10-11T22:14:15.003Z"),
			hostname:  syslogtesting.StringAddress("example.com"),
			appname:   syslogtesting.StringAddress("evnts"),
			procID:    nil,
			msgID:     syslogtesting.StringAddress("ID27"),
			structuredData: &map[string]map[string]string{
				"exampleSDID@32473": {
					"iut":         "3",
					"eventSource": "Application",
					"eventID":     "1011",
				},
				"examplePriority@32473": {
					"class": "high",
				},
			},
			message: nil,
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
			priority:  syslogtesting.Uint8Address(165),
			facility:  syslogtesting.Uint8Address(20),
			severity:  syslogtesting.Uint8Address(5),
			version:   3,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-10-11T22:14:15.003Z"),
			hostname:  syslogtesting.StringAddress("example.com"),
			appname:   syslogtesting.StringAddress("evnts"),
			msgID:     syslogtesting.StringAddress("ID27"),
			structuredData: &map[string]map[string]string{
				"id1": {},
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
			priority:  syslogtesting.Uint8Address(165),
			facility:  syslogtesting.Uint8Address(20),
			severity:  syslogtesting.Uint8Address(5),
			version:   3,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-10-11T22:14:15.003Z"),
			hostname:  syslogtesting.StringAddress("example.com"),
			appname:   syslogtesting.StringAddress("evnts"),
			msgID:     syslogtesting.StringAddress("ID27"),
			structuredData: &map[string]map[string]string{
				"id1": {},
				"dupe": {
					"e": "1",
				},
			},
		},
	},
	// Valid, with structured data w/o msg
	{
		[]byte(`<165>4 2003-10-11T22:14:15.003Z mymachine.it e - 1 [ex@32473 iut="3" eventSource="A"] An application event log entry...`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(20),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(165),
			version:   4,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2003-10-11T22:14:15.003Z"),
			hostname:  syslogtesting.StringAddress("mymachine.it"),
			appname:   syslogtesting.StringAddress("e"),
			procID:    nil,
			msgID:     syslogtesting.StringAddress("1"),
			structuredData: &map[string]map[string]string{
				"ex@32473": {
					"iut":         "3",
					"eventSource": "A",
				},
			},
			message: syslogtesting.StringAddress("An application event log entry..."),
		},
		"",
		nil,
	},
	// Valid, with double quotes in the message
	{
		[]byte(`<29>1 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [origin x-service="svcname"][meta sequenceId="1"] 127.0.0.1 - - 1452819643 "GET"`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   1,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("some-host-name"),
			appname:   syslogtesting.StringAddress("SEKRETPROGRAM"),
			procID:    syslogtesting.StringAddress("prg"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"origin": {
					"x-service": "svcname",
				},
				"meta": {
					"sequenceId": "1",
				},
			},
			message: syslogtesting.StringAddress("127.0.0.1 - - 1452819643 \"GET\""),
		},
		"",
		nil,
	},
	// Valid, with empty structured data param value
	{
		[]byte(`<1>1 - - - - - [id pk=""]`),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			structuredData: &map[string]map[string]string{
				"id": {
					"pk": "",
				},
			},
		},
		"",
		nil,
	},
	// Valid, with escaped character within param value
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]"] some "mex"`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   2,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("some-host-name"),
			appname:   syslogtesting.StringAddress("SEKRETPROGRAM"),
			procID:    syslogtesting.StringAddress("prg"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"meta": {
					"escape": "]",
				},
			},
			message: syslogtesting.StringAddress(`some "mex"`),
		},
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\\"]`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   2,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("some-host-name"),
			appname:   syslogtesting.StringAddress("SEKRETPROGRAM"),
			procID:    syslogtesting.StringAddress("prg"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"meta": {
					"escape": `\`,
				},
			},
		},
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\""]`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   2,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("some-host-name"),
			appname:   syslogtesting.StringAddress("SEKRETPROGRAM"),
			procID:    syslogtesting.StringAddress("prg"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"meta": {
					"escape": `"`,
				},
			},
		},
		"",
		nil,
	},
	{
		[]byte(`<29>2 2016-01-15T01:00:43Z some-host-name SEKRETPROGRAM prg - [meta escape="\]\"\\\\\]\""]`),
		true,
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   2,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("some-host-name"),
			appname:   syslogtesting.StringAddress("SEKRETPROGRAM"),
			procID:    syslogtesting.StringAddress("prg"),
			msgID:     nil,
			structuredData: &map[string]map[string]string{
				"meta": {
					"escape": `]"\\]"`,
				},
			},
		},
		"",
		nil,
	},
	// Invalid, param value can not contain closing square bracket - ie., ]
	{
		[]byte(`<29>3 2016-01-15T01:00:43Z hn S - - [meta escape="]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 50),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   3,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>5 2016-01-15T01:00:43Z hn S - - [meta escape="]q"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 50),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   5,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>4 2016-01-15T01:00:43Z hn S - - [meta escape="p]"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 51),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   4,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	// Invalid, param value can not contain doublequote char - ie., ""
	{
		[]byte(`<29>4 2017-01-15T01:00:43Z hn S - - [meta escape="""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   4,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2017-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>6 2016-01-15T01:00:43Z hn S - - [meta escape="a""] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 52),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   6,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>4 2018-01-15T01:00:43Z hn S - - [meta escape=""b"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrSdParam+ColumnPositionTemplate, 51),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   4,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2018-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	// Invalid, param value can not contain backslash - ie., \
	{
		[]byte(`<29>5 2019-01-15T01:00:43Z hn S - - [meta escape="\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 52),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   5,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2019-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>7 2019-01-15T01:00:43Z hn S - - [meta escape="a\"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 53),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   7,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2019-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	{
		[]byte(`<29>8 2016-01-15T01:00:43Z hn S - - [meta escape="\n"] 127.0.0.1 - - 1452819643 "GET"`),
		false,
		nil,
		fmt.Sprintf(ErrEscape+ColumnPositionTemplate, 51),
		&SyslogMessage{
			facility:  syslogtesting.Uint8Address(3),
			severity:  syslogtesting.Uint8Address(5),
			priority:  syslogtesting.Uint8Address(29),
			version:   8,
			timestamp: syslogtesting.TimeParse(RFC3339MICRO, "2016-01-15T01:00:43Z"),
			hostname:  syslogtesting.StringAddress("hn"),
			appname:   syslogtesting.StringAddress("S"),
			structuredData: &map[string]map[string]string{
				"meta": {},
			},
		},
	},
	// Valid, message starting with byte order mark (BOM, \uFEFF)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\ufeff"),
		},
		"",
		nil,
	},
	// Valid, greek
	{
		[]byte("<1>1 - - - - - - κόσμε"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("κόσμε"),
		},
		"",
		nil,
	},
	// Valid, 2 octet sequence
	{
		[]byte("<1>1 - - - - - - "),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress(""),
		},
		"",
		nil,
	},
	// Valid, spanish (2 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xc3\xb1"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("ñ"),
		},
		"",
		nil,
	},
	// Valid, colon currency sign (3 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xe2\x82\xa1"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("₡"),
		},
		"",
		nil,
	},
	// Valid, gothic letter (4 octet sequence)
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF \xf0\x90\x8c\xbc"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\ufeff 𐌼"),
		},
		"",
		nil,
	},
	// Valid, 5 octet sequence
	{
		[]byte("<1>1 - - - - - - \xC8\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("Ȁ000"),
		},
		"",
		nil,
	},
	// Valid, 6 octet sequence
	{
		[]byte("<1>1 - - - - - - \xE4\x80\x80\x30\x30\x30"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("䀀000"),
		},
		"",
		nil,
	},
	// Valid, UTF-8 boundary conditions
	{
		[]byte("<1>1 - - - - - - \xC4\x90\x30\x30\x30"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("Đ000"),
		},
		"",
		nil,
	},
	{
		[]byte("<1>1 - - - - - - \x0D\x37\x46\x46"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\r7FF"),
		},
		"",
		nil,
	},
	// Valid, Tamil poetry of Subramaniya Bharathiyar
	{
		[]byte("<1>1 - - - - - - யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("யாமறிந்த மொழிகளிலே தமிழ்மொழி போல் இனிதாவது எங்கும் காணோம், பாமரராய் விலங்குகளாய், உலகனைத்தும் இகழ்ச்சிசொலப் பான்மை கெட்டு, நாமமது தமிழரெனக் கொண்டு இங்கு வாழ்ந்திடுதல் நன்றோ? சொல்லீர்! தேமதுரத் தமிழோசை உலகமெலாம் பரவும்வகை செய்தல் வேண்டும்."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Milanese)
	{
		[]byte("<1>1 - - - - - - Sôn bôn de magnà el véder, el me fa minga mal."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("Sôn bôn de magnà el véder, el me fa minga mal."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Romano)
	{
		[]byte("<1>1 - - - - - - Me posso magna' er vetro, e nun me fa male."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("Me posso magna' er vetro, e nun me fa male."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Braille)
	{
		[]byte("<1>1 - - - - - - ⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("⠊⠀⠉⠁⠝⠀⠑⠁⠞⠀⠛⠇⠁⠎⠎⠀⠁⠝⠙⠀⠊⠞⠀⠙⠕⠑⠎⠝⠞⠀⠓⠥⠗⠞⠀⠍⠑"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Sanskrit)
	{
		[]byte("<1>1 - - - - - - काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("काचं शक्नोम्यत्तुम् । नोपहिनस्ति माम् ॥"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Urdu)
	{
		[]byte("<1>1 - - - - - - میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("میں کانچ کھا سکتا ہوں اور مجھے تکلیف نہیں ہوتی ۔"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Yiddish)
	{
		[]byte("<1>1 - - - - - - איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("איך קען עסן גלאָז און עס טוט מיר נישט װײ."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Polish)
	{
		[]byte("<1>1 - - - - - - Mogę jeść szkło, i mi nie szkodzi."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("Mogę jeść szkło, i mi nie szkodzi."),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Japanese)
	{
		[]byte("<1>1 - - - - - - 私はガラスを食べられます。それは私を傷つけません。"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("私はガラスを食べられます。それは私を傷つけません。"),
		},
		"",
		nil,
	},
	// Valid, I Can Eat Glass (Arabic)
	{
		[]byte("<1>1 - - - - - - أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("أنا قادر على أكل الزجاج و هذا لا يؤلمني."),
		},
		"",
		nil,
	},
	// Valid, russian alphabet
	{
		[]byte("<1>1 - - - - - - абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
		},
		"",
		nil,
	},
	// Valid, armenian letters
	{
		[]byte("<1>1 - - - - - - ԰ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ՗՘ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆևֈ։֊֋֌֍֎֏"),
		true,
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\u0530ԱԲԳԴԵԶԷԸԹԺԻԼԽԾԿՀՁՂՃՄՅՆՇՈՉՊՋՌՍՎՏՐՑՒՓՔՕՖ\u0557\u0558ՙ՚՛՜՝՞՟աբգդեզէըթիլխծկհձղճմյնշոչպջռսվտրցւփքօֆև\u0588։֊\u058b\u058c֍֎֏"),
		},
		"",
		nil,
	},
	// Valid, new line within message
	{
		[]byte("<1>1 - - - - - - x\x0Ay"),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("x\ny"),
		},
		"",
		nil,
	},
	{
		[]byte(`<1>2 - - - - - - x
y`),
		true,
		&SyslogMessage{
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			priority: syslogtesting.Uint8Address(1),
			version:  2,
			message:  syslogtesting.StringAddress("x\ny"),
		},
		"",
		nil,
	},
	// Invalid, out of range code within message
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xC1"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 20),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF"),
		},
	},
	{
		[]byte("<1>2 - - - - - - \xC1"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  2,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 21),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xc3"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc3\x28"), // invalid 2 octet sequence
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xc3"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 20),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xa0\xa1"), // invalid sequence identifier
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 21),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xe2"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x28\xa1"), // invalid 3 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xe2"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 22),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xe2\x82"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe2\x82\x28"), // invalid 3 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 19),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xe2\x82"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 21),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\xbc"), // invalid 4 octet sequence (2nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 22),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xf0\x90"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x90\x28\xbc"), // invalid 4 octet sequence (3nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 19),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xf0\x90"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xEF\xBB\xBF\xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 21),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xEF\xBB\xBF\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x28\x8c\x28"), // invalid 4 octet sequence (4nd octet)
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xf0"),
		},
	},
	// Invalid, impossible bytes
	{
		[]byte("<1>1 - - - - - - \xfe\xfe\xff\xff"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xfe"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xff"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Invalid, overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x80\x80\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x80\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x80\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x80\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xe0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc0\xaf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	// Invalid, maximum overlong sequences
	{
		[]byte("<1>1 - - - - - - \xfc\x83\xbf\xbf\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf8\x87\xbf\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
		},
	},
	{
		[]byte("<1>1 - - - - - - \xf0\x8f\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xf0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xe0\x9f\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xe0"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xc1\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 17),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  nil,
		},
	},
	// Invalid, illegal code positions, single utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xad\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xae\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xaf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xb0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbe\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	{
		[]byte("<1>1 - - - - - - \xed\xbf\xbf"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	// Invalid, illegal code positions, paired utf-16 surrogates
	{
		[]byte("<1>1 - - - - - - \xed\xa0\x80\xed\xb0\x80"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 18),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("\xed"),
		},
	},
	// Invalid, out of range code within message after valid string
	{
		[]byte("<1>1 - - - - - - valid\xEF\xBB\xBF\xC1"),
		false,
		nil,
		fmt.Sprintf(ErrMsg+ColumnPositionTemplate, 25),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  1,
			message:  syslogtesting.StringAddress("valid\ufeff"),
		},
	},
	// Invalid, missing whitespace after nil timestamp
	{
		[]byte("<1>10 -- - - - -"),
		false,
		nil,
		fmt.Sprintf(ErrParse+ColumnPositionTemplate, 7),
		&SyslogMessage{
			priority: syslogtesting.Uint8Address(1),
			facility: syslogtesting.Uint8Address(0),
			severity: syslogtesting.Uint8Address(1),
			version:  10,
		},
	},

	// (fixme) > evaluate non characters for UTF-8 security concerns, eg. \xef\xbf\xbe
}

// genIncompleteTimestampTestCases generates test cases with incomplete timestamp part.
func genIncompleteTimestampTestCases() []testCase {
	incompleteTimestamp := []byte("2003-11-02T23:12:46.012345")
	prefix := []byte("<1>1 ")
	mex := &SyslogMessage{
		priority: syslogtesting.Uint8Address(1),
		severity: syslogtesting.Uint8Address(1),
		facility: syslogtesting.Uint8Address(0),
		version:  1,
	}
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

		mex := &SyslogMessage{
			priority: syslogtesting.Uint8Address(uint8(randp)),
			severity: syslogtesting.Uint8Address(uint8(randp % 8)),
			facility: syslogtesting.Uint8Address(uint8(randp / 8)),
			version:  uint16(randv),
		}
		switch part {
		case 0:
			mex.hostname = syslogtesting.StringAddress(string(prev))
		case 1:
			mex.appname = syslogtesting.StringAddress(string(prev))
		case 2:
			mex.procID = syslogtesting.StringAddress(string(prev))
		case 3:
			mex.msgID = syslogtesting.StringAddress(string(prev))
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
	for _, tc := range testCases {
		tc := tc
		t.Run(syslogtesting.RightPad(string(tc.input), 50), func(t *testing.T) {
			t.Parallel()

			message, merr := NewMachine().Parse(tc.input)
			partial, perr := NewMachine(WithBestEffort()).Parse(tc.input)

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
