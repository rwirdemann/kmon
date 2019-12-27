package monitor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type LogStream struct {
	filename  string
	snapshots []Snapshot
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
	snap := m.getOrCreateSnapshot(start)
	if strings.Contains(line, "http status=200") {
		snap.successes++
	}
	if strings.Contains(line, "http status=400") {
		snap.failures++
	}
	fmt.Println(snap)
}

func (m *LogStream) getOrCreateSnapshot(t time.Time) *Snapshot {
	if len(m.snapshots) > 0 {
		s := &m.snapshots[len(m.snapshots)-1]
		if t.After(s.start) {
			s.end = t
		} else {
			return s
		}
	}
	end := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, t.Location())
	s := Snapshot{id: len(m.snapshots) + 1, start: t, end: end}
	m.snapshots = append(m.snapshots, s)
	return &s
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
