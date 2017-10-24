package pangram

import (
	"strings"
)

var chars = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

const testVersion = 2

func IsPangram(s string) bool {
	var charCount = make(map[string]int)

	for _, char := range s {
		for _, c := range chars {
			if strings.ToLower(string(char)) == c {
				charCount[string(char)]++
			}
		}
	}

	if len(charCount) >= 26 {
		return true
	}
	return false
}
