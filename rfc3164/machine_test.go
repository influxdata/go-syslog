package rfc3164

import (
	"testing"
	"time"

	"github.com/influxdata/go-syslog/v2"
	syslogtesting "github.com/influxdata/go-syslog/v2/testing"
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
	{
		[]byte(`<34>Jan 12 06:30:00 xxx apache: 1.2.3.4 - - [12/Jan/2011:06:29:59 +0100] "GET /foo/bar.html HTTP/1.1" 301 96 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.9.2.12) Gecko/20101026 Firefox/3.6.12 ( .NET CLR 3.5.30729)" PID 18904 Time Taken 0`),
		true,
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
	// todo > other test cases pleaaaase
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
