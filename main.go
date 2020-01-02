package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

func main() {
	filename := flag.String("logfile", "/tmp/jobdog.log", "name of the logfile to monitor")
	minConsecutiveSuccesses := flag.Int("minsuccesses", 10, "min number of consecutive successes")
	flag.Parse()
	m := NewLogStream(*filename, *minConsecutiveSuccesses)
	m.Run()
}

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

	var bytesRead int64 = 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				stat, _ := os.Stat(m.filename)
				if stat.Size() < int64(bytesRead) {
					f.Close()
					f, _ = os.Open(m.filename)
					reader = bufio.NewReader(f)
					bytesRead = 0
				}
			} else {
				break
			}
		}
		if len(line) > 0 {
			m.process(string(line))
			runeCount := int64(utf8.RuneCountInString(line))
			bytesRead = bytesRead + runeCount
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
