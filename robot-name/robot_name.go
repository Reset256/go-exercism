package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var alreadySeen = make(map[string]bool)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		for {
			robotName := generateRobotName()
			if alreadySeen[robotName] != true {
				alreadySeen[robotName] = true
				r.name = robotName
				break
			}
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateRobotName() string {
	rand.Seed(time.Now().UnixNano())
	l1 := rand.Int31n(26)
	l2 := rand.Int31n(26)
	return fmt.Sprintf("%v%v%03d", letters[l1:l1+1], letters[l2:l2+1], rand.Int31n(999))
}
