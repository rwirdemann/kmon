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
	filename                string
	successes               int
	consecutiveSuccesses    int
	minConsecutiveSuccesses int
	failures                int
	posted                  int
}

func NewLogStream(filename string, minConsecutiveSuccesses int) LogStream {
	m := LogStream{filename: filename, consecutiveSuccesses: 1000, minConsecutiveSuccesses: minConsecutiveSuccesses}
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
	if strings.Contains(line, "Published job") {
		m.posted++
	}

	if strings.Contains(line, "status=200") {
		m.successes++
		m.consecutiveSuccesses++
	}
	if strings.Contains(line, "status=400") {
		m.failures++
		m.consecutiveSuccesses = 0
	}

	if m.consecutiveSuccesses >= m.minConsecutiveSuccesses {
		logrus.WithFields(logrus.Fields{"published": m.posted, "success": m.successes, "failures": m.failures}).Info()
	} else {
		logrus.WithFields(logrus.Fields{"published": m.posted, "success": m.successes, "failures": m.failures}).Error()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
