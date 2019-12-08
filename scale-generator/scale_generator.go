// Package scale generates musical scales.
// Thanks to skuridin's solution on this for helping me out.
package scale

import (
	"strings"
)

// Scale generates the musical scale starting with the tonic
// and following the specified interval pattern
func Scale(tonic string, gap string) []string {
	retScale := []string{}
	scl := notes(tonic)
	tonic = strings.Title(tonic)

	for i, s := range scl {
		if s == tonic {
			scl = append(scl[i:], scl[:i]...)
		}
	}

	if gap == "" {
		gap = "mmmmmmmmmmmm"
	}

	for _, s := range strings.Split(gap, "") {
		retScale = append(retScale, scl[0])

		switch s {
		case "m":
			scl = scl[1:]
		case "M":
			scl = scl[2:]
		case "A":
			scl = scl[3:]
		}
	}

	return retScale
}

func notes(tonic string) []string {
	noteScales := []string{}
	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		noteScales = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		noteScales = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}

	return noteScales
}
