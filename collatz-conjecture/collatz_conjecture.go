package collatzconjecture

import "github.com/pkg/errors"

func CollatzConjecture(input int) (int, error) {
	var result int
	if input < 1 {
		return 0, errors.New("Error!")
	}
	for {
		if IsPowerOfTwo(input) {
			result += GetPowerOfTwo(input)
			return result, nil
		}
		if !IsEven(input) {
			input = input*3 + 1
		} else {
			input >>=  1
		}
		result++
	}

}

func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}

func IsEven(n int) bool {
	return n & 1 == 0
}

func GetPowerOfTwo(n int) int {
	result := 0
	if n > 1 {
		for {
			n = n >> 1
			result++
			if n == 1 {
				return result
			}
		}
	} else {
		return result
	}
}
