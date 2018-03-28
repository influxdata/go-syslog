package syslog

import (
	"testing"
)

func TestDumb(t *testing.T) {
	line := "<101>"
	RFC5424(line)
}
