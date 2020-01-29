package rfc3164

import (
	"github.com/davecgh/go-spew/spew"
)

func output(out interface{}) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Dump(out)
}

func Example() {
	i := []byte(`<13>Dec  2 16:31:03 host app: Test`)
	p := NewParser()
	m, _ := p.Parse(i)
	output(m)
	// Output:
	// (*rfc3164.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(1),
	//   Severity: (*uint8)(5),
	//   Priority: (*uint8)(13),
	//   Timestamp: (*time.Time)(0000-12-02 16:31:03 +0000 UTC),
	//   Hostname: (*string)((len=4) "host"),
	//   Appname: (*string)((len=3) "app"),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   Message: (*string)((len=4) "Test")
	//  }
	// })
}

func Example_currentyear() {
	i := []byte(`<13>Dec  2 16:31:03 host app: Test`)
	p := NewParser(WithYear(CurrentYear{}))
	m, _ := p.Parse(i)
	output(m)
	// Output:
	// (*rfc3164.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(1),
	//   Severity: (*uint8)(5),
	//   Priority: (*uint8)(13),
	//   Timestamp: (*time.Time)(2020-12-02 16:31:03 +0000 UTC),
	//   Hostname: (*string)((len=4) "host"),
	//   Appname: (*string)((len=3) "app"),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   Message: (*string)((len=4) "Test")
	//  }
	// })
}

func Example_besteffort() {
	i := []byte(`<13>Dec  2 16:31:03 -`)
	p := NewParser(WithBestEffort())
	m, _ := p.Parse(i)
	output(m)
	// Output:
	// (*rfc3164.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(1),
	//   Severity: (*uint8)(5),
	//   Priority: (*uint8)(13),
	//   Timestamp: (*time.Time)(0000-12-02 16:31:03 +0000 UTC),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   Message: (*string)(<nil>)
	//  }
	// })
}
