package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
	"os"
)

func randomZeroOrOne() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	return int(n.Int64()), err
}

func main() {
	health, _ := randomZeroOrOne()
	hostname, _ := os.Hostname()
	if health == 1 {
		fmt.Println(hostname +" Server will be ready forever")
	} else {
		fmt.Println(hostname +" Server will be ready for 30 seconds, then 30 second not, and again")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if health == 1 {
			w.WriteHeader(200)
			w.Write([]byte(hostname +" Server will be ready forever\n"))
		} else {
			if time.Now().Second() > 30 {
				w.WriteHeader(500)
				w.Write([]byte(hostname +" Not Ready\n"))
			} else {
				w.WriteHeader(200)
				w.Write([]byte(hostname +" Ready\n"))
			}
		}
	})
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if health == 1 {
			w.WriteHeader(200)
			w.Write([]byte("Ready"))
		} else {
			if time.Now().Second() < 30 {
				w.WriteHeader(200)
				w.Write([]byte("Ready"))
			} else {
				w.WriteHeader(500)
				w.Write([]byte("Not Ready"))
			}
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
