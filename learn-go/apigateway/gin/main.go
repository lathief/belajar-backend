package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lathief/test-go/apigateway/internal/gateway"
)

func main() {
	router := gin.Default()
	router.GET("/api/products", gin.WrapH(gateway.Products()))
	log.Fatal(router.Run("127.0.0.1:8082"))
}
