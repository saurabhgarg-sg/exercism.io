// Package acronym converts a given line to corresponding acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate accpets a line input and spits out its acronym.
func Abbreviate(input string) string {

	// split the line on spaces & hyphen
	delimiters := regexp.MustCompile(`(\s+|\-)`)
	tokens := delimiters.Split(input, -1)

	// Parse each word, pick first letter seen & append it to output
	re := regexp.MustCompile(`[[:alpha:]]{1}`)
	output := make([]string, 0, 30)
	for _, token := range tokens {
		match := re.FindString(token)

		// Ignore if there is no letter in this token
		if match != "" {
			output = append(output, match)
		}
	}

	// Capitalize the output before returning it to the user
	return strings.ToUpper(strings.Join(output, ""))
}
