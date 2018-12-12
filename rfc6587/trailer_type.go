package rfc6587

import (
	"fmt"
)

// TrailerType is the king of supported trailers for non-transparent frames.
type TrailerType int

const (
	// LF is the line feed - ie., byte 10. Also the default one.
	LF TrailerType = iota
	// NUL is the nul byte - ie., byte 0.
	NUL
)

var names = [...]string{"LF", "NUL"}
var bytes = []int{10, 0}

func (t TrailerType) String() string {
	if t < LF || t > NUL {
		return ""
	}

	return names[t]
}

// Value returns the byte corresponding to the receiving TrailerType.
func (t TrailerType) Value() (int, error) {
	if t < LF || t > NUL {
		return -1, fmt.Errorf("unknown TrailerType")
	}

	return bytes[t], nil
}
