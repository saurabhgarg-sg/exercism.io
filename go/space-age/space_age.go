package space

import (
	"fmt"
	"strconv"
)

// Planet holds the name of the planet
type Planet string

// Age calculates the age on different planets
func Age(secs float64, name Planet) float64 {
	factor := 1.0
	switch name {
	case "Earth":
		// No more conversions are needed for Earth
	case "Mercury":
		factor = 0.2408467
	case "Venus":
		factor = 0.61519726
	case "Mars":
		factor = 1.8808158
	case "Jupiter":
		factor = 11.862615
	case "Saturn":
		factor = 29.447498
	case "Uranus":
		factor = 84.016846
	case "Neptune":
		factor = 164.79132
	}
	return calculateAge(secs, factor)
}

// calculateAgeFactor performs divison and rounds up the result to two decimal digits
func calculateAge(secs float64, factor float64) float64 {
	age := fmt.Sprintf("%.2f", (secs / (factor * 31557600)))
	roundedAge, _ := strconv.ParseFloat(age, 64)
	return roundedAge
}
