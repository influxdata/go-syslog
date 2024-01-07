package rfc3164

import (
	"testing"
	"time"

	"github.com/influxdata/go-syslog/v3"
	syslogtesting "github.com/influxdata/go-syslog/v3/testing"
	"github.com/stretchr/testify/assert"
)

// todo > add support for testing `best effort` mode

type testCase struct {
	input        []byte
	valid        bool
	shouldSkip   bool
	value        syslog.Message
	errorString  string
	partialValue syslog.Message
}

var testCases = []testCase{
	{
		[]byte(`<34>Jan 12 06:30:00 xxx apache: 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
		true,
		false,
		&SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(34),
				Severity:  syslogtesting.Uint8Address(2),
				Facility:  syslogtesting.Uint8Address(4),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Jan 12 06:30:00"),
				Hostname:  syslogtesting.StringAddress("xxx"),
				Appname:   syslogtesting.StringAddress("apache"),
				Message:   syslogtesting.StringAddress(`1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
			},
		},
		"",
		nil,
	},
	{
		[]byte(`<34>Aug  7 06:30:00 xxx aaa: message from 1.2.3.4`),
		true,
		false,
		&SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(34),
				Severity:  syslogtesting.Uint8Address(2),
				Facility:  syslogtesting.Uint8Address(4),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Aug 7 06:30:00"),
				Hostname:  syslogtesting.StringAddress("xxx"),
				Appname:   syslogtesting.StringAddress("aaa"),
				Message:   syslogtesting.StringAddress(`message from 1.2.3.4`),
			},
		},
		"",
		nil,
	},
	{
		input:      []byte(`<85>Jan 24 15:50:41 ip-172-31-30-110 sudo[6040]: ec2-user : TTY=pts/0 ; PWD=/var/log ; USER=root ; COMMAND=/bin/tail secure`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(85),
				Facility:  syslogtesting.Uint8Address(10),
				Severity:  syslogtesting.Uint8Address(5),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Jan 24 15:50:41"),
				Hostname:  syslogtesting.StringAddress("ip-172-31-30-110"),
				Appname:   syslogtesting.StringAddress("sudo"),
				ProcID:    syslogtesting.StringAddress("6040"),
				Message:   syslogtesting.StringAddress(`ec2-user : TTY=pts/0 ; PWD=/var/log ; USER=root ; COMMAND=/bin/tail secure`),
			},
		},
	},
	{
		input:      []byte(`<166>Jul  6 20:33:28 ABC-1-234567 Some message here`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(166),
				Facility:  syslogtesting.Uint8Address(20),
				Severity:  syslogtesting.Uint8Address(6),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Jul  6 20:33:28"),
				Hostname:  syslogtesting.StringAddress("ABC-1-234567"),
				Message:   syslogtesting.StringAddress("Some message here"),
			},
		},
	},
	{
		input:      []byte(`<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(34),
				Facility:  syslogtesting.Uint8Address(4),
				Severity:  syslogtesting.Uint8Address(2),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Oct 11 22:14:15"),
				Hostname:  syslogtesting.StringAddress("mymachine"),
				Appname:   syslogtesting.StringAddress("su"),
				Message:   syslogtesting.StringAddress("'su root' failed for lonvick on /dev/pts/8"),
			},
		},
	},
	{
		input:      []byte(`<13>Feb  5 17:32:18 10.0.0.99 Use the BFG!`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(13),
				Facility:  syslogtesting.Uint8Address(1),
				Severity:  syslogtesting.Uint8Address(5),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Feb  5 17:32:18"),
				Hostname:  syslogtesting.StringAddress("10.0.0.99"),
				Message:   syslogtesting.StringAddress("Use the BFG!"),
			},
		},
	},
	{
		input:      []byte(`<165>Aug 24 05:34:00 mymachine myproc[10]: %% It's time to make the do-nuts.  %%  Ingredients: Mix=OK, Jelly=OK # Devices: Mixer=OK, Jelly_Injector=OK, Frier=OK # Transport: Conveyer1=OK, Conveyer2=OK # %%`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(165),
				Facility:  syslogtesting.Uint8Address(20),
				Severity:  syslogtesting.Uint8Address(5),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Aug 24 05:34:00"),
				Hostname:  syslogtesting.StringAddress("mymachine"),
				Appname:   syslogtesting.StringAddress("myproc"),
				ProcID:    syslogtesting.StringAddress("10"),
				Message:   syslogtesting.StringAddress(`%% It's time to make the do-nuts.  %%  Ingredients: Mix=OK, Jelly=OK # Devices: Mixer=OK, Jelly_Injector=OK, Frier=OK # Transport: Conveyer1=OK, Conveyer2=OK # %%`),
			},
		},
	},
	{
		input:      []byte(`<0>Oct 22 10:52:01 10.1.2.3 sched[0]: That's All Folks!`),
		valid:      true,
		shouldSkip: false,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(0),
				Facility:  syslogtesting.Uint8Address(0),
				Severity:  syslogtesting.Uint8Address(0),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Oct 22 10:52:01"),
				Hostname:  syslogtesting.StringAddress("10.1.2.3"),
				Appname:   syslogtesting.StringAddress("sched"),
				ProcID:    syslogtesting.StringAddress("0"),
				Message:   syslogtesting.StringAddress(`That's All Folks!`),
			},
		},
	},
	{
		input:      []byte(`Oct 22 10:52:01 10.1.2.3 sched[0]: That's All Folks!`),
		valid:      true,
		shouldSkip: true,
		value: &SyslogMessage{
			Base: syslog.Base{
				Priority:  syslogtesting.Uint8Address(0),
				Facility:  syslogtesting.Uint8Address(0),
				Severity:  syslogtesting.Uint8Address(0),
				Timestamp: syslogtesting.TimeParse(time.Stamp, "Oct 22 10:52:01"),
				Hostname:  syslogtesting.StringAddress("10.1.2.3"),
				Appname:   syslogtesting.StringAddress("sched"),
				ProcID:    syslogtesting.StringAddress("0"),
				Message:   syslogtesting.StringAddress(`That's All Folks!`),
			},
		},
	},
	{
		input:       []byte(`Oct 22 10:52:01 10.1.2.3 sched[0]: That's All Folks!`),
		valid:       false,
		shouldSkip:  false,
		value:       nil,
		errorString: "expecting a priority value within angle brackets [col 0]",
	},
	// todo > other test cases pleaaaase
}

func TestMachineParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(syslogtesting.RightPad(string(tc.input), 50), func(t *testing.T) {
			t.Parallel()
			var merr, perr error
			var message, partial syslog.Message

			if tc.shouldSkip {
				message, merr = NewMachine(WithAllowSkipPri()).Parse(tc.input)
				partial, perr = NewMachine(WithBestEffort(), WithAllowSkipPri()).Parse(tc.input)
			} else {
				message, merr = NewMachine().Parse(tc.input)
				partial, perr = NewMachine(WithBestEffort()).Parse(tc.input)
			}

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

			if tc.errorString != "" {
				assert.Equal(t, tc.value, message)
			}
		})
	}
}
