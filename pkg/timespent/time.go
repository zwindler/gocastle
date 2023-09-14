// Package timespent provides functionality for working with time spent on tasks.
package timespent

import (
	"fmt"
)

// Format represents the format of the time spent.
type Format string

// ShortFormat and LongFormat are the two possible formats for the time spent.
const (
	ShortFormat Format = "short"
	LongFormat  Format = "long"
)

var timeSinceBegin int = 0

// Increment is a helper function to handle in-game time.
func Increment(value int) {
	timeSinceBegin += value
}

// FormatDuration is a helper function that returns a string to get nice logs.
func FormatDuration(format Format) string {
	if Get() < 0 {
		return "Invalid duration"
	}

	seconds := Get()

	days := seconds / (60 * 60 * 24)
	seconds -= days * (60 * 60 * 24)

	hours := seconds / (60 * 60)
	seconds -= hours * (60 * 60)

	minutes := seconds / 60
	seconds -= minutes * 60

	if format == LongFormat {
		durationStr := ""
		if days > 0 {
			durationStr += fmt.Sprintf("%d day(s), ", days)
		}
		if hours > 0 {
			durationStr += fmt.Sprintf("%d hour(s), ", hours)
		}
		if minutes > 0 {
			durationStr += fmt.Sprintf("%d minute(s), ", minutes)
		}
		durationStr += fmt.Sprintf("%d second(s)", seconds)
		return durationStr
	}

	return fmt.Sprintf("%dd%d:%d:%d", days, hours, minutes, seconds)
}

// Get returns the timeSinceBegin variable.
func Get() int {
	return timeSinceBegin
}

// Set sets the timeSinceBegin variable.
func Set(value int) {
	timeSinceBegin = value
}
