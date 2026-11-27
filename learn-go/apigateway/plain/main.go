package main

import (
	"log"
	"net/http"

	"github.com/lathief/test-go/apigateway/internal/gateway"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /api/products", gateway.Products())
	log.Fatal(http.ListenAndServe("127.0.0.1:8083", mux))
}
