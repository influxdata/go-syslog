package rfc5424

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type testCase struct {
	input string
	valid bool
	// value SyslogMessage
}

func TestDumb(t *testing.T) {
	line := "<101>122 "
	msg, err := Parse(line)
	if err != nil {
		t.Error(err)
		return
	}

	spew.Dump(msg)

	// if msg.Prival.Value != 101 {
	// 	t.Errorf("Prival value not matching")
	// }
}
