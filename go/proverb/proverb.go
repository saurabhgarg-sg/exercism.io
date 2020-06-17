// Package proverb generates the age-old rhyme from given set of words.
package proverb

import "fmt"

// Proverb generates the substituted rhyme
func Proverb(rhyme []string) []string {
	base := [100]string{}
	inputLength := len(rhyme)
	output := base[:0]

	// It returns empty string on empty input
	if inputLength == 0 {
		return output
	}

	if inputLength > 1 {
		for index := 0; index < inputLength-1; index++ {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1]))
		}
	}

	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return output
}
