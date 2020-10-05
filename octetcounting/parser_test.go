package octetcounting

import (
	"fmt"
	"strings"
	"testing"

	"github.com/influxdata/go-syslog/v3"
	"github.com/influxdata/go-syslog/v3/rfc5424"
	syslogtesting "github.com/influxdata/go-syslog/v3/testing"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	descr             string
	input             string
	results           []syslog.Result
	bestEffortResults []syslog.Result
	maxMessageLength  int
}

var testCases []testCase

func getTimestampError(col int) error {
	return fmt.Errorf(rfc5424.ErrTimestamp+rfc5424.ColumnPositionTemplate, col)
}

func getParsingError(col int) error {
	return fmt.Errorf(rfc5424.ErrParse+rfc5424.ColumnPositionTemplate, col)
}

func getTestCases() []testCase {
	return []testCase{
		{
			descr: "empty",
			input: "",
			results: []syslog.Result{
				{Error: fmt.Errorf("found %s, expecting a %s", EOF, MSGLEN)},
			},
			bestEffortResults: []syslog.Result{
				{Error: fmt.Errorf("found %s, expecting a %s", EOF, MSGLEN)},
			},
		},
		{
			descr: "1st ok/2nd mf", // mf means malformed syslog message
			input: "16 <1>1 - - - - - -17 <2>12 A B C D E -",
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Error: getTimestampError(6),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(2).SetVersion(12),
					Error:   getTimestampError(6),
				},
			},
		},
		{
			descr: "1st ok/2nd ko", // ko means wrong token
			input: "16 <1>1 - - - - - -xaaa",
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Error: fmt.Errorf("found %s, expecting a %s", Token{ILLEGAL, []byte("x")}, MSGLEN),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Error: fmt.Errorf("found %s, expecting a %s", Token{ILLEGAL, []byte("x")}, MSGLEN),
				},
			},
		},
		{
			descr: "1st ml/2nd ko",
			input: "16 <1>1 A B C D E -xaaa",
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: getTimestampError(5),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
					Error:   getTimestampError(5),
				},
				{
					Error: fmt.Errorf("found %s, expecting a %s", Token{ILLEGAL, []byte("x")}, MSGLEN),
				},
			},
		},
		{
			descr: "1st ok//utf8",
			input: "23 <1>1 - - - - - - hellø", // msglen MUST be the octet count
			//results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(1).
						SetVersion(1).
						SetMessage("hellø"),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(1).
						SetVersion(1).
						SetMessage("hellø"),
				},
			},
		},
		{
			descr: "1st ko//incomplete SYSLOGMSG",
			input: "16 <1>1",
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, EOF, "<1>1", SYSLOGMSG, 16),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
					Error:   getParsingError(4),
					// Error:   fmt.Errorf(`found %s after "%s", expecting a %s containing %d octets`, EOF, "<1>1", SYSLOGMSG, 16),
				},
			},
		},
		{
			descr: "1st ko//missing WS found ILLEGAL",
			input: "16<1>1",
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: fmt.Errorf("found %s, expecting a %s", Token{ILLEGAL, []byte("<")}, WS),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Error: fmt.Errorf("found %s, expecting a %s", Token{ILLEGAL, []byte("<")}, WS),
				},
			},
		},
		{
			descr: "1st ko//missing WS found EOF",
			input: "1",
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: fmt.Errorf("found %s, expecting a %s", EOF, WS),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Error: fmt.Errorf("found %s, expecting a %s", EOF, WS),
				},
			},
		},
		{
			descr: "1st ok/2nd ok/3rd ok",
			input: "48 <1>1 2003-10-11T22:14:15.003Z host.local - - - -25 <3>1 - host.local - - - -38 <2>1 - host.local su - - - κόσμε",
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(1).
						SetVersion(1).
						SetTimestamp("2003-10-11T22:14:15.003Z").
						SetHostname("host.local"),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(3).
						SetVersion(1).
						SetHostname("host.local"),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(2).
						SetVersion(1).
						SetHostname("host.local").
						SetAppname("su").
						SetMessage("κόσμε"),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(1).
						SetVersion(1).
						SetTimestamp("2003-10-11T22:14:15.003Z").
						SetHostname("host.local"),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(3).
						SetVersion(1).
						SetHostname("host.local"),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(2).
						SetVersion(1).
						SetHostname("host.local").
						SetAppname("su").
						SetMessage("κόσμε"),
				},
			},
		},
		{
			descr: "1st ok/2nd mf/3rd ok", // mf means malformed syslog message
			input: "16 <1>1 - - - - - -17 <2>12 A B C D E -16 <1>1 - - - - - -",
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Error: getTimestampError(6),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(2).SetVersion(12),
					Error:   getTimestampError(6),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
			},
		},
		{
			descr: "1st ok//max",
			input: fmt.Sprintf(
				"8192 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
		},
		{
			descr: "1st ok//longer-max",
			input: fmt.Sprintf(
				"65529 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.LongerMaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
			},
			maxMessageLength: 65529,
		},
		{
			descr: "1st ok/2nd ok//max/max",
			input: fmt.Sprintf(
				"8192 <%d>%d %s %s %s %s %s - %s8192 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
		},
		{
			descr: "1st ok/2nd ok//longer-max/longer-max",
			input: fmt.Sprintf(
				"65529 <%d>%d %s %s %s %s %s - %s65529 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.LongerMaxMessage),
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.LongerMaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.LongerMaxMessage)),
				},
			},
			maxMessageLength: 65529,
		},
		{
			descr: "1st ok/2nd ok/3rd ok//max/no/max",
			input: fmt.Sprintf(
				"8192 <%d>%d %s %s %s %s %s - %s16 <1>1 - - - - - -8192 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(1),
				},
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
				},
			},
		},
		{
			descr: "1st ml//maxlen gt 8192", // maxlength greather than the buffer size
			input: fmt.Sprintf(
				"8193 <%d>%d %s %s %s %s %s - %s",
				syslogtesting.MaxPriority,
				syslogtesting.MaxVersion,
				syslogtesting.MaxRFC3339MicroTimestamp,
				string(syslogtesting.MaxHostname),
				string(syslogtesting.MaxAppname),
				string(syslogtesting.MaxProcID),
				string(syslogtesting.MaxMsgID),
				string(syslogtesting.MaxMessage),
			),
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: fmt.Errorf(
						"found %s after \"%s\", expecting a %s containing %d octets",
						EOF,
						fmt.Sprintf(
							"<%d>%d %s %s %s %s %s - %s", syslogtesting.MaxPriority,
							syslogtesting.MaxVersion,
							syslogtesting.MaxRFC3339MicroTimestamp,
							string(syslogtesting.MaxHostname),
							string(syslogtesting.MaxAppname),
							string(syslogtesting.MaxProcID),
							string(syslogtesting.MaxMsgID),
							string(syslogtesting.MaxMessage),
						),
						SYSLOGMSG,
						8193,
					),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).
						SetPriority(syslogtesting.MaxPriority).
						SetVersion(syslogtesting.MaxVersion).
						SetTimestamp(syslogtesting.MaxRFC3339MicroTimestamp).
						SetHostname(string(syslogtesting.MaxHostname)).
						SetAppname(string(syslogtesting.MaxAppname)).
						SetProcID(string(syslogtesting.MaxProcID)).
						SetMsgID(string(syslogtesting.MaxMsgID)).
						SetMessage(string(syslogtesting.MaxMessage)),
					Error: fmt.Errorf(
						"found %s after \"%s\", expecting a %s containing %d octets",
						EOF,
						fmt.Sprintf(
							"<%d>%d %s %s %s %s %s - %s", syslogtesting.MaxPriority,
							syslogtesting.MaxVersion,
							syslogtesting.MaxRFC3339MicroTimestamp,
							string(syslogtesting.MaxHostname),
							string(syslogtesting.MaxAppname),
							string(syslogtesting.MaxProcID),
							string(syslogtesting.MaxMsgID),
							string(syslogtesting.MaxMessage),
						),
						SYSLOGMSG,
						8193,
					),
				},
			},
		},
		{
			descr: "1st uf/2nd ok//incomplete SYSLOGMSG/notdetectable",
			input: "16 <1>217 <11>1 - - - - - -",
			// results w/o best effort
			results: []syslog.Result{
				{
					Error: getTimestampError(7),
				},
			},
			// results with best effort
			bestEffortResults: []syslog.Result{
				{
					Message: (&rfc5424.SyslogMessage{}).SetPriority(1).SetVersion(217),
					Error:   getTimestampError(7),
				},
				{
					Error: fmt.Errorf("found %s, expecting a %s", WS, MSGLEN),
				},
			},
		},
	}
}

