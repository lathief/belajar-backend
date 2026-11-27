package main

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /profil", func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok || user != "demo" || subtle.ConstantTimeCompare([]byte(password), []byte("belajar")) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="latihan"`)
			http.Error(w, "kredensial diperlukan", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"user": user})
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8083", mux))
}
