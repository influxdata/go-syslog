**A parser for RFC5424 syslog messages**.

> Blazing fast syslog parser

## Installation

```
go get github.com/influxdata/go-syslog
```

## Docs

API documentation can be found [here](https://godoc.org/github.com/influxdata/go-syslog/rfc5424).

The [docs](docs/) directory contains images representing the FSM parts of a RFC5424 syslog message.

## Usage


```go
i := []byte(`<165>4 2018-10-11T22:14:15.003Z mymach.it e - 1 [ex@32473 iut="3"] An application event log entry...`)
p := NewParser()
m, e := p.Parse(i, nil)
fmt.Printf("%#v\n", m)
// &rfc5424.SyslogMessage{
//  Priority:  (*uint8)(0xc420098a73)(165),
//  facility:  (*uint8)(0xc420098a74)(20),
//  severity:  (*uint8)(0xc420098a75)(5),
//  Version:   (uint16) 4,
//  Timestamp: (*time.Time)(0xc42011a8e0)(2018-10-11 22:14:15.003 +0000 UTC),
//  Hostname:  (*string)(0xc42008d5a0)((len=9) "mymach.it"),
//  Appname:   (*string)(0xc42008d5b0)((len=1) "e"),
//  MsgID:     (*string)(0xc42008d5d0)((len=1) "1"),
//  StructuredData: (*map[string]map[string]string)(0xc42009a080)((len=1) {
//   (string) (len=8) "ex@32473": (map[string]string) (len=1) {
//    (string) (len=3) "iut": (string) (len=1) "3"
//   }
//  }),
//  Message: (*string)(0xc42008d5e0)((len=33) "An application event log entry...")
// }
fmt.Println(e)
// <nil>
```

### Best effort mode

This modality enables partial parsing.

When the parsing process errors out it returns the message collected until that position, and the error that caused the parser to stop.

Notice that in this modality the output is returned iff it represents a minimally valid message - ie., a message containing almost a priority field in `[1,191]` within angular brackets, followed by a version in `]0,999]`.

```go
bestEffort := true
i := []byte("<1>1 - - - - - X")
p := NewParser()
m, e := p.Parse(i, &bestEffort)
fmt.Printf("%#v\n", m)
// &rfc5424.SyslogMessage{
//	Priority: (*uint8)(0xc4200988bd)(1),
//  facility: (*uint8)(0xc4200988be)(0),
//  severity: (*uint8)(0xc4200988bf)(1),
//  Version:  (uint16) 1
// }
fmt.Println(e)
// expecting a structured data section containing one or more elements (`[id( key="value")*]+`) or a nil value [col 15]
```

---
