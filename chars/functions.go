package chars

// UnsafeUTF8DecimalCodePointsToInt converts a slice containing
// a series of UTF-8 decimal code points into their integer rapresentation.
//
// It assumes input code points are in the range 48-57.
// Returns a pointer since an empty slice is equal to nil and not to the zero value of the codomain (ie., `int`).
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
