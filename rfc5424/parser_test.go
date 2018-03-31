package rfc5424

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input       string
	valid       bool
	value       *SyslogMessage
	errorString string
}

func timeParse(layout, value string) *time.Time {
	t, _ := time.Parse(layout, value)
	return &t
}

var testCases = []testCase{
	// Wrong date
	{
		"<101>122 201-11-22",
		false,
		nil,
		"error parsing <nilvalue>",
	},
	// Incomplete date
	{
		"<191>123 2018-02-29",
		false,
		nil,
		"error parsing <nilvalue>",
	},
	// Right date
	{
		"<187>222 1985-04-12T23:20:50.003Z",
		true,
		&SyslogMessage{
			Header: Header{
				Pri: Pri{
					Prival: Prival{
						Facility: Facility{
							Code: 23,
						},
						Severity: Severity{
							Code: 3,
						},
						Value: 187,
					},
				},
				Version: Version{
					Value: 222,
				},
				Timestamp: timeParse(time.RFC3339Nano, "1985-04-12T23:20:50.003Z"),
			},
		},
		"",
	},
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			msg, err := Parse(tc.input)

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
	for _, tc := range testCases {
		tc := tc
		b.Run(tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchParseResult, _ = Parse(tc.input)
			}
		})
	}
}
