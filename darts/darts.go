package darts

import "math"

func Score(x float64, y float64) int {

	result := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case result <= 1:
		return 10
	case result <= 5:
		return 5
	case result <= 10:
		return 1
	default:
		return 0
	}

}
