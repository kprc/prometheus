package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var opsProced prometheus.Counter

func main() {

	opsProced = promauto.NewCounter(prometheus.CounterOpts{
		Name: "A_test_counter",
		Help: "Help for test counter",
	})

	time.Sleep(time.Second)

	go func() {
		for {
			opsProced.Inc()
			time.Sleep(time.Second * 2)
		}

	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
