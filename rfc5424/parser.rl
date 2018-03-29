package rfc5424

import (
  "fmt"
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

func parse(data string) (*Message, error) {
    cs, p, pe := 0, 0, len(data)

    privalChars := []uint8{}
    versionChars := []uint8{}
    var prival *Prival
    var version *Version

    %%{
      include rfc5424 "machine.rl";
      main := header;
      write init;
      write exec;
    }%%

    if cs < rfc5424_first_final {
        return nil, fmt.Errorf("error")
    } 

    return &Message{
      Prival: *prival,
      Version: *version,
    }, nil
}
