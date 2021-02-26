package speed

import (
	main "github.com/dragonDScript/golang-metric-system"
)

const (
	// MetresPerSecond represents M/S
	MetresPerSecond main.Length = 1
	// KilometresPerHour represents M/S
	KilometresPerHour main.Length = (MetresPerSecond * 18) / 5
)

// MpsToKmh converts m/s to km/h
func MpsToKmh(origin main.Length) main.Length {
	return (origin * 18) / 5
}

// KmhToMps converts km/h to m/s
func KmhToMps(origin main.Length) main.Length {
	return (origin / 18) * 5
}
