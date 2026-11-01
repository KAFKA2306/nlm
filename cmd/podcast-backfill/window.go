package main

import (
	"fmt"
	"strings"
	"time"
)

// parseUntil returns an exclusive upper bound. Date-only values include the
// entire JST calendar day; RFC3339 values include the specified instant.
func parseUntil(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if t, err := time.Parse(time.RFC3339Nano, value); err == nil {
		return t.Add(time.Nanosecond), nil
	}
	jst := time.FixedZone("JST", 9*60*60)
	if t, err := time.ParseInLocation("2006-01-02", value, jst); err == nil {
		return t.AddDate(0, 0, 1), nil
	}
	return time.Time{}, fmt.Errorf("invalid --until %q: use YYYY-MM-DD or RFC3339", value)
}

func inWindow(when, since, until time.Time) bool {
	return !when.Before(since) && (until.IsZero() || when.Before(until))
}
