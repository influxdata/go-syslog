package nontransparent

import (
	"github.com/davecgh/go-spew/spew"
	"io"
	"math/rand"
	"strings"

	"github.com/influxdata/go-syslog/v3"
	"time"
)

func Example_withoutTrailerAtEnd() {
	results := []syslog.Result{}
	acc := func(res *syslog.Result) {
		results = append(results, *res)
	}
	// Notice the message ends without trailer but we catch it anyway
	r := strings.NewReader("<1>1 2003-10-11T22:14:15.003Z host.local - - - - mex")
	NewParser(syslog.WithListener(acc)).Parse(r)
	output(results)
	// Output:
	// ([]syslog.Result) (len=1) {
	//  (syslog.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Base: (syslog.Base) {
	//     Facility: (*uint8)(0),
	//     Severity: (*uint8)(1),
	//     Priority: (*uint8)(1),
	//     Timestamp: (*time.Time)(2003-10-11 22:14:15.003 +0000 UTC),
	//     Hostname: (*string)((len=10) "host.local"),
	//     Appname: (*string)(<nil>),
	//     ProcID: (*string)(<nil>),
	//     MsgID: (*string)(<nil>),
	//     Message: (*string)((len=3) "mex")
	//    },
	//    Version: (uint16) 1,
	//    StructuredData: (*map[string]map[string]string)(<nil>)
	//   }),
	//   Error: (*ragel.ReadingError)(unexpected EOF)
	//  }
	// }
}

func Example_bestEffortOnLastOne() {
	results := []syslog.Result{}
	acc := func(res *syslog.Result) {
		results = append(results, *res)
	}
	r := strings.NewReader("<1>1 - - - - - - -\n<3>1\n")
	NewParser(syslog.WithBestEffort(), syslog.WithListener(acc)).Parse(r)
	output(results)
	// Output:
	// ([]syslog.Result) (len=2) {
	//  (syslog.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Base: (syslog.Base) {
	//     Facility: (*uint8)(0),
	//     Severity: (*uint8)(1),
	//     Priority: (*uint8)(1),
	//     Timestamp: (*time.Time)(<nil>),
	//     Hostname: (*string)(<nil>),
	//     Appname: (*string)(<nil>),
	//     ProcID: (*string)(<nil>),
	//     MsgID: (*string)(<nil>),
	//     Message: (*string)((len=1) "-")
	//    },
	//    Version: (uint16) 1,
	//    StructuredData: (*map[string]map[string]string)(<nil>)
	//   }),
	//   Error: (error) <nil>
	//  },
	//  (syslog.Result) {
	//   Message: (*rfc5424.SyslogMessage)({
	//    Base: (syslog.Base) {
	//     Facility: (*uint8)(0),
	//     Severity: (*uint8)(3),
	//     Priority: (*uint8)(3),
	//     Timestamp: (*time.Time)(<nil>),
	//     Hostname: (*string)(<nil>),
	//     Appname: (*string)(<nil>),
	//     ProcID: (*string)(<nil>),
	//     MsgID: (*string)(<nil>),
	//     Message: (*string)(<nil>)
	//    },
	//    Version: (uint16) 1,
	//    StructuredData: (*map[string]map[string]string)(<nil>)
	//   }),
	//   Error: (*errors.errorString)(parsing error [col 4])
	//  }
	// }
}

func Example_intoChannelWithLF() {
	messages := []string{
		"<2>1 - - - - - - A\nB",
		"<1>1 -",
		"<1>1 - - - - - - A\nB\nC\nD",
	}

	r, w := io.Pipe()

	go func() {
		defer w.Close()

		for _, m := range messages {
			// Write message (containing trailers to be interpreted as part of the syslog MESSAGE)
			w.Write([]byte(m))
			// Write non-transparent frame boundary
			w.Write([]byte{10})
			// Wait a random amount of time
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	results := make(chan *syslog.Result)
	ln := func(x *syslog.Result) {
		// Emit the result
		results <- x
	}

	p := NewParser(syslog.WithListener(ln), syslog.WithBestEffort())
	go func() {
		defer close(results)
		defer r.Close()
		p.Parse(r)
	}()

	// Consume results
	for r := range results {
		output(r)
	}

	// Output:
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   Base: (syslog.Base) {
	//    Facility: (*uint8)(0),
	//    Severity: (*uint8)(2),
	//    Priority: (*uint8)(2),
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)(<nil>),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    Message: (*string)((len=3) "A\nB")
	//   },
	//   Version: (uint16) 1,
	//   StructuredData: (*map[string]map[string]string)(<nil>)
	//  }),
	//  Error: (error) <nil>
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   Base: (syslog.Base) {
	//    Facility: (*uint8)(0),
	//    Severity: (*uint8)(1),
	//    Priority: (*uint8)(1),
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)(<nil>),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    Message: (*string)(<nil>)
	//   },
	//   Version: (uint16) 1,
	//   StructuredData: (*map[string]map[string]string)(<nil>)
	//  }),
	//  Error: (*errors.errorString)(parsing error [col 6])
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   Base: (syslog.Base) {
	//    Facility: (*uint8)(0),
	//    Severity: (*uint8)(1),
	//    Priority: (*uint8)(1),
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)(<nil>),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    Message: (*string)((len=7) "A\nB\nC\nD")
	//   },
	//   Version: (uint16) 1,
	//   StructuredData: (*map[string]map[string]string)(<nil>)
	//  }),
	//  Error: (error) <nil>
	// })

}

func Example_intoChannelWithNUL() {
	messages := []string{
		"<2>1 - - - - - - A\x00B",
		"<1>1 -",
		"<1>1 - - - - - - A\x00B\x00C\x00D",
	}

	r, w := io.Pipe()

	go func() {
		defer w.Close()

		for _, m := range messages {
			// Write message (containing trailers to be interpreted as part of the syslog MESSAGE)
			w.Write([]byte(m))
			// Write non-transparent frame boundary
			w.Write([]byte{0})
			// Wait a random amount of time
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	results := make(chan *syslog.Result)
	ln := func(x *syslog.Result) {
		// Emit the result
		results <- x
	}

	p := NewParser(syslog.WithListener(ln), WithTrailer(NUL))

	go func() {
		defer close(results)
		defer r.Close()
		p.Parse(r)
	}()

	// Range over the results channel
	for r := range results {
		output(r)
	}

	// Output:
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   Base: (syslog.Base) {
	//    Facility: (*uint8)(0),
	//    Severity: (*uint8)(2),
	//    Priority: (*uint8)(2),
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)(<nil>),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    Message: (*string)((len=3) "A\x00B")
	//   },
	//   Version: (uint16) 1,
	//   StructuredData: (*map[string]map[string]string)(<nil>)
	//  }),
	//  Error: (error) <nil>
	// })
	// (*syslog.Result)({
	//  Message: (syslog.Message) <nil>,
	//  Error: (*errors.errorString)(parsing error [col 6])
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   Base: (syslog.Base) {
	//    Facility: (*uint8)(0),
	//    Severity: (*uint8)(1),
	//    Priority: (*uint8)(1),
	//    Timestamp: (*time.Time)(<nil>),
	//    Hostname: (*string)(<nil>),
	//    Appname: (*string)(<nil>),
	//    ProcID: (*string)(<nil>),
	//    MsgID: (*string)(<nil>),
	//    Message: (*string)((len=7) "A\x00B\x00C\x00D")
	//   },
	//   Version: (uint16) 1,
	//   StructuredData: (*map[string]map[string]string)(<nil>)
	//  }),
	//  Error: (error) <nil>
	// })
}

func output(out interface{}) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Dump(out)
}
