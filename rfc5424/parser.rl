package rfc5424

import (
  "fmt"
  "os"
)
 
%%{
machine rfc5424;
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

func parse(data string) Message {
    cs, p, pe := 0, 0, len(data)

    privalChars := []uint8{}
    var prival *Prival

    %%{
      include rfc5424 "machines.rl";
      main := pri;
      write init;
      write exec;
    }%%

    if cs < rfc5424_first_final {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error"))
    }

    return Message{
      Prival: *prival,
    }
}
