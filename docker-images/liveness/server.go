package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
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
	if health {
		fmt.Printf("Server will be healthy forever")
	} else {
		fmt.Println("Server will be healthy for 30 seconds")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if health {
			w.Write([]byte("Server will be healthy forever"))
		} else {
			w.Write([]byte("Server will be healthy for 30 seconds"))
		}
	})
	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if health {
			w.WriteHeader(200)
			w.Write([]byte("ok always"))
		} else {
			if duration.Seconds() < 30 {
				w.WriteHeader(200)
				w.Write([]byte("ok for 30 seconds"))
			} else {
				w.WriteHeader(500)
				w.Write([]byte("err"))
			}
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
