package main

// Convert converts a certain unit to a certain unit.
func Convert(origin Length, to Length) Length {
	if origin > Metre {
		return to * origin
	}
	if origin < Metre {
		return to / origin
	}
	return 0
}
