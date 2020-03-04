package hamming

import (
	"github.com/pkg/errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Lengths are not equal")
	}
	if a == b {
		return 0, nil
	}
	result := 0
	for x, y := range a {
		if uint8(y) != b[x] {
			result++
		}
	}
	return result, nil
}
