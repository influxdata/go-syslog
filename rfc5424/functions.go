package rfc5424

// unsafeUTF8DecimalCodePointsToInt converts a slice containing
// a series of UTF-8 decimal code points into their integer rapresentation.
//
// It assumes input code points are in the range 48-57.
// Returns a pointer since an empty slice is equal to nil and not to the zero value of the codomain (ie., `int`).
func unsafeUTF8DecimalCodePointsToInt(chars []uint8) int {
	out := 0
	ord := 1
	for i := len(chars) - 1; i >= 0; i-- {
		curchar := int(chars[i])
		out += (curchar - '0') * ord
		ord *= 10
	}
	return out
}

func stripSlashes(str string) string {
	l := len(str)
	buf := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		buf = append(buf, str[i])
		if l > i+1 && str[i+1] == 92 {
			i++
		}
	}

	return string(buf)
}
