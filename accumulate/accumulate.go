package accumulate

func Accumulate(given []string, converter func(string) string) []string {
	result := make([]string, len(given))

	for number, value := range given {
		result[number] = converter(value)
	}
	return result
}
