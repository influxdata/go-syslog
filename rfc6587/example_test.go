package rfc6587

import (
	"github.com/davecgh/go-spew/spew"
	"io"
	"math/rand"

	"github.com/influxdata/go-syslog"
	"time"
)

func Example_channel_lf() {
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
	//   priority: (*uint8)(2),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(2),
	//   version: (uint16) 1,
	//   timestamp: (*time.Time)(<nil>),
	//   hostname: (*string)(<nil>),
	//   appname: (*string)(<nil>),
	//   procID: (*string)(<nil>),
	//   msgID: (*string)(<nil>),
	//   structuredData: (*map[string]map[string]string)(<nil>),
	//   message: (*string)((len=3) "A\nB")
	//  }),
	//  Error: (error) <nil>
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   priority: (*uint8)(1),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(1),
	//   version: (uint16) 1,
	//   timestamp: (*time.Time)(<nil>),
	//   hostname: (*string)(<nil>),
	//   appname: (*string)(<nil>),
	//   procID: (*string)(<nil>),
	//   msgID: (*string)(<nil>),
	//   structuredData: (*map[string]map[string]string)(<nil>),
	//   message: (*string)(<nil>)
	//  }),
	//  Error: (*errors.errorString)(parsing error [col 6])
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   priority: (*uint8)(1),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(1),
	//   version: (uint16) 1,
	//   timestamp: (*time.Time)(<nil>),
	//   hostname: (*string)(<nil>),
	//   appname: (*string)(<nil>),
	//   procID: (*string)(<nil>),
	//   msgID: (*string)(<nil>),
	//   structuredData: (*map[string]map[string]string)(<nil>),
	//   message: (*string)((len=7) "A\nB\nC\nD")
	//  }),
	//  Error: (error) <nil>
	// })
}

func Example_channel_nul() {
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
	//(*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   priority: (*uint8)(2),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(2),
	//   version: (uint16) 1,
	//   timestamp: (*time.Time)(<nil>),
	//   hostname: (*string)(<nil>),
	//   appname: (*string)(<nil>),
	//   procID: (*string)(<nil>),
	//   msgID: (*string)(<nil>),
	//   structuredData: (*map[string]map[string]string)(<nil>),
	//   message: (*string)((len=3) "A\x00B")
	//  }),
	//  Error: (error) <nil>
	// })
	// (*syslog.Result)({
	//  Message: (syslog.Message) <nil>,
	//  Error: (*errors.errorString)(parsing error [col 6])
	// })
	// (*syslog.Result)({
	//  Message: (*rfc5424.SyslogMessage)({
	//   priority: (*uint8)(1),
	//   facility: (*uint8)(0),
	//   severity: (*uint8)(1),
	//   version: (uint16) 1,
	//   timestamp: (*time.Time)(<nil>),
	//   hostname: (*string)(<nil>),
	//   appname: (*string)(<nil>),
	//   procID: (*string)(<nil>),
	//   msgID: (*string)(<nil>),
	//   structuredData: (*map[string]map[string]string)(<nil>),
	//   message: (*string)((len=7) "A\x00B\x00C\x00D")
	//  }),
	//  Error: (error) <nil>
	// })
}

func output(out interface{}) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Dump(out)
}
