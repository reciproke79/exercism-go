package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 3

// Scan only for chars from alphabet, ignore case
var regexChars, _ = regexp.Compile(`([a-z][A-Z])`)

// Scan for separators
var regexSpecialChars, _ = regexp.Compile(`[-;:]`)

func Abbreviate(s string) (acronym string) {
	// Clear string of all special chars
	clear := regexChars.ReplaceAllString(s, "${1}")
	fmt.Println(clear)
	clear = regexSpecialChars.ReplaceAllString(clear, " ")
	fmt.Println(clear)

	for _, el := range strings.Split(clear, " ") {
		if el != "" {
			acronym += strings.ToUpper(string(el[0]))
		}
	}
	return acronym
}
