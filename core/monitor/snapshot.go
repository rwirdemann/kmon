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
	return fmt.Sprintf("%d: [%s-%s] ok: %d nok: %d", s.id, s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.successes, s.failures)
}
