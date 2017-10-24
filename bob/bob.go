package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(s string) string {
	// trim s
	s = strings.TrimSpace(s)

	switch {
	case s == "":
		return "Fine. Be that way!"
	case check(s, unicode.IsUpper) && !check(s, unicode.IsLower):
		return "Whoa, chill out!"
	case s[len(s)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}

}

// Checks if strings passes some test
func check(s string, test func(rune) bool) bool {
	for _, string := range s {
		if test(string) {
			return true
		}
	}
	return false
}
