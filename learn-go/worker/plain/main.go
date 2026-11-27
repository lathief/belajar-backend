package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lathief/test-go/worker/internal/worker"
)

func main() {
	if err := worker.RunHTTP("127.0.0.1:8083", func(queue *worker.Queue) http.Handler {
		mux := http.NewServeMux()
		mux.HandleFunc("POST /jobs", func(w http.ResponseWriter, r *http.Request) {
			status, result := queue.Submit()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				log.Printf("write job response: %v", err)
			}
		})
		return mux
	}); err != nil {
		log.Fatal(err)
	}
}
