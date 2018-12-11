package rfc6587

import (
	"io"

	"fmt"
	syslog "github.com/influxdata/go-syslog"
	"testing"
	"time"
)

func TestMeToo(t *testing.T) {
	messages := []string{
		"<2>1 - - - - - - A\nB",
		"<1>1 -",
		"<1>1 - - - - - - A\nB\nC\nD",
	}

	r, w := io.Pipe()

	go func() {
		defer w.Close()

		for _, m := range messages {
			w.Write([]byte(m))  // message (containing trailers to be interpreted as message)
			w.Write([]byte{10}) // trailer
			time.Sleep(time.Millisecond * 1)
		}
	}()

	results := make(chan syslog.Result)
	ln := func(x syslog.Result) {
		fmt.Println("EMIT", x)
		results <- x
	}

	go func() {
		for {
			fmt.Println("RECV", <-results)
		}
	}()

	New(WithListener(ln)).Parse(r)

	close(results)
	r.Close()
}
