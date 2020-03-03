package bob

import "strings"

func Hey(remark string) string {
	space := strings.TrimSpace(remark)
	switch {
	case isEmptyString(space):
		return "Fine. Be that way!"
	case isCapital(space) && isQuestion(space):
		return "Calm down, I know what I'm doing!"
	case isQuestion(space):
		return "Sure."
	case isCapital(space):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isCapital(remark string) bool {
	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1:] == "?"
}

func isEmptyString(s string) bool {
	return len(s)==0
}
