package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lathief/test-go/worker/internal/worker"
)

func main() {
	if err := worker.RunHTTP("127.0.0.1:8082", func(queue *worker.Queue) http.Handler {
		router := gin.New()
		router.POST("/jobs", func(c *gin.Context) {
			status, result := queue.Submit()
			c.JSON(status, result)
		})
		return router
	}); err != nil {
		log.Fatal(err)
	}
}
