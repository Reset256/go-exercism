package brackets

const (
	p_o = "("
	p_c = ")"
	c_o = "{"
	c_c = "}"
	b_o = "["
	b_c = "]"
)

func Bracket(input string) bool {
	stack := make([]string, 0)

	for i := 0; i < len(input); i++ {
		current := input[i : i+1]
		if current == p_o || current == c_o || current == b_o {
			stack = append(stack, current)
		} else if current == p_c || current == c_c || current == b_c {
			switch {
			case current == p_c && (getLast(stack) == p_o):
				stack = removeLast(stack)
			case current == c_c && (getLast(stack) == c_o):
				stack = removeLast(stack)
			case current == b_c && (getLast(stack) == b_o):
				stack = removeLast(stack)
			default:
				return false
			}
		}
	}
	return len(stack) == 0

}

func getLast(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[len(stack)-1:][0]
}

func removeLast(stack []string) []string {
	if len(stack) == 1 {
		return make([]string, 0)
	}
	return stack[0 : len(stack)-1 ]
}
