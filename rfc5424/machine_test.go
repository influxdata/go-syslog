package rfc5424

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeParse(layout, value string) *time.Time {
	t, _ := time.Parse(layout, value)
	return &t
}

func getStringAddress(str string) *string {
	return &str
}

func createName(x []byte) string {
	str := string(x)
	lim := 30
	if len(str) > lim {
		return string(str[0:30])
	}
	return str
}

// []byte("<22>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [id1 a=\"b\" c=\"d\" e=\"\"][id2 z=\"w\"]")

type testCase struct {
	input       []byte
	valid       bool
	value       *SyslogMessage
	errorString string
}

var testCases = []testCase{
	// Malformed pri
	{
		[]byte("(190>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
	},
	// (fixme) > Malformed pri outputs wrong error (inner error regarding prival, not the outer one regarding pri)
	// {
	// 	[]byte("<190)122 2018-11-22"),
	// 	false,
	// 	nil,
	// 	"expecting a priority value within angle brackets [col 4]",
	// },
	// Missing pri
	{
		[]byte("122 2018-11-22"),
		false,
		nil,
		"expecting a priority value within angle brackets [col 0]",
	},
	// Missing prival
	{
		[]byte("<>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 1]",
	},
	// Prival with too much digits
	{
		[]byte("<19000021>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 4]",
	},
	// Prival too high
	{
		[]byte("<192>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 3]",
	},
	// 0 starting prival
	{
		[]byte("<002>122 2018-11-22"),
		false,
		nil,
		"expecting a priority value in the range 1-191 or equal to 0 [col 2]",
	},
	// Missing version
	{
		[]byte("<100> 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
	},
	// 0 version
	{
		[]byte("<100>0 2018-11-22"),
		false,
		nil,
		"expecting a version value in the range 1-999 [col 5]",
	},
	// Wrong year
	{
		[]byte("<101>122 201-11-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 12]",
	},
	// Wrong month
	{
		[]byte("<101>122 2018-112-22"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 16]",
	},
	// Wrong day
	{
		[]byte("<191>123 2018-02-32"),
		false,
		nil,
		"expecting a RFC3339 or a RFC3339NANO timestamp or a nil value [col 18]",
	},
	// Non existing date
	{
		[]byte("<34>11 2003-02-30T05:14:15.000003-07:00 "), // (fixme) > needed space here ...
		false,
		nil,
		"parsing time \"2003-02-30T05:14:15.000003-07:00\": day out of range",
	},
	// All right but without structured data
	{
		[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			Header: Header{
				Pri: Pri{
					Prival: Prival{
						Facility: 4,
						Severity: 2,
						Value:    34,
					},
				},
				Version:   1,
				Timestamp: timeParse(time.RFC3339Nano, "2003-10-11T22:14:15.003Z"),
				Hostname:  getStringAddress("mymachine.example.com"),
				Appname:   getStringAddress("su"),
				ProcID:    nil,
				MsgID:     getStringAddress("ID47"),
			},
			StructuredData: nil,
			Message:        getStringAddress("BOM'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
	},
	// Nil timestamp
	{
		[]byte("<187>222 - mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"),
		true,
		&SyslogMessage{
			Header: Header{
				Pri: Pri{
					Prival: Prival{
						Facility: 23,
						Severity: 3,
						Value:    187,
					},
				},
				Version:   222,
				Timestamp: nil,
				Hostname:  getStringAddress("mymachine.example.com"),
				Appname:   getStringAddress("su"),
				ProcID:    nil,
				MsgID:     getStringAddress("ID47"),
			},
			StructuredData: nil,
			Message:        getStringAddress("BOM'su root' failed for lonvick on /dev/pts/8"),
		},
		"",
	},
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(createName(tc.input), func(t *testing.T) {
			t.Parallel()

			// (fixme) > instantiating a new machine each time since otherwise there are racing conditions
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

// This is here to avoid compiler optimizations that
// could remove the actual call we are benchmarking
// during benchmarks
var benchParseResult *SyslogMessage

func BenchmarkParse(b *testing.B) {
	parser := NewMachine()
	for _, tc := range testCases {
		tc := tc
		b.Run(createName(tc.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchParseResult, _ = parser.Parse(tc.input)
			}
		})
	}
}
