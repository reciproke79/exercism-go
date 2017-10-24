package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (diff int, e error) {
	if len(a) != len(b) {
		return 0, errors.New("Strand size is different")
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
