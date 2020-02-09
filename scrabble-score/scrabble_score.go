package scrabble

import (
	"regexp"
	"strings"
	"sync"
)

var initOnce sync.Once

var letterRegExp = regexp.MustCompile("[A-Z]")
var resultMap = make(map[rune]int)

func initLetterMap() {
	var m = map[int]string{
		1:  "A, E, I, O, U, L, N, R, S, T",
		2:  "D, G",
		3:  "B, C, M, P",
		4:  "F, H, V, W, Y",
		5:  "K",
		8:  "J, X",
		10: "Q, Z",
	}

	for key, value := range m {
		for _, char := range letterRegExp.FindAllString(value, -1) {
			resultMap[[]rune(char)[0]] = key
		}
	}
}

func Score(input string) (result int) {
	initOnce.Do(initLetterMap)
	if len(input) == 0 {
		return 0
	}
	for _, letter := range letterRegExp.FindAllString(strings.ToUpper(input), -1) {
		result += resultMap[[]rune(letter)[0]]
	}
	return
}
