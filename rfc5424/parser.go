
//line rfc5424/parser.rl:1
package rfc5424

import (
  "fmt"
  "time"
)
 

//line rfc5424/parser.go:12
const rfc5424_start int = 1
const rfc5424_first_final int = 48
const rfc5424_error int = 0

const rfc5424_en_line int = 50
const rfc5424_en_main int = 1


//line rfc5424/parser.rl:11


func Parse(data string) (*SyslogMessage, error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    _ = eof

    cr := GetCharsRepo()

    poss := make(map[string]int, 0)

    err := fmt.Errorf("generic error")

    var prival *Prival
    var version *Version
    var timestamp *time.Time

    
//line rfc5424/parser.go:40
	{
	cs = rfc5424_start
	}

//line rfc5424/parser.go:45
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
	case 6:
		goto st_case_6
	case 48:
		goto st_case_48
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 49:
		goto st_case_49
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 50:
		goto st_case_50
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st0
tr9:
//line rfc5424/machine.rl:30

    err = fmt.Errorf("error parsing <nilvalue>");

	goto st0
//line rfc5424/parser.go:170
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

    cr.Add(data[p])

	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line rfc5424/parser.go:201
		if data[p] == 62 {
			goto tr5
		}
		goto st0
tr5:
//line rfc5424/machine.rl:8

    prival = NewPrival(*cr.Reduce())

	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line rfc5424/parser.go:217
		if 49 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
tr6:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line rfc5424/parser.go:233
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto st0
tr7:
//line rfc5424/machine.rl:12

    version = NewVersion(*cr.Reduce())

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line rfc5424/parser.go:252
		if data[p] == 45 {
			goto st48
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr11
		}
		goto tr9
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		goto st0
tr11:
//line rfc5424/machine.rl:30

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/machine.rl:16

    poss["timestamp:ini"] = p

	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line rfc5424/parser.go:281
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 45 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 48:
			goto st12
		case 49:
			goto st42
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 49 <= data[p] && data[p] <= 57 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 45 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 48:
			goto st15
		case 51:
			goto st41
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto st40
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if 49 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 84 {
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 50 {
			goto st39
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= data[p] && data[p] <= 57 {
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 58 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if 48 <= data[p] && data[p] <= 53 {
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 58 {
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= data[p] && data[p] <= 53 {
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 46:
			goto st32
		case 90:
			goto st49
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 50 {
			goto st31
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if 48 <= data[p] && data[p] <= 57 {
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 58 {
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if 48 <= data[p] && data[p] <= 53 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if 48 <= data[p] && data[p] <= 57 {
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if 48 <= data[p] && data[p] <= 51 {
			goto st28
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if 48 <= data[p] && data[p] <= 57 {
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 43:
			goto st26
		case 45:
			goto st26
		case 90:
			goto st49
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if 48 <= data[p] && data[p] <= 51 {
			goto st19
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if 48 <= data[p] && data[p] <= 49 {
			goto st16
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if 48 <= data[p] && data[p] <= 50 {
			goto st13
		}
		goto st0
tr8:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line rfc5424/parser.go:685
		if data[p] == 32 {
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr48
		}
		goto st0
tr48:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line rfc5424/parser.go:704
		if data[p] == 32 {
			goto tr7
		}
		goto st0
tr3:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line rfc5424/parser.go:720
		switch data[p] {
		case 57:
			goto tr49
		case 62:
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 56 {
			goto tr4
		}
		goto st0
tr4:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line rfc5424/parser.go:742
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
tr49:
//line rfc5424/machine.rl:4

    cr.Add(data[p])

	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line rfc5424/parser.go:761
		if data[p] == 62 {
			goto tr5
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto tr2
		}
		goto st0
tr50:
//line rfc5424/machine.rl:82
 {goto st1 } 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line rfc5424/parser.go:778
		switch data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto tr50
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 49:
//line rfc5424/machine.rl:20

    t, e := time.Parse(time.RFC3339Nano, data[poss["timestamp:ini"]:p])
    if e != nil {
        err = fmt.Errorf("error %s [col %d:%d]", e, poss["timestamp:ini"], p);
        p--
 {goto st50 }
    } else {
        timestamp = &t
    }

		case 6:
//line rfc5424/machine.rl:30

    err = fmt.Errorf("error parsing <nilvalue>");

//line rfc5424/parser.go:858
		}
	}

	_out: {}
	}

//line rfc5424/parser.rl:33


    if cs < rfc5424_first_final {
      return nil, err
    }

    return &SyslogMessage{
      Header: Header{
        Pri: Pri{
          Prival: *prival,
        },
        Version: *version,
        Timestamp: timestamp,
      },
    }, nil
}
