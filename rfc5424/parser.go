
//line rfc5424/parser.rl:1
package rfc5424

import (
  "fmt"
)
 

//line rfc5424/parser.go:11
const rfc5424_start int = 1
const rfc5424_first_final int = 11
const rfc5424_error int = 0

const rfc5424_en_main int = 1


//line rfc5424/parser.rl:10


func utf8ToNum(bseq []uint8) int {
  out := 0
  ord := 1
  for i := len(bseq) - 1; i >= 0; i-- {
    out += (int(bseq[i]) - '0') * ord
    ord *= 10
  }
  return out
}

func parse(data string) (*Message, error) {
    cs, p, pe := 0, 0, len(data)

    privalChars := []uint8{}
    versionChars := []uint8{}
    var prival *Prival
    var version *Version

    
//line rfc5424/parser.go:41
	{
	cs = rfc5424_start
	}

//line rfc5424/parser.go:46
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
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 11:
		goto st_case_11
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
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
//line rfc5424/machine.rl:4

    privalChars = append(privalChars, data[p])

	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line rfc5424/parser.go:113
		if data[p] == 62 {
			goto tr5
		}
		goto st0
tr5:
//line rfc5424/machine.rl:8

    prival = NewPrival(utf8ToNum(privalChars))

	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line rfc5424/parser.go:129
		if 49 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
tr6:
//line rfc5424/machine.rl:12

    versionChars = append(versionChars, data[p])

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line rfc5424/parser.go:145
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto st0
tr7:
//line rfc5424/machine.rl:16

    version = NewVersion(utf8ToNum(versionChars))

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line rfc5424/parser.go:164
		goto st0
tr8:
//line rfc5424/machine.rl:12

    versionChars = append(versionChars, data[p])

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line rfc5424/parser.go:177
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
tr9:
//line rfc5424/machine.rl:12

    versionChars = append(versionChars, data[p])

	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line rfc5424/parser.go:196
		if data[p] == 32 {
			goto tr7
		}
		goto st0
tr3:
//line rfc5424/machine.rl:4

    privalChars = append(privalChars, data[p])

	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line rfc5424/parser.go:212
		switch data[p] {
		case 57:
			goto tr10
		case 62:
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 56 {
			goto tr4
		}
		goto st0
tr4:
//line rfc5424/machine.rl:4

    privalChars = append(privalChars, data[p])

	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line rfc5424/parser.go:234
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
tr10:
//line rfc5424/machine.rl:4

    privalChars = append(privalChars, data[p])

	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line rfc5424/parser.go:253
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
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line rfc5424/parser.rl:35


    if cs < rfc5424_first_final {
        return nil, fmt.Errorf("error")
    } 

    return &Message{
      Prival: *prival,
      Version: *version,
    }, nil
}
