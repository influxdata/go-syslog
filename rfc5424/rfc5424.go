package rfc5424

func Parse(data string) (*Message, error) {
	return parse(data)
}
