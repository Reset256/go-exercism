package isogram

import (
	"regexp"
	"strings"
)

var letterRegExp = regexp.MustCompile("[A-Z]")

func IsIsogram(input string) (result bool) {

	result = true
	var resultMap = make(map[string]int)
	for _, letter := range letterRegExp.FindAllString(strings.ToUpper(input), -1) {
		resultMap[letter] += 1
	}

	for _, value := range resultMap {
		if value > 1 {
			result = false
			break
		}
	}
	return
}
