package model

import "fmt"

var (
	TimeSinceBegin int = 0
)

func FormatDuration(seconds int) string {
	if seconds < 0 {
		return "Invalid duration"
	}

	days := seconds / (60 * 60 * 24)
	seconds -= days * (60 * 60 * 24)

	hours := seconds / (60 * 60)
	seconds -= hours * (60 * 60)

	minutes := seconds / 60
	seconds -= minutes * 60

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
