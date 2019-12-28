package monitor

import (
	"fmt"
	"time"
)

type Snapshot struct {
	id        int
	start     time.Time
	end       time.Time
	successes int
	failures  int
}

func (s Snapshot) String() string {
	return fmt.Sprintf("[%s-%s] posted: %d failed: %d", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.successes, s.failures)
}
