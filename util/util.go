package util

import "fmt"

const (
	errInvalidCodepoint = "the provided codepoint does not represent a digit: %d"
)

// UTF8DecimalCodePointsToInt convert a slice containing
// a series of UTF-8 decimal code points into their integer rapresentation
// Returns a pointer because an empty slice is equal to nil and not to zero
func UTF8DecimalCodePointsToInt(chars []uint8) (*int, error) {
	if len(chars) == 0 {
		return nil, nil
	}
	out := 0
	ord := 1
	for i := len(chars) - 1; i >= 0; i-- {
		curchar := int(chars[i])
		if curchar < 48 || curchar > 57 {
			return nil, fmt.Errorf(errInvalidCodepoint, curchar)
		}
		out += (curchar - '0') * ord
		ord *= 10
	}
	return &out, nil
}
