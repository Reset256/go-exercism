package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	expression, _ := regexp.Compile("[^a-z0-9]+")
	formatted := expression.ReplaceAllString(strings.ToLower(pt), "")
	lenOfFormatted := len(formatted)
	rectangleSide1 := int(math.Sqrt(float64(lenOfFormatted)))
	var rectangleSide2 int
	if lenOfFormatted/rectangleSide1 == rectangleSide1 && lenOfFormatted%rectangleSide1 == 0 {
		rectangleSide2 = rectangleSide1
	} else {
		rectangleSide2 = rectangleSide1 + 1
	}
	result := ""
	for x := 0; x < rectangleSide2; x++ {
		for y := 0; y < rectangleSide1; y++ {
			index := y*rectangleSide2 + x
			fmt.Println("x - ", x, "y - ", y, "index - ", index)
			if index < lenOfFormatted {
				result += formatted[index : index+1]
			} else {
				result += " "
			}
		}
		result += " "
	}
	result = strings.TrimSpace(result)
	//result = strings.Replace(result, "  ", " ", rectangleSide1)
	//return result[0 : len(result)-1]
	//resultLen := len(result)
	if rectangleSide1 == rectangleSide2 {
		return result + " "
	}
	return result

}
