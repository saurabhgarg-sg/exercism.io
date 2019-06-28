// Package triangle classifies a given triangle as equilateral, isosceles, or scalene
package triangle

import (
	"math"
)

// Kind is the type of triangle found
type Kind string

const (
	// NaT is not a triangle
	NaT = "NaT"
	// Equ is a triage with all 3 sides equal
	Equ = "Equ"
	// Iso is a triage with 2 of 3 sides equal
	Iso = "Iso"
	// Sca is a triangle all three different sides
	Sca = "Sca"
)

// KindFromSides checks the edges to classify the triangle
func KindFromSides(a, b, c float64) Kind {
	if (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) ||
		(math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)) ||
		(a <= 0 || b <= 0 || c <= 0) ||
		(a+b < c || a+c < b || b+c < a) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
