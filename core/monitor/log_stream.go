package monitor

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type LogStream struct {
	filename  string
	Snapshots []*Snapshot
}

func NewLogStream(filename string) LogStream {
	m := LogStream{filename: filename}
	return m
}

func (m LogStream) Run() {
	f, err := os.Open(m.filename)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}
		if len(line) > 0 {
			m.process(string(line))
		}
	}
}

func (m *LogStream) process(line string) {
	start, _ := extractTime(line)
	start = round(start, time.Second)
	end := start.Add(1 * time.Minute)
	snap := m.findOrCreateSnapshot(start, end)
	if strings.Contains(line, "http status=200") {
		snap.successes++
	}
	if strings.Contains(line, "http status=400") {
		snap.failures++
	}

	if snap.failures > snap.successes {
		logrus.Error(snap)
	} else if snap.failures == snap.successes {
		logrus.Warn(snap)
	} else {
		logrus.Info(snap)
	}

}

func (m *LogStream) findOrCreateSnapshot(start time.Time, end time.Time) *Snapshot {
	for _, s := range m.Snapshots {
		if s.start.Equal(start) && s.end.Equal(end) {
			return s
		}
	}

	s := &Snapshot{id: len(m.Snapshots) + 1, start: start, end: end}
	m.Snapshots = append(m.Snapshots, s)

	return s
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
