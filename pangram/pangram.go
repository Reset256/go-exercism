package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToUpper(input)
	for i := 'A'; i<='Z'; i++ {
		if !strings.Contains(input, string(i)) {
			return false
		}
	}
	return true

}

