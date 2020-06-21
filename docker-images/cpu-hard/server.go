package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		sum := 0.
		for i := 1.; i < 100000000; i++ {
			sum = math.Sqrt(sum + i)
		}
		duration := time.Now().Sub(started)
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%v\n", duration.Seconds())))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
