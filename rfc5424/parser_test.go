package rfc5424

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input       string
	valid       bool
	value       *SyslogMessage
	errorString string
}

var testCases = []testCase{
	{
		"<101>122 201-11-22",
		false,
		nil,
		"errore generico",
	},
	{
		"<101>122 2018-11-22",
		true,
		&SyslogMessage{
			Header: Header{
				Pri: Pri{
					Prival: Prival{
						Facility: Facility{
							Code: 12,
						},
						Severity: Severity{
							Code: 5,
						},
						Value: 101,
					},
				},
				Version: Version{
					Value: 122,
				},
			},
		},
		"",
	},
	{
		"<191>123 2018-02-29",
		false,
		nil,
		"error parsing <nilvalue>",
	},
	// {
	// 	"<187>222 1985-04-12T23:20:50.003Z",
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
	// 			// (fixme)
	// 			Timestamp: time.Parse(time.RFC3339Nano, "1985-04-12 23:20:50.003 +0000 UTC"),
	// 		},
	// 	},
	// 	"",
	// },
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			msg, err := Parse(tc.input)

			fmt.Printf("%+v\n", msg)

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

func BenchmarkParse(b *testing.B) {
	for _, tc := range testCases {
		b.Run(tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Parse(tc.input)
			}
		})
	}
}
