package diffsquares

import "math"

func Difference(n int) int {
	var sums int
	var sumOfSquares int
	for i := 1; i <= n; i++ {
		sums += i
		sumOfSquares += int(math.Pow(float64(i), 2))
	}

	return int(math.Pow(float64(sums), 2)) - sumOfSquares
}
func SumOfSquares(n int) int {
	var sumOfSquares int
	for i := 1; i <= n; i++ {
		sumOfSquares += int(math.Pow(float64(i), 2))
	}

	return sumOfSquares
}

func SquareOfSum(n int) int {
	var sums int
	for i := 1; i <= n; i++ {
		sums += i
	}
	return int(math.Pow(float64(sums), 2))
}
