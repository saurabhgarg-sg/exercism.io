// Package gigasecond implements adding gigasecond to age
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond add 1 giga second to the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
