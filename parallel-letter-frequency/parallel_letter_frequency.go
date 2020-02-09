package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	c := make(chan FreqMap, len(strings))
	for _, value := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(value)
	}
	result := FreqMap{}
	for range strings {
		result.merge(<-c)
	}
	return result
}

func (f FreqMap) merge(donor FreqMap) FreqMap {
	for key, value := range donor {
		f[key] += value
	}
	return f
}
