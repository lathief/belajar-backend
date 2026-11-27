package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lathief/test-go/ratelimiter/internal/limit"
)

func main() {
	mux := http.NewServeMux()
	var window limit.Window
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		if wait := window.RetryAfter(); wait > 0 {
			w.Header().Set("Retry-After", strconv.Itoa(wait))
			http.Error(w, "terlalu banyak permintaan", http.StatusTooManyRequests)
			return
		}
		fmt.Fprint(w, "pong")
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8083", mux))
}
