package chars

import "fmt"

const (
	errInvalidDigitCodepoint = "the provided codepoint does not represent a digit: %d"
)

// UTF8DecimalCodePointsToInt converts a slice containing
// a series of UTF-8 decimal code points into their integer rapresentation.
//
// Returns a pointer since an empty slice is equal to nil and not to the zero value of the codomain (ie., `int`).
func UTF8DecimalCodePointsToInt(chars []uint8) (*int, error) {
	if len(chars) == 0 {
		return nil, nil
	}

	out := 0
	ord := 1
	for i := len(chars) - 1; i >= 0; i-- {
		curchar := int(chars[i])
		if curchar < 48 || curchar > 57 {
			return nil, fmt.Errorf(errInvalidDigitCodepoint, curchar)
		}
		out += (curchar - '0') * ord
		ord *= 10
	}
	return &out, nil
}

// UTF8DecimalCodePointsToString converts a slice containing
// a series of UTF-8 decimal code point into their string representation.
//
// Returns a pointer since an empty slice is equal to nil and not to the zero value of the codomain (ie., `string`).
func UTF8DecimalCodePointsToString(chars []uint8) (*string, error) {
	if len(chars) == 0 {
		return nil, nil
	}

	out := ""
	for i := 0; i < len(chars); i++ {
		out += string(chars[i])
	}

	return &out, nil
}

func UnsafeUTF8DecimalCodePointsToInt(chars []uint8) int {
	out := 0
	ord := 1
	for i := len(chars) - 1; i >= 0; i-- {
		curchar := int(chars[i])
		out += (curchar - '0') * ord
		ord *= 10
	}
	return out
}
