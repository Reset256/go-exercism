package isbn

import (
	"regexp"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	pattern := `^\d-?\d{3}-?\d{5}-?[\dX]$`
	matchString, _ := regexp.MatchString(pattern, isbn)
	if !matchString {
		return false
	}
	cleared := ""
	for i, char := range isbn {
		if (char >= '0' && char <= '9') || char == 'X' {
			cleared += isbn[i : i+1]
		}
	}

	var calc int

	for i, value := range Reverse(cleared) {
		if value == 'X' {
			calc += 10 * (i + 1)
		} else {
			atoi, _ := strconv.Atoi(string(value))
			calc += atoi * (i + 1)
		}
	}

	return calc%11 == 0

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
