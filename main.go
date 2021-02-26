package main

// Length is a float64. It is used by this package to calculate the metric conversions.
type Length float64

const (
	// Metre represents a meter in the metric system. It is the base unit.
	Metre Length = 1
	// Attometre represents an attometre in the metric system (am). It is the smaller unit.
	Attometre Length = 100000000
	// Femtometre represents a femtometre (fm).
	Femtometre Length = 10000000
	// Picometre represents a picometre (pm).
	Picometre Length = 1000000
	// Nanometre represents a nanometre (nm).
	Nanometre Length = 100000
	// Micrometre represents a micrometre in the metric system.
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
