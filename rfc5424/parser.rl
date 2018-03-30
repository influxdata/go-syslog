package rfc5424

import (
  "fmt"
  "time"
)
 
%%{
machine rfc5424;
write data;
}%%

func Parse(data string) (*SyslogMessage, error) {
    cs, p, pe, eof := 0, 0, len(data), len(data)

    _ = eof

    cr := GetCharsRepo()

    poss := make(map[string]int, 0)

    err := fmt.Errorf("generic error")

    var prival *Prival
    var version *Version
    var timestamp *time.Time

    %%{
      include rfc5424 "machine.rl";
      
      write init;
      write exec;
    }%%

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
