package syslog

import (
  "fmt"
  "os"
  "github.com/influxdata/go-syslog/syslog/rfc5424"
  "github.com/davecgh/go-spew/spew"
)
 
%%{

machine rfc5424;

# 0..191
prival = (('1' ( '9' ( '0'..'1' ){,1} | '0'..'8' ( '0'..'9' ){,1} ){,1}) | ( '2'..'9' ('0'..'9'){,1} )) | '0';

write data;

}%%

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

    %%{
        action acc_prival { privalChars = append(privalChars, fc) }
        action int_prival { prival = rfc5424.NewPrival(utf8ToNum(privalChars)) }

        pri = '<' . prival @acc_prival %int_prival . '>';

        main := pri;
    }%%

    
    %%write init;
    %%write exec;
    if cs < rfc5424_first_final {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error"))
    }

    spew.Dump(prival)
}