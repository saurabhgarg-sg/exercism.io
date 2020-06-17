// Package bob implements decoding of answer patterns of a teenager
package bob

import "strings"

const (
	chars = "abcdefghijklmnopqrstuvwxyz"
)

// Hey formats the answers based upon the speech pattern
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	if len(remark) < 1 {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if matchCase(remark) {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	} else if matchCase(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func matchCase(r string) bool {
	return strings.ToUpper(r) == r && strings.IndexAny(strings.ToLower(r), chars) != -1
}