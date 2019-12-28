package main

import (
	"github.com/rwirdemann/kmon/core/monitor"
)

func main() {
	m := monitor.NewLogStream("job-postings.log")
	m.Run()
}
