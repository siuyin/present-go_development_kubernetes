// Package time provides the current time.
package time

import (
	gt "time" // go time package, imported as gt
)

// Nower types have a Now method.
type Nower interface {
	Now() gt.Time
}

// Clock implements the Now() method.
type Clock struct{}

// Now returns the current time.
func (c Clock) Now() gt.Time {
	return gt.Now()
}

// Now returns the current time as a string "15:04:05".
func Now(c Nower) string {
	return c.Now().Format("15:04:05")
}
