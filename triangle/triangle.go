package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

type Kind int

// Enumerated list of triangle types
const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

func KindFromSides(a, b, c float64) Kind {
	// Put all values into slice and sort
	s := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(s))

	switch {
	case math.IsNaN(s[0]) || s[0] <= 0 || s[2] == math.Inf(1) ||
		s[1]+s[0] < s[2]:
		return NaT
	case s[0] == s[1] && s[1] == s[2]:
		return Equ
	case s[0] == s[1] || s[1] == s[2]:
		return Iso
	}
	return Sca
}
