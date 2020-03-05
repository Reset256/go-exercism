package twelve

import "strings"

type Line struct {
	dayNumber string
	verse     string
}

var songMap = map[int]Line{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	result := make([]string, 0)
	for i := 1; i <= 12; i++ {
		result = append(result, Verse(i))
	}
	return strings.Join(result, "\n")
}

func Verse(input int) string {
	verse := ""
	for i := input; i > 0; i-- {
		verse = verse + songMap[i].verse
		switch i {
		case 2:
			verse += ", and "
		case 1:
		default:
			verse += ", "

		}
	}
	return "On the " + songMap[input].dayNumber + " day of Christmas my true love gave to me: " + verse + "."
}
