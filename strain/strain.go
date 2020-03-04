package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(predicate func(int) bool) Ints {
	if i == nil {
		return i
	}
	result := make(Ints, 0)
	for _, value := range i {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
func (i Ints) Discard(predicate func(int) bool) Ints {
	if i == nil {
		return i
	}
	result := make(Ints, 0)
	for _, value := range i {
		if !predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func (i Lists) Keep(predicate func([]int) bool) Lists {
	if i == nil {
		return i
	}
	result := make(Lists, 0)
	for _, value := range i {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func (i Strings) Keep(predicate func(string) bool) Strings {
	if i == nil {
		return i
	}
	result := make(Strings, 0)
	for _, value := range i {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
