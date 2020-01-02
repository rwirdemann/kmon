package monitor

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Snapshot struct {
	id        int
	start     time.Time
	end       time.Time
	successes int
	failures  int
	posted    int
	lastlog   string
}

func (s *Snapshot) Process(line string) {
	if strings.Contains(line, "Published job") {
		s.posted++
	}

	if strings.Contains(line, "status=200") {
		s.successes++
	}
	if strings.Contains(line, "status=400") {
		s.failures++
	}
}

func (s *Snapshot) Log(limit float64) {
	if s.posted == 0 {
		return
	}

	rate := float64(s.failures) / float64(s.posted)
	if rate > limit {
		current := fmt.Sprintf("[%s-%s] error rate execeeds limit [posted=%d, failed=%d, limit=%.3f, rate=%.3f]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, limit, rate)
		if current != s.lastlog {
			logrus.Error(current)
			s.lastlog = current
		}
		return
	}

	if s.failures > 0 {
		current := fmt.Sprintf("[%s-%s] error rate increases [posted=%d, failed=%d, limit=%.3f, rate=%.3f]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, limit, rate)
		if current != s.lastlog {
			logrus.Warn(current)
			s.lastlog = current
		}
		return
	}

	current := fmt.Sprintf("[%s-%s] all fine [posted=%d, failed=%d, limit=%.3f, rate=%.3f]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, limit, rate)
	if current != s.lastlog {
		logrus.Info(current)
		s.lastlog = current
	}
}
