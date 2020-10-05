package octetcounting

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/influxdata/go-syslog/v3"
	syslogtesting "github.com/influxdata/go-syslog/v3/testing"
)

// This is here to avoid compiler optimizations that
// could remove the actual call we are benchmarking
// during benchmarks
var benchParseResult syslog.Message

type benchCase struct {
	input     []byte
	label     string
	maxLength int
}

var benchCases = []benchCase{
	{
		label: "Small Message Size",
		input: []byte("48 <1>1 2003-10-11T22:14:15.003Z host.local - - - -25 <3>1 - host.local - - - -38 <2>1 - host.local su - - - κόσμε"),
	},
	{
		label: "Default Max Message Size",
		input: []byte(fmt.Sprintf(
			"8192 <%d>%d %s %s %s %s %s - %s",
			syslogtesting.MaxPriority,
			syslogtesting.MaxVersion,
			syslogtesting.MaxRFC3339MicroTimestamp,
			string(syslogtesting.MaxHostname),
			string(syslogtesting.MaxAppname),
			string(syslogtesting.MaxProcID),
			string(syslogtesting.MaxMsgID),
			string(syslogtesting.MaxMessage),
		)),
	},
	{
		label: "UDP Max Message Size",
		input: []byte(fmt.Sprintf(
			"65529 <%d>%d %s %s %s %s %s - %s",
			syslogtesting.MaxPriority,
			syslogtesting.MaxVersion,
			syslogtesting.MaxRFC3339MicroTimestamp,
			string(syslogtesting.MaxHostname),
			string(syslogtesting.MaxAppname),
			string(syslogtesting.MaxProcID),
			string(syslogtesting.MaxMsgID),
			string(syslogtesting.LongerMaxMessage),
		)),
		maxLength: 65529,
	},
}

func BenchmarkParse(b *testing.B) {
	for _, tc := range benchCases {
		tc := tc
		if tc.maxLength == 0 {
			tc.maxLength = 8192
		}
		m := NewParser(syslog.WithBestEffort())
		b.Run(syslogtesting.RightPad(tc.label, 50), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := bytes.NewReader(tc.input)
				m.Parse(reader)
			}
		})
	}
}
