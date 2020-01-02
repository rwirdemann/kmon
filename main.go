package main

import (
	"flag"

	"github.com/rwirdemann/kmon/monitor"
)

func main() {
	filename := flag.String("logfile", "/tmp/jobdog.log", "name of the logfile to monitor")
	minConsecutiveSuccesses := flag.Int("minsuccesses", 10, "min number of consecutive successes")
	flag.Parse()
	m := monitor.NewLogStream(*filename, *minConsecutiveSuccesses)
	m.Run()
}
