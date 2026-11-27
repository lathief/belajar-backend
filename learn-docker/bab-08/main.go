package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /buku", func(w http.ResponseWriter, r *http.Request) {
		buku, err := os.ReadFile("data/buku.txt")
		if err != nil {
			log.Printf("gagal membaca katalog: %v", err)
			http.Error(w, "katalog belum tersedia", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(buku)
	})
	fmt.Println("katalog siap di port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
