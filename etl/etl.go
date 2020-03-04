package etl

import "strings"

func Transform(m map[int][]string) map[string]int {
	result := make (map[string]int, 0)
	for score, letterList := range m{
		for _, letter := range letterList {
			result[strings.ToLower(letter)]=score
		}
	}
	return result
}
