package timespent_test

import (
	"testing"

	"github.com/zwindler/gocastle/pkg/timespent"
)

func TestIncrement(t *testing.T) {
	timespent.Set(0)
	timespent.Increment(10)
	if timespent.Get() != 10 {
		t.Errorf("Expected timeSinceBegin to be 10, but got %d", timespent.Get())
	}
}

func TestFormatDurationShortFormat(t *testing.T) {
	timespent.Set(60 * 60 * 24 * 2) // 2 days
	durationStr := timespent.FormatDuration(timespent.ShortFormat)
	if durationStr != "2d0:0:0" {
		t.Errorf("Expected duration string to be '2d0:0:0', but got '%s'", durationStr)
	}
}

func TestFormatDurationLongFormat(t *testing.T) {
	timespent.Set(60 * 60 * 24 * 2) // 2 days
	durationStr := timespent.FormatDuration(timespent.LongFormat)
	if durationStr != "2 day(s), 0 hour(s), 0 minute(s), 0 second(s)" {
		t.Errorf("Expected duration string to be '2 day(s), 0 hour(s), 0 minute(s), 0 second(s)', but got '%s'", durationStr)
	}
}

func TestFormatDurationInvalidDuration(t *testing.T) {
	timespent.Set(-1)
	durationStr := timespent.FormatDuration(timespent.ShortFormat)
	if durationStr != "Invalid duration" {
		t.Errorf("Expected duration string to be 'Invalid duration', but got '%s'", durationStr)
	}
}
