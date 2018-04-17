**A parser for syslog messages**.

> Blazing fast RFC5424-compliant parser

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

## Performance

To run the benchmark execute the following command.

```bash
make bench
```

On my machine<sup>[1](#mymachine)</sup> this are the results obtained in best effort mode.

```
goos: linux
goarch: amd64
pkg: github.com/influxdata/go-syslog/rfc5424
BenchmarkParse/[no]_empty_input__________________________________-4             30000000               250 ns/op             176 B/op          3 allocs/op
BenchmarkParse/[no]_multiple_syslog_messages_on_multiple_lines___-4             20000000               478 ns/op             224 B/op         15 allocs/op
BenchmarkParse/[no]_impossible_timestamp_________________________-4             10000000              1098 ns/op             424 B/op         17 allocs/op
BenchmarkParse/[no]_malformed_structured_data____________________-4             20000000               560 ns/op             320 B/op         15 allocs/op
BenchmarkParse/[no]_with_duplicated_structured_data_id___________-4              5000000              1435 ns/op             656 B/op         29 allocs/op
BenchmarkParse/[ok]_minimal______________________________________-4             30000000               288 ns/op             167 B/op         12 allocs/op
BenchmarkParse/[ok]_average_message______________________________-4              3000000              2212 ns/op            1544 B/op         37 allocs/op
BenchmarkParse/[ok]_complicated_message__________________________-4              5000000              1852 ns/op            1272 B/op         37 allocs/op
BenchmarkParse/[ok]_very_long_message____________________________-4              2000000              3978 ns/op            2472 B/op         43 allocs/op
BenchmarkParse/[ok]_all_max_length_and_complete__________________-4              3000000              3014 ns/op            1880 B/op         43 allocs/op
BenchmarkParse/[ok]_all_max_length_except_structured_data_and_mes-4              5000000              1934 ns/op             841 B/op         23 allocs/op
BenchmarkParse/[ok]_minimal_with_message_containing_newline______-4             20000000               363 ns/op             186 B/op         14 allocs/op
BenchmarkParse/[ok]_w/o_procid,_w/o_structured_data,_with_message-4             10000000              1075 ns/op             336 B/op         19 allocs/op
BenchmarkParse/[ok]_minimal_with_UTF-8_message___________________-4             20000000               539 ns/op             295 B/op         14 allocs/op
BenchmarkParse/[ok]_with_structured_data_id,_w/o_structured_data_-4             10000000              1024 ns/op             552 B/op         22 allocs/op
BenchmarkParse/[ok]_with_multiple_structured_data________________-4              5000000              1757 ns/op            1197 B/op         32 allocs/op
BenchmarkParse/[ok]_with_escaped_backslash_within_structured_data-4              5000000              1385 ns/op             888 B/op         28 allocs/op
BenchmarkParse/[ok]_with_UTF-8_structured_data_param_value,_with_-4              5000000              1560 ns/op            1032 B/op         31 allocs/op
```

As you can see it takes:

* ~300ns to parse the smallest legal message

* ~2.2µs to parse an average legal message

* ~4.0µs to parse a very long legal message

Other RFC5424 implementations, like this [one](https://github.com/roguelazer/rust-syslog-rfc5424) in Rust, spend 8µs to parse an average legal message.

_TBD: comparation against other golang parsers_.

---

* <a name="mymachine">[1]</a>: Intel Core i7-7600U CPU @ 2.80GHz
