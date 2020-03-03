package scale

import (
	"strings"
)

var shifts = map[string]int{
	"m": 1,
	"M": 2,
	"A": 3,
}

func Scale(tonic string, interval string) []string {
	var result []string
	var currentRange []string
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		currentRange = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		currentRange = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}
	position := FindTonicPositionInSlice(FormatTonic(tonic), currentRange)
	finalRange := append(currentRange[position:], currentRange[0:position]...)
	if interval == "" {
		return finalRange
	}
	sliceOfIndicies := []int{
		0,
	}
	for _, value := range interval {
		i := sliceOfIndicies[len(sliceOfIndicies)-1] + shifts[string(value)]
		if i < len(finalRange) {
			sliceOfIndicies = append(sliceOfIndicies, i)
		}
	}
	for _, value := range sliceOfIndicies {
		result = append(result, finalRange[value%len(finalRange)])
	}
	return result
}

func FindTonicPositionInSlice(tonic string, slice []string) int {
	for index, value := range slice {
		if tonic == value {
			return index
		}
	}
	return -1
}
func FormatTonic(tonic string) string {
	if len(tonic) == 2 {
		return strings.ToUpper(tonic[0:1]) + tonic [1:]
	} else {
		return strings.ToUpper(tonic)
	}

}
