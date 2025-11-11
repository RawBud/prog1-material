package length

type Length int

// FromCentimeters konstruiert eine Länge aus einer Zentimetern-Angabe.
func FromCentimeters(c int) Length {
	return Length(c)
}

// FromMeters konstruiert eine Länge aus einer Meter-Angabe.
func FromMeters(m int) Length {
	return Length(m * 100)
}

// FromKilometers konstruiert eine Länge aus einer Kilometer-Angabe.
func FromKilometers(k int) Length {
	return Length(k * 1000 * 100)
}

// Centimeters gibt die Länge in Zentimetern zurück.
func (l Length) Centimeters() int {
	return int(l)
}

// Meters gibt die Länge in Metern zurück.
func (l Length) Meters() int {
	return int(l) / 100
}

// Kilometers gibt die Länge in Kilometern zurück.
func (l Length) Kilometers() int {
	return int(l) / (1000 * 100)
}
