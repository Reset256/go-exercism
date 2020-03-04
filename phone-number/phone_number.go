package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Number(input string) (string, error) {
	strippingFunc := func(r rune) rune {
		if r < '0' || r > '9' {
			return -1
		}
		return r
	}
	result := strings.Map(strippingFunc, input)
	if len(regexp.MustCompile(`^1?[2-9]\d{2}[2-9]\d{6}$`).FindAllString(result, -1)) != 1 {
		return "", errors.New("wrong")
	}
	return result[len(result)-10:], nil
}

func Format(input string) (string, error) {
	result, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", result[0:3], result[3:6], result[6:], ), nil

}

func AreaCode(input string) (string, error) {
	result, err := Number(input)
	if err != nil {
		return "", err
	}
	return result[0:3], nil
}
