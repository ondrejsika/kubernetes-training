package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func randomZeroOrOne() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	return int(n.Int64()), err
}

func main() {
	started := time.Now()
	health, _ := randomZeroOrOne()
	if health == 1 {
		fmt.Println("Server will be healthy forever")
	} else {
		fmt.Println("Server will be healthy for 30 seconds")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if health == 1 {
			w.Write([]byte("Server will be healthy forever"))
		} else {
			w.Write([]byte("Server will be healthy for 30 seconds"))
		}
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if health == 1 {
			w.WriteHeader(200)
			w.Write([]byte("ok always"))
		} else {
			if duration.Seconds() > 30 {
				w.WriteHeader(500)
				w.Write([]byte("err"))
			} else {
				w.WriteHeader(200)
				w.Write([]byte("ok for 30 seconds"))
			}
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
