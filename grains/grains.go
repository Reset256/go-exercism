package grains

import (
	"errors"
	"sync"
)

var once sync.Once
var total uint64

func Square(input int) (numberOnSquare uint64, err error) {
	if input < 1 || input > 64 {
		return 0, errors.New("not an appropriate number")
	}
	numberOnSquare = 1 << (input - 1)
	return
}

func Total() uint64 {
	once.Do(countTotal)
	return total
}

func countTotal() {
	total = 1<<64 - 1
}
