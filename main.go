package metric

// Length is a float64. It is used by this package to calculate the metric conversions.
type Length float64

const (
	// Metre represents a meter in the metric system. It is the base unit.
	Metre Length = 1
	// Micrometre is the smallest unit in this package.
	Micrometre Length = 10000
	// Millimetre represents a millimeter (mm).
	Millimetre Length = 1000
	// Centimetre represents a centimeter (cm).
	Centimetre Length = 100
	// Decimetre represents a decimetre (dm).
	Decimetre Length = 10
	// Decametre represents a decameter in the metric system.
	Decametre Length = 10
	// Hectometre represents an hectometer in the metric system.
	Hectometre Length = 100
	// Kilometre represents a kilometer in the metric system.
	Kilometre Length = 1000
)
