
//line syslog.rl:1
package syslog

import (
  "fmt"
  "os"
  "github.com/influxdata/go-syslog/syslog/rfc5424"
  "github.com/davecgh/go-spew/spew"
)
 

//line syslog.go:14
const rfc5424_start int = 1
const rfc5424_first_final int = 7
const rfc5424_error int = 0

const rfc5424_en_main int = 1


//line syslog.rl:19


func utf8ToNum(bseq []uint8) int {
  out := 0
  ord := 1
  for i := len(bseq) - 1; i >= 0; i-- {
    out += (int(bseq[i]) - '0') * ord
    ord *= 10
  }
  return out
}

// RFC5424 is ...
func RFC5424(data string) {
    cs, p, pe := 0, 0, len(data)

    privalChars := []uint8{}
    var prival *rfc5424.Prival

    
//line syslog.rl:45


    
    
//line syslog.go:48
	{
	cs = rfc5424_start
	}

//line syslog.rl:49
    
//line syslog.go:55
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 7:
		goto st_case_7
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	}
	goto st_out
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 48:
			goto tr2
		case 49:
			goto tr3
		}
		if 50 <= data[p] && data[p] <= 57 {
			goto tr4
		}
		goto st0
tr2:
//line syslog.rl:39
 privalChars = append(privalChars, data[p]) 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line syslog.go:112
		if data[p] == 62 {
			goto tr5
		}
		goto st0
tr5:
//line syslog.rl:40
 prival = rfc5424.NewPrival(utf8ToNum(privalChars)) 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line syslog.go:126
		goto st0
tr3:
//line syslog.rl:39
 privalChars = append(privalChars, data[p]) 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line syslog.go:137
		switch data[p] {
		case 57:
			goto tr6
		case 62:
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 56 {
			goto tr4
		}
		goto st0
tr4:
//line syslog.rl:39
 privalChars = append(privalChars, data[p]) 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line syslog.go:157
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
tr6:
//line syslog.rl:39
 privalChars = append(privalChars, data[p]) 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line syslog.go:174
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto tr2
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line syslog.rl:50
    if cs < rfc5424_first_final {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error"))
    }

    spew.Dump(prival)
}