package monitor

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func extractTime(s string) (time.Time, error) {
	r := regexp.MustCompile("time=\"[0-9-T:+]*\"")
	match := r.FindStringSubmatch(s)
	if match == nil {
		return time.Now(), errors.New(fmt.Sprintf("string contains no time [s=%s]", s))
	}

	v := strings.ReplaceAll(strings.Split(match[0], "=")[1], "\"", "")
	t, err := time.Parse(time.RFC3339, v)
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
