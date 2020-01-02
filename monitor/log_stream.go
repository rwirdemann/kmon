package monitor

import (
	"bufio"
	"fmt"
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
	if strings.Contains(line, "status=200") {
		m.posted++
		m.successes++
		m.consecutiveSuccesses++
	}
	if strings.Contains(line, "status=400") {
		m.posted++
		m.failures++
		m.consecutiveSuccesses = 0
	}

	s := fmt.Sprintf("published: %d, successes: %d, failures: %d", m.posted, m.successes, m.failures)
	if m.consecutiveSuccesses >= m.minConsecutiveSuccesses {
		logrus.Info(s)
	} else {
		logrus.Error(s)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
