package raindrops

import (
	"strconv"
	"strings"
)

func Convert(input int) string {
	var result strings.Builder
	var m = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	fun := func(input int, divider int, s string) string {
		if input%divider == 0 {
			return s
		}
		return ""
	}
	for i, s := range m {
		result.WriteString(fun(input, i, s))
	}
	if result.Len() == 0 {
		return strconv.Itoa(input)
	}
	return result.String()
}
