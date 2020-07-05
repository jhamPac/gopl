package main

import (
	"flag"
	"fmt"
)

// Celsius represents °C
type Celsius float64

// Fahrenheit represents °F
type Fahrenheit float64

// Kelvin represents °K
type Kelvin float64

// CToF converts °C to °F
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

// FToC converts °F to °C
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

// KToC converts °K to °C
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.3g°C", c)
}

type celsiusFlag struct{ Celsius }

func (cf *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		cf.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		cf.Celsius = FToC(Fahrenheit(value))
	case "K":
		cf.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}
