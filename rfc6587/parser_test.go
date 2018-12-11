package rfc6587

import (
	"io"
	"sync"

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

	wg := sync.WaitGroup{}

	r, w := io.Pipe()

	go func() {
		defer w.Close()

		for _, m := range messages {
			wg.Add(1)
			w.Write([]byte(m))  // message (containing trailers to be interpreted as message)
			w.Write([]byte{10}) // trailer
			time.Sleep(time.Second * 1)
		}
	}()

	results := make(chan *syslog.Result)
	defer close(results)
	ln := func(x *syslog.Result) {
		fmt.Println("EMIT", x)
		results <- x
	}

	go func() {
		for {
			fmt.Println("RECV", <-results)
			wg.Done()
		}
	}()

	NewParser(WithListener(ln)).Parse(r)

	wg.Wait()
	r.Close()
}
