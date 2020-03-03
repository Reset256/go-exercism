package series

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func All(n int, s string) []string  {
	result := make([]string, len(s) +1 -n)
	for i := 0; i <= len(s) - n; i++ {
		result[i] = s[i:i+n]
	}
	return result
}