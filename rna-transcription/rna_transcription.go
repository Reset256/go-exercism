package strand

func ToRNA(dna string) string {
	var result string
	for i := 0; i < len(dna); i++ {
		curr := dna[i : i+1]
		switch {
		case curr == "G":
			result += "C"
		case curr == "C":
			result += "G"
		case curr == "T":
			result += "A"
		case curr == "A":
			result += "U"
		}
	}
	return result
}
