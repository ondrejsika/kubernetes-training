package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter_requests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "example",
		Name:      "requests",
		Help:      "Number of requests",
	})

func main() {
	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(counter_requests)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter_requests.Add(1)

		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":80", nil))
}
