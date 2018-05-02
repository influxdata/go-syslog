[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](LICENSE)

**A parser for syslog messages**.

> Blazing fast RFC5424-compliant parser

## Installation

```
go get github.com/influxdata/go-syslog
```

## Docs

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/influxdata/go-syslog)

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

## Performance

To run the benchmark execute the following command.

```bash
make bench
```

On my machine<sup>[1](#mymachine)</sup> this are the results obtained in best effort mode.

```
[no]_empty_input__________________________________-4	30000000       253 ns/op     224 B/op       3 allocs/op
[no]_multiple_syslog_messages_on_multiple_lines___-4	20000000       433 ns/op     304 B/op      12 allocs/op
[no]_impossible_timestamp_________________________-4	10000000      1080 ns/op     528 B/op      11 allocs/op
[no]_malformed_structured_data____________________-4	20000000       552 ns/op     400 B/op      12 allocs/op
[no]_with_duplicated_structured_data_id___________-4	 5000000      1246 ns/op     688 B/op      17 allocs/op
[ok]_minimal______________________________________-4	30000000       264 ns/op     247 B/op       9 allocs/op
[ok]_average_message______________________________-4	 5000000      1984 ns/op    1536 B/op      26 allocs/op
[ok]_complicated_message__________________________-4	 5000000      1644 ns/op    1280 B/op      25 allocs/op
[ok]_very_long_message____________________________-4	 2000000      3826 ns/op    2464 B/op      28 allocs/op
[ok]_all_max_length_and_complete__________________-4	 3000000      2792 ns/op    1888 B/op      28 allocs/op
[ok]_all_max_length_except_structured_data_and_mes-4	 5000000      1830 ns/op     883 B/op      13 allocs/op
[ok]_minimal_with_message_containing_newline______-4	20000000       294 ns/op     250 B/op      10 allocs/op
[ok]_w/o_procid,_w/o_structured_data,_with_message-4	10000000       956 ns/op     364 B/op      11 allocs/op
[ok]_minimal_with_UTF-8_message___________________-4	20000000       586 ns/op     359 B/op      10 allocs/op
[ok]_with_structured_data_id,_w/o_structured_data_-4	10000000       998 ns/op     592 B/op      14 allocs/op
[ok]_with_multiple_structured_data________________-4	 5000000      1538 ns/op    1232 B/op      22 allocs/op
[ok]_with_escaped_backslash_within_structured_data-4	 5000000      1316 ns/op     920 B/op      20 allocs/op
[ok]_with_UTF-8_structured_data_param_value,_with_-4	 5000000      1580 ns/op    1050 B/op      21 allocs/op
```

As you can see it takes:

* ~250ns to parse the smallest legal message

* ~2µs to parse an average legal message

* ~4µs to parse a very long legal message

Other RFC5424 implementations, like this [one](https://github.com/roguelazer/rust-syslog-rfc5424) in Rust, spend 8µs to parse an average legal message.

_TBD: comparation against other golang parsers_.

---

* <a name="mymachine">[1]</a>: Intel Core i7-7600U CPU @ 2.80GHz


