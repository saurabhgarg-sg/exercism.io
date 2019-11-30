// Package darts implements scoring scheme for a Dart game.
package darts

import (
	"math"
)

var radii = map[string]float64{
	"inner":  1.0,
	"middle": 5.0,
	"outer":  10.0,
}

var scores = map[string]int{
	"inner":  10,
	"middle": 5,
	"outer":  1,
}

// Score computes points for given dart Cartesian coordinates.
func Score(x float64, y float64) int {

	// As it is concentric circles, equation is simpler
	// Equation: d = sqrt of (sq(|x2 - x1|) - sq(|y2-y1|))
	// location = sqrt(sq(x) + sq(y)) as other coordinates are (0,0)

	sum := math.Pow(x, 2) + math.Pow(y, 2)
	switch location := math.Sqrt(sum); {

	case location <= radii["inner"]:
		return scores["inner"]

	case location <= radii["middle"]:
		return scores["middle"]

	case location <= radii["outer"]:
		return scores["outer"]

	default:
		return 0
	}
}
