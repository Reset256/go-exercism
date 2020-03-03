package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	d, _ := time.ParseDuration("1000000000s")
	t = t.Add(d)
	return t
}
