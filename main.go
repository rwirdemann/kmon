package main

import (
	"flag"

	"github.com/rwirdemann/kmon/monitor"
)

func main() {
	filename := flag.String("logfile", "/tmp/jobdog.log", "name of the logfile to monitor")
	limit := flag.Float64("limit", 0.1, "acceptable error limit")
	flag.Parse()
	m := monitor.NewLogStream(*filename, *limit)
	m.Run()
}
