package listops

type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int


func (l IntList) Foldl(fn binFunc, initial int) int {
	for _ ,value := range l {
		initial = fn(initial, value)
	}

	return initial
}

func (l IntList) Foldr(function binFunc, initial int) int {
	for i := len(l) - 1; i>=0; i-- {
		initial = function(l[i], initial)
	}
	return initial
}


func (l IntList) Filter(fn predFunc) IntList {
	result := make(IntList, 0)
	for _, value := range l {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}

func (l IntList) Concat(args []IntList) IntList {
	result := make(IntList, 0)
	for _, value := range l {
		result = append(result, value)
	}
	for _, value := range args {
		for _, innerValue := range value {
			result = append(result, innerValue)
		}
	}
	return result
}

func (l IntList) Length() int {
	return len(l)
}

func (l IntList) Append(this IntList) IntList {
	for _, value := range this {
		l = append(l, value)
	}
	return l
}

func (l IntList) Reverse() IntList {
	result := make(IntList, len(l))
	for i := range l {
		result[i] = l[len(l)-1-i]
	}
	return result
}

func (l IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, len(l))
	for i, value := range l {
		result[i] = fn(value)
	}
	return result
}

