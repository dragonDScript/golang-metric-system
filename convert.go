package metric

// Convert converts a certain unit to a certain unit.
func Convert(origin Length, to Length) Length {
	return to * origin
}
