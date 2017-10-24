package raindrops

import (
	"strconv"
	"strings"
)

const testVersion = 3

var factors = map[string]int{
	"factor3": 3,
	"factor5": 5,
	"factor7": 7,
}

func Convert(i int) string {
	var strs []string

	for k, v := range factors {
		if i%v == 0 {
			switch k {
			case "factor3":
				{
					strs = append(strs, "Pling")
				}
			case "factor5":
				{
					strs = append(strs, "Plang")
				}
			case "factor7":
				{
					strs = append(strs, "Plong")
				}
			}
		}
	}
	if strs == nil {
		return strconv.Itoa(i)
	}
	return strings.Join(strs, "")
}
