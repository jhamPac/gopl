package main

// Celsius represents C
type Celsius float64

// Fahrenheit represents F
type Fahrenheit float64

// Kelvin represents K
type Kelvin float64

// CToF converts C to F
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

// FToC converts F to C
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}
