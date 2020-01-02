package monitor

import (
	"bufio"
	"io"
	"os"
	"time"
)

type LogStream struct {
	filename  string
	limit     float64
	Snapshots []*Snapshot
}

func NewLogStream(filename string, limit float64) LogStream {
	m := LogStream{filename: filename, limit: limit}
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
				time.Sleep(10 * time.Second)
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
	snap.Process(line)
	snap.Log(m.limit)
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
