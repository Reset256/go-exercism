package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	//fields := strings.Fields(s)
	allString := regexp.MustCompile("[a-zA-Z']+").FindAllString(s, -1)
	result := []string{}
	for i := range allString {
		result = append(result, strings.ToUpper(getFirstLetter(allString[i])))
	}
	return strings.Join(result, "")
}

func getFirstLetter(s string) string {
	for i := range s {
		if unicode.IsLetter(rune(s[i])) {
			return string(s[i])
		}
	}
	return ""
}