func init() {
	testCases = getTestCases()
}

func TestParse(t *testing.T) {
	for i := range testCases {
		tc := testCases[i] // tests could be running in parallel, needs to be scoped.
		if tc.maxMessageLength == 0 {
			tc.maxMessageLength = 8192
		}
		t.Run(fmt.Sprintf("strict/%s", tc.descr), func(t *testing.T) {
			t.Parallel()

			res := []syslog.Result{}
			strictParser := NewParser(
				syslog.WithListener(func(r *syslog.Result) {
					res = append(res, *r)
				}),
				syslog.WithMaxMessageLength(tc.maxMessageLength))
			strictParser.Parse(strings.NewReader(tc.input))

			assert.Equal(t, tc.results, res)
		})
		t.Run(fmt.Sprintf("effort/%s", tc.descr), func(t *testing.T) {
			t.Parallel()

			res := []syslog.Result{}
			effortParser := NewParser(
				syslog.WithBestEffort(),
				syslog.WithListener(func(r *syslog.Result) {
					res = append(res, *r)
				}),
				syslog.WithMaxMessageLength(tc.maxMessageLength),
			)
			effortParser.Parse(strings.NewReader(tc.input))

			assert.Equal(t, tc.bestEffortResults, res)
		})
	}
}

func TestParserBestEffortOption(t *testing.T) {
	p1 := NewParser().(syslog.BestEfforter)
	assert.False(t, p1.HasBestEffort())

	p2 := NewParser(syslog.WithBestEffort()).(syslog.BestEfforter)
	assert.True(t, p2.HasBestEffort())
}
