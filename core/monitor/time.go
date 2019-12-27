package monitor

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func extractTime(s string) (time.Time, error) {
	split := strings.Split(s, " ")
	for _, s := range split {
		if strings.HasPrefix(s, "time") {
			split := strings.Split(s, "=")
			t, err := time.Parse(time.RFC3339, split[1][1:len(split[1])-1])
			if err != nil {
				return time.Now(), err
			}
			return round(t, time.Second), nil
		}
	}
	return time.Now(), errors.New(fmt.Sprintf("string contains no time [s=%s]", s))
}

func round(t time.Time, d time.Duration) time.Time {
	if d == time.Second {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
	}
	return t
}
