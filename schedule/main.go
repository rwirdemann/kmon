package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		id := rand.Intn(9000) + 1
		status := status()
		if status == http.StatusOK {
			logrus.WithFields(logrus.Fields{"id": id, "http status": status}).Info("job successfully posted")
		} else {
			logrus.WithFields(logrus.Fields{"id": id, "http status": status}).Error("job posting failed")
		}
		n := rand.Intn(2000)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func status() int {
	i := rand.Intn(2)
	if i == 0 {
		return http.StatusOK
	}
	return http.StatusBadRequest
}
