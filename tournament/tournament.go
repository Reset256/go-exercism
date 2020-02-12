package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
)

type score struct {
	win  int
	loss int
	draw int
}

var table = make(map[string]score)

var regExp = regexp.MustCompile("([a-zA-Z]+ [a-zA-Z]+);([a-zA-Z]+ [a-zA-Z]+);(win|loss|draw)")

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if !validateRow(line) {
			return errors.New("failed")
		}
		submatch := regExp.FindAllStringSubmatch(line, -1)[0]
		switch submatch[3] {
		case "win":
			win(submatch[1])
			loss(submatch[2])
		case "draw":
			draw(submatch[1])
			draw(submatch[2])
		case "loss":
			win(submatch[2])
			loss(submatch[1])
		}
	}
	keys := make([]string, 0, len(table))
	for k := range table {
		keys = append(keys, k)
	}
	// TODO sort first by points, then by name
	sort.Strings(keys)
	fmt.Printf("%v\n", table)
	result := "Team                           | MP |  W |  D |  L |  P\n"
	for _, key := range keys {
		result += fmt.Sprintf("%-31v|  %v |  %v |  %v |  %v |  %v \n",
			key, table[key].win+table[key].loss+table[key].draw, table[key].win,
			table[key].draw, table[key].loss, table[key].win*3+table[key].draw)
	}
	fmt.Println(result)
	return nil
}

func win(teamName string) {
	score := table[teamName]
	score.win++
	table[teamName] = score
}

func loss(teamName string) {
	score := table[teamName]
	score.loss++
	table[teamName] = score
}

func draw(teamName string) {
	score := table[teamName]
	score.draw++
	table[teamName] = score
}

func validateRow(s string) bool {
	return regExp.MatchString(s)
}
