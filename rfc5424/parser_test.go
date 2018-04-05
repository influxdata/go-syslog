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
	// {
	// 	"<101>122 201-11-22",
	// 	false,
	// 	nil,
	// 	"error parsing <nilvalue>",
	// },
	// Prival too high
	// {
	// 	"<333>122 2018-11-22",
	// 	false,
	// 	nil,
	// 	"generic error",
	// },
	// Missing version
	// {
	// 	"<100> 2018-11-22",
	// 	false,
	// 	nil,
	// 	"generic error",
	// },
	// Incomplete date
	// {
	// 	"<191>123 2018-02-29",
	// 	false,
	// 	nil,
	// 	"error parsing <nilvalue>",
	// },
	// All right but without structured data
	{
		"<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8", // need bom of utf8 encoded string
		true,
		&SyslogMessage{
			Header: Header{
				Pri: Pri{
					Prival: Prival{
						Facility: Facility{
							Code: 4,
						},
						Severity: Severity{
							Code: 2,
						},
						Value: 34,
					},
				},
				Version: Version{
					Value: 1,
				},
				Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
				Hostname:  "mymachine.example.com",
				Appname:   "su",
				ProcID:    "",
				MsgID:     "ID47",
			},
		},
		"",
	},
	// All right with nilvalue (-) as date
	// {
	// 	"<187>222 -",
	// 	true,
	// 	&SyslogMessage{
	// 		Header: Header{
	// 			Pri: Pri{
	// 				Prival: Prival{
	// 					Facility: Facility{
	// 						Code: 23,
	// 					},
	// 					Severity: Severity{
	// 						Code: 3,
	// 					},
	// 					Value: 187,
	// 				},
	// 			},
	// 			Version: Version{
	// 				Value: 222,
	// 			},
	// 			Timestamp: nil,
	// 		},
	// 	},
	// 	"",
	// },
	// All right but prival too high
	// {
	// 	"<999>222 1985-04-12T23:20:50.003Z",
	// 	false,
	// 	nil,
	// 	"generic error",
	// },
	// {
	// 	"<999>222 1985-04-12T23:20:50.003Z hostname",
	// 	false,
	// 	nil,
	// 	"generic error",
	// },
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
