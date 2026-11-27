package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lathief/test-go/ratelimiter/internal/limit"
)

func main() {
	router := gin.Default()
	var window limit.Window
	router.GET("/ping", func(c *gin.Context) {
		if wait := window.RetryAfter(); wait > 0 {
			c.Header("Retry-After", strconv.Itoa(wait))
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.String(http.StatusOK, "pong")
	})
	log.Fatal(router.Run("127.0.0.1:8082"))
}
