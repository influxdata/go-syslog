package rfc5424

// Version represents the version value
type Version struct {
	Value int
}

// NewVersion constructs a complete Version starting from its value
func NewVersion(value int) *Version {
	return &Version{
		Value: value,
	}
}
