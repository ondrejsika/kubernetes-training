package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func random() bool {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	handleError(err)
	return int(n.Int64()) == 1
}

func main() {
	started := time.Now()
	health := random()
	hostname, err := os.Hostname()
	handleError(err)
	if health {
		fmt.Printf("[%s] Server will be healthy forever\n", hostname)
	} else {
		fmt.Printf("[%s] Server will be healthy for 30 seconds\n", hostname)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if health {
			w.Write([]byte(fmt.Sprintf("[%s] Server will be healthy forever\n", hostname)))
		} else {
			w.Write([]byte(fmt.Sprintf("[%s] Server will be healthy for 30 seconds\n", hostname)))
		}
	})
	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if health {
			w.WriteHeader(200)
			w.Write([]byte("OK\n"))
		} else {
			if duration.Seconds() < 30 {
				w.WriteHeader(200)
				w.Write([]byte("OK\n"))
			} else {
				w.WriteHeader(500)
				w.Write([]byte("ERR\n"))
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
