package twofer

import "fmt"

func ShareWith(name string) string {
	return fmt.Sprintf("One for %v, one for me.", DefaultIfEmpty(name, "you.. "))
}

func DefaultIfEmpty(s string, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}