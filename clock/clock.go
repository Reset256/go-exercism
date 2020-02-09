package clock

import (
	"fmt"
)

type Clock struct {
	totalMinutes int
}

func New(h int, m int) Clock {
	totalMinutes := 60*h + m
	return New2(totalMinutes)
}

func New2(totalMinutes int) Clock {
	if totalMinutes < 0 {
		totalMinutes = 1440 + (totalMinutes % 1440)
	} else if totalMinutes > 1439 {
		totalMinutes = totalMinutes % 1440
	}
	return Clock{
		totalMinutes: totalMinutes,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.totalMinutes / 60, c.totalMinutes % 60)
}

func (c Clock) Add(m int) Clock {
	return New2(c.totalMinutes + m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}
