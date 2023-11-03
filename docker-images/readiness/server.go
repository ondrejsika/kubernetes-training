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
	health := random()
	hostname, err := os.Hostname()
	handleError(err)
	if health {
		fmt.Printf("[%s] Server will be ready forever\n", hostname)
	} else {
		fmt.Printf("[%s] Server will be ready only for first 30 seconds in minute\n", hostname)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if health {
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprintf("[%s] Server will be healthy forever\n", hostname)))
		} else {
			if time.Now().Second() > 30 {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("[%s] Not Ready\n", hostname)))
			} else {
				w.WriteHeader(200)
				w.Write([]byte(fmt.Sprintf("[%s] Ready\n", hostname)))
			}
		}
	})
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if health {
			w.WriteHeader(200)
			w.Write([]byte("OK\n"))
		} else {
			if time.Now().Second() < 30 {
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
