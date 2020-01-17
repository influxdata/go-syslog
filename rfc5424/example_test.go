package rfc5424

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func output(out interface{}) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Dump(out)
}

func Example() {
	i := []byte(`<165>4 2018-10-11T22:14:15.003Z mymach.it e - 1 [ex@32473 iut="3"] An application event log entry...`)
	p := NewParser()
	m, _ := p.Parse(i)
	output(m)
	// Output:
	// (*rfc5424.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(20),
	//   Severity: (*uint8)(5),
	//   Priority: (*uint8)(165),
	//   Timestamp: (*time.Time)(2018-10-11 22:14:15.003 +0000 UTC),
	//   Hostname: (*string)((len=9) "mymach.it"),
	//   Appname: (*string)((len=1) "e"),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)((len=1) "1"),
	//   Message: (*string)((len=33) "An application event log entry...")
	//  },
	//  Version: (uint16) 4,
	//  StructuredData: (*map[string]map[string]string)((len=1) {
	//   (string) (len=8) "ex@32473": (map[string]string) (len=1) {
	//    (string) (len=3) "iut": (string) (len=1) "3"
	//   }
	//  })
	// })
}

func Example_besteffort() {
	i := []byte(`<1>1 A - - - - - -`)
	p := NewParser(WithBestEffort())
	m, e := p.Parse(i)
	output(m)
	fmt.Println(e)
	// Output:
	// (*rfc5424.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(0),
	//   Severity: (*uint8)(1),
	//   Priority: (*uint8)(1),
	//   Timestamp: (*time.Time)(<nil>),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   Message: (*string)(<nil>)
	//  },
	//  Version: (uint16) 1,
	//  StructuredData: (*map[string]map[string]string)(<nil>)
	// })
	// expecting a RFC3339MICRO timestamp or a nil value [col 5]
}

func Example_builder() {
	msg := &SyslogMessage{}
	msg.SetTimestamp("not a RFC3339MICRO timestamp")
	fmt.Println("Valid?", msg.Valid())
	msg.SetPriority(191)
	msg.SetVersion(1)
	fmt.Println("Valid?", msg.Valid())
	output(msg)
	str, _ := msg.String()
	fmt.Println(str)
	// Output:
	// Valid? false
	// Valid? true
	// (*rfc5424.SyslogMessage)({
	//  Base: (syslog.Base) {
	//   Facility: (*uint8)(23),
	//   Severity: (*uint8)(7),
	//   Priority: (*uint8)(191),
	//   Timestamp: (*time.Time)(<nil>),
	//   Hostname: (*string)(<nil>),
	//   Appname: (*string)(<nil>),
	//   ProcID: (*string)(<nil>),
	//   MsgID: (*string)(<nil>),
	//   Message: (*string)(<nil>)
	//  },
	//  Version: (uint16) 1,
	//  StructuredData: (*map[string]map[string]string)(<nil>)
	// })
	// <191>1 - - - - - -
}
