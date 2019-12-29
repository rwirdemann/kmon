package main

import (
	"flag"

	"github.com/rwirdemann/kmon/monitor"
)

func main() {
	filename := flag.String("logfile", "/tmp/jobdog.log", "name of the logfile to monitor")
	flag.Parse()
	m := monitor.NewLogStream(*filename)
	m.Run()
}
