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

func (s *Snapshot) Log() {
	if s.posted == 0 {
		return
	}

	if float32(s.failures/s.posted) > 0.05 {
		current := fmt.Sprintf("[%s-%s] error rate execeeds limit [posted=%d, failed=%d, limit=0.05, rate=%d]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, s.failures/s.posted)
		if current != s.lastlog {
			logrus.Error(current)
			s.lastlog = current
		}
		return
	}

	if s.failures > 0 {
		current := fmt.Sprintf("[%s-%s] error rated increases [posted=%d, failed=%d, limit=0.05, rate=%d]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, s.failures/s.posted)
		if current != s.lastlog {
			logrus.Warn(current)
			s.lastlog = current
		}
		return
	}

	current := fmt.Sprintf("[%s-%s] all fine [posted=%d, failed=%d, limit=0.05, rate=%d]", s.start.Format("15:04:05"), s.end.Format("15:04:05"), s.posted, s.failures, s.failures/s.posted)
	if current != s.lastlog {
		logrus.Info(current)
		s.lastlog = current
	}
}
