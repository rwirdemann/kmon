package monitor

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const layout = "2006-01-02 15:04:05.999"

func extractTime(s string) (time.Time, error) {
	r := regexp.MustCompile("[0-9-]* [0-9:.]*")
	match := r.FindStringSubmatch(s)
	if match == nil {
		return time.Now(), errors.New(fmt.Sprintf("string contains no time [s=%s]", s))
	}
	t, err := time.Parse(layout, match[0])
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}

func round(t time.Time, d time.Duration) time.Time {
	if d == time.Second {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
	}
	return t
}
