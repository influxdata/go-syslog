package rfc5424

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	valid bool
	value *SyslogMessage
}

var testCases = []testCase{
	{
		"<101>122 201-11-22",
		false,
		nil,
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
	},
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			msg, err := Parse(tc.input)
			if !tc.valid {
				assert.Nil(t, msg)
				assert.Error(t, err)
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
