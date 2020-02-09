package luhn

import "strings"

func Valid(input string) bool {
	runes := []rune(strings.ReplaceAll(input, " ", ""))
	lenght := len(runes)
	if lenght < 2 {
		return false
	}
	sum := 0
	for i := lenght - 2; i > -1; i -= 2 {
		i2 := int(runes[i] - '0')
		if i2 > 9 || i2 < 0 {
			return false
		}
		current := i2 * 2
		if current > 9 {
			current -= 9
		}
		sum += current
	}

	for i := lenght - 1; i > -1; i -= 2 {
		current := int(runes[i] - '0')
		if current > 9 || current < 0 {
			return false
		}
		sum += current
	}

	return sum%10 == 0

}
