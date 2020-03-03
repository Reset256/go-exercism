package triangle

import (
	"math"
)

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	switch {
	case a <= 0 || b <= 0 || c <= 0 || (a+b < c) || (b+c < a) || (a+c < b) || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || a == pinf || a == ninf || b == pinf || b == ninf || c == pinf || c == ninf:
		k = 0
	case a == b && b == c:
		k = 1
	case a == b || b == c || c == a:
		k = 2
	default:
		k = 3
	}
	return k
}
