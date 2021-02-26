package main

// Convert converts a certain unit to a certain unit.
func Convert(origin Length, to Length) Length {
	if origin < to {
		return to * origin
	}
	if origin > to {
		return to / origin
	}
	return 0
}
