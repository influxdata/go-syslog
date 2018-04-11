package rfc5424

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(createName(tc.input), func(t *testing.T) {
			t.Parallel()

			msg, err := NewMachine().Parse(tc.input)

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
