package rfc5424

import (
	"testing"
)

func TestDumb(t *testing.T) {
	line := "<101>"
	msg := Parse(line)
	if msg.Prival.Value != 101 {
		t.Errorf("Prival value not matching")
	}
}
